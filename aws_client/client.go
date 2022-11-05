package aws_client

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	smithy "github.com/aws/smithy-go"
	"github.com/selefra/selefra-utils/pkg/if_expression"
)

const (
	defaultRegion         = "us-east-1"
	cloudfrontScopeRegion = defaultRegion
)

var envVarsToCheck = []string{
	"AWS_PROFILE",
	"AWS_ACCESS_KEY_ID",
	"AWS_SECRET_ACCESS_KEY",
	"AWS_CONFIG_FILE",
	"AWS_ROLE_ARN",
	"AWS_SESSION_TOKEN",
	"AWS_SHARED_CREDENTIALS_FILE",
}

func checkEnvVariables() string {
	var result []string
	for _, v := range envVarsToCheck {
		if _, present := os.LookupEnv(v); present {
			result = append(result, v)
		}
	}
	return strings.Join(result, ",")
}

type Option func(client *Client)

type AwsOrg struct {
	OrganizationUnits           []string    `yaml:"organization_units,omitempty"  mapstructure:"organization_units"`
	AdminAccount                *AwsAccount `yaml:"admin_account"  mapstructure:"admin_account"`
	MemberCredentials           *AwsAccount `yaml:"member_trusted_principal"  mapstructure:"member_trusted_principal"`
	ChildAccountRoleName        string      `yaml:"member_role_name,omitempty"  mapstructure:"member_role_name"`
	ChildAccountRoleSessionName string      `yaml:"member_role_session_name,omitempty"  mapstructure:"member_role_session_name"`
	ChildAccountExternalID      string      `yaml:"member_external_id,omitempty"  mapstructure:"member_external_id"`
	ChildAccountRegions         []string    `yaml:"member_regions,omitempty"  mapstructure:"member_regions"`
}

type AwsAccount struct {
	AccountName            string   `yaml:"account_name,omitempty"  mapstructure:"account_name"`
	SharedConfigProfile    string   `yaml:"shared_config_profile,omitempty"  mapstructure:"shared_config_profile"`
	SharedConfigFiles      []string `yaml:"shared_config_files"  mapstructure:"shared_config_files"`
	SharedCredentialsFiles []string `yaml:"shared_credentials_files"  mapstructure:"shared_credentials_files"`
	RoleARN                string   `yaml:"role_arn,omitempty"  mapstructure:"role_arn"`
	RoleSessionName        string   `yaml:"role_session_name,omitempty"  mapstructure:"role_session_name"`
	ExternalID             string   `yaml:"external_id,omitempty"  mapstructure:"external_id"`
	DefaultRegion          string   `yaml:"default_region,omitempty"  mapstructure:"default_region"`
	Regions                []string `yaml:"regions,omitempty"  mapstructure:"regions"`
	source                 string
}

type AwsProviderConfigs struct {
	Providers []AwsProviderConfig `yaml:"providers"  mapstructure:"providers"`
}

const (
	DefaultMaxAttempts = 10
	DefaultMaxBackoff  = 30
)

type AwsProviderConfig struct {
	Accounts     []AwsAccount `yaml:"accounts"  mapstructure:"accounts"`
	MaxAttempts  int          `yaml:"max_attempts,omitempty"  mapstructure:"max_attempts"`
	MaxBackoff   int          `yaml:"max_backoff,omitempty" mapstructure:"max_backoff"`
	Organization *AwsOrg      `yaml:"org" mapstructure:"org"`
	Regions      []string
}

func verifyRegions(regions []string) error {
	if serviceRegionDataTransport == nil {
		return errors.New("service Region data not initialized")
	}

	var hasWildcard bool
	for i, region := range regions {
		if region == "*" {
			hasWildcard = true
		}
		if (i != 0 && region == "*") || (i > 0 && hasWildcard) {
			return errors.New("Region wildcard \"*\" is only supported as first argument")
		}

		_, regionExist := serviceRegionDataTransport.regionSet[region]
		if !hasWildcard && !regionExist {
			return fmt.Errorf("Region %s is not supported", region)
		}
	}
	return nil
}

func filterDisabledRegions(regions []string, enabledRegions []ec2types.Region) []string {
	regionsMap := map[string]bool{}
	for _, r := range enabledRegions {
		if r.RegionName != nil && r.OptInStatus != nil && *r.OptInStatus != "not-opted-in" {
			regionsMap[*r.RegionName] = true
		}
	}

	var filteredRegions []string

	if isAllRegions(regions) {
		for region := range regionsMap {
			filteredRegions = append(filteredRegions, region)
		}
	} else {
		for _, r := range regions {
			if regionsMap[r] {
				filteredRegions = append(filteredRegions, r)
			}
		}
	}
	return filteredRegions
}

func isAllRegions(regions []string) bool {
	err := verifyRegions(regions)
	if err != nil {
		return false
	}

	wildcardAllRegions := false
	if (len(regions) == 1 && regions[0] == "*") || (len(regions) == 0) {
		wildcardAllRegions = true
	}
	return wildcardAllRegions
}

func newAwsConfig(ctx context.Context, cfg *AwsProviderConfig, account AwsAccount, stsClient *sts.Client) (aws.Config, error) {
	var err error
	var awsCfg aws.Config
	configFns := []func(*config.LoadOptions) error{
		config.WithDefaultRegion(defaultRegion),
		config.WithRetryer(func() aws.Retryer {
			return retry.NewStandard(func(o *retry.StandardOptions) {
				o.MaxAttempts = cfg.MaxAttempts
				o.MaxBackoff = time.Second * time.Duration(cfg.MaxBackoff)
			})
		}),
	}

	if account.SharedConfigProfile != "" {
		configFns = append(configFns, config.WithSharedConfigProfile(account.SharedConfigProfile))
	}

	if len(account.SharedConfigFiles) != 0 {
		configFns = append(configFns, config.WithSharedConfigFiles(account.SharedConfigFiles))
	}

	if len(account.SharedCredentialsFiles) != 0 {
		configFns = append(configFns, config.WithSharedCredentialsFiles(account.SharedCredentialsFiles))
	}

	if account.DefaultRegion != "" {
		configFns = append(configFns, config.WithDefaultRegion(account.DefaultRegion))
	}

	awsCfg, err = config.LoadDefaultConfig(ctx, configFns...)

	if err != nil {
		return awsCfg, err
	}

	if account.RoleARN != "" {
		opts := make([]func(*stscreds.AssumeRoleOptions), 0, 1)
		if account.ExternalID != "" {
			opts = append(opts, func(opts *stscreds.AssumeRoleOptions) {
				opts.ExternalID = &account.ExternalID
			})
		}
		if account.RoleSessionName != "" {
			opts = append(opts, func(opts *stscreds.AssumeRoleOptions) {
				opts.RoleSessionName = account.RoleSessionName
			})
		}
		if stsClient == nil {
			stsClient = sts.NewFromConfig(awsCfg)
		}
		provider := stscreds.NewAssumeRoleProvider(stsClient, account.RoleARN, opts...)

		awsCfg.Credentials = aws.NewCredentialsCache(provider)
	}

	if _, err := awsCfg.Credentials.Retrieve(ctx); err != nil {
		var ae smithy.APIError
		if errors.As(err, &ae) {
			if strings.Contains(ae.ErrorCode(), "InvalidClientTokenId") {
				return awsCfg, errors.New("the credentials being used to assume role are invalid. Please check that your credentials are valid in the Partition you are using. If you are using a Partition other than the AWS commercial Region, be sure set the default_region attribute in the config file")
			}
		}
		return awsCfg, errors.New("couldn't find any credentials in environment variables or configuration files ")
	}
	return awsCfg, nil
}

type Client struct {
	config aws.Config

	Partition string
	AccountID string
	Region    string

	AutoscalingNamespace string
	WAFScope             types.Scope

	accountAwsServiceManager *AwsServicesManager
}

func NewClients(configs AwsProviderConfigs) ([]*Client, error) {
	var clients []*Client

	for _, c := range configs.Providers {
		cls, err := newClient(c)
		if err != nil {
			return nil, err
		}
		clients = append(clients, cls...)
	}
	return clients, nil
}

func newClient(config AwsProviderConfig) ([]*Client, error) {
	ctx := context.Background()
	var stsClient *sts.Client
	var clients []*Client
	var accounts []AwsAccount

	config.MaxAttempts = if_expression.ReturnInt(config.MaxAttempts == 0, DefaultMaxAttempts, config.MaxAttempts)
	config.MaxBackoff = if_expression.ReturnInt(config.MaxBackoff == 0, DefaultMaxBackoff, config.MaxBackoff)

	if config.Organization != nil {
		var err error
		accounts, stsClient, err = loadOrgAccounts(ctx, &config)
		if err != nil {
			var ae smithy.APIError
			if errors.As(err, &ae) {
				if strings.Contains(ae.ErrorCode(), "AccessDenied") {
					return nil, errors.New("failed to list Org member accounts. Make sure that your credentials have the proper permissions")
				}
			}
			return nil, err
		}
	}

	for _, account := range config.Accounts {
		account.source = "AccountID"
		accounts = append(accounts, account)
	}

	if len(accounts) == 0 {
		accounts = append(accounts, AwsAccount{
			AccountName: "default",
			source:      "default",
		})
	}
	for _, account := range accounts {
		localRegions := account.Regions
		if len(localRegions) == 0 {
			localRegions = config.Regions
		}

		if err := verifyRegions(localRegions); err != nil {
			return nil, err
		}

		awsCfg, err := newAwsConfig(ctx, &config, account, stsClient)
		if err != nil {
			if account.source == "org" {
				principal := "unknown principal"

				awsAdminCfg, _ := newAwsConfig(ctx, &config, *config.Organization.AdminAccount, nil)
				output, accountErr := sts.NewFromConfig(awsAdminCfg).GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
				if accountErr == nil {
					principal = *output.Arn
				}

				fmt.Printf("ensure that %s has access to be able perform `sts:AssumeRole` on %s \n", principal, account.RoleARN)

				continue
			}
			var ae smithy.APIError
			if errors.As(err, &ae) {
				if strings.Contains(ae.ErrorCode(), "AccessDenied") {
					fmt.Printf("failed to retrieve credentials for AccountID %s. AWS Error: %s, detected aws env variables: %s \n", account.AccountName, err.Error(), checkEnvVariables())
					continue
				}
			}

			return nil, err
		}

		res, err := ec2.NewFromConfig(awsCfg).DescribeRegions(ctx,
			&ec2.DescribeRegionsInput{AllRegions: aws.Bool(false)},
			func(o *ec2.Options) {
				o.Region = defaultRegion
				if account.DefaultRegion != "" {
					o.Region = account.DefaultRegion
				}

				if len(localRegions) > 0 && !isAllRegions(localRegions) {
					o.Region = localRegions[0]
				}
			})

		if err != nil {
			fmt.Printf("failed to find disabled regions for AccountID %s. AWS Error: %s", account.AccountName, err.Error())
			continue
		}

		account.Regions = filterDisabledRegions(localRegions, res.Regions)

		if len(account.Regions) == 0 {
			fmt.Printf("no enabled regions provided in config for AccountID %s", account.AccountName)
			continue
		}
		awsCfg.Region = account.Regions[0]
		output, err := sts.NewFromConfig(awsCfg).GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
		if err != nil {

			fmt.Printf("failed to get caller identity. AWS Error: %s", err.Error())
			continue
		}
		iamArn, err := arn.Parse(*output.Arn)
		if err != nil {
			return nil, err
		}

		client := NewAwsClient(awsCfg, SetAccount(*output.Account), SetPartition(iamArn.Partition))

		accountAwsServiceManager := NewAwsServiceCache(&awsCfg)
		client.accountAwsServiceManager = accountAwsServiceManager
		var register bool

		for _, region := range account.Regions {
			if region == defaultRegion {
				register = true
			}
			accountAwsServiceManager.initAwsServices(client.Partition, region)
		}
		if !register {
			accountAwsServiceManager.initAwsServices(client.Partition, defaultRegion)
		}

		accountAwsServiceManager.initAwsServices(client.Partition, cloudfrontScopeRegion)

		if client.accountAwsServiceManager.CacheCount() == 0 {
			return nil, errors.New("no accounts instantiated")
		}
		clients = append(clients, client)
	}

	if len(clients) == 0 {
		return nil, errors.New("no accounts instantiated")
	}

	return clients, nil

}

func (c *Client) Copy(opts ...Option) *Client {
	cc := &Client{
		config:                   c.config,
		AccountID:                c.AccountID,
		Region:                   c.Region,
		AutoscalingNamespace:     c.AutoscalingNamespace,
		WAFScope:                 c.WAFScope,
		Partition:                c.Partition,
		accountAwsServiceManager: c.accountAwsServiceManager,
	}
	for _, opt := range opts {
		opt(cc)
	}
	return cc
}

func (c *Client) GetAccount() string {
	return c.AccountID
}

func (c *Client) GetRegion() string {
	return c.Region
}

func (c *Client) GetWAFScope() types.Scope {
	return c.WAFScope
}

func (c *Client) GetPartition() string {
	return c.Partition
}

func (c *Client) GetAutoscalingNamespace() string {
	return c.AutoscalingNamespace
}

func SetRegion(region string) Option {
	return func(c *Client) {
		if region == "" {
			return
		}
		c.Region = region
	}
}

func SetPartition(partition string) Option {
	return func(c *Client) {
		if partition == "" {
			return
		}
		c.Partition = partition
	}
}

func SetAccount(account string) Option {
	return func(c *Client) {
		if account == "" {
			return
		}
		c.AccountID = account
	}
}

func SetAutoscalingNamespace(autoscalingNamespace string) Option {
	return func(client *Client) {
		client.AutoscalingNamespace = autoscalingNamespace
	}
}

func SetWAFScope(cp types.Scope) Option {
	return func(client *Client) {
		client.WAFScope = cp
	}
}

func (c *Client) IsNotFoundError(err error) bool {
	if isNotFoundError(err) {
		return true
	}
	return false
}

func (c *Client) IsAccessDeniedError(err error) bool {
	if isAccessDeniedError(err) {
		return true
	}
	return false
}

func (c *Client) AwsServices() *AwsServices {

	awsServices := c.accountAwsServiceManager.Get(c.Partition, c.Region)
	if awsServices == nil && c.WAFScope == types.ScopeCloudfront {

		return c.accountAwsServiceManager.Get(c.Partition, "")
	}
	return awsServices
}

func (c *Client) ARN(service string, idParts ...string) string {
	return makeARN(service, c.Partition, c.AccountID, c.GetRegion(), idParts...).String()
}

func NewAwsClient(cfg aws.Config, ops ...Option) *Client {
	cli := &Client{
		config:               cfg,
		Region:               defaultRegion,
		AutoscalingNamespace: "",
		WAFScope:             "",
		Partition:            "",
	}
	for _, op := range ops {
		op(cli)
	}
	return cli
}

func (c *Client) AccountGlobalARN(service string, idParts ...string) string {
	return makeARN(service, c.Partition, c.AccountID, "", idParts...).String()
}
