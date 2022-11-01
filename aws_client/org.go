package aws_client

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	orgTypes "github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func loadOrgAccounts(ctx context.Context, awsConfig *AwsProviderConfig) ([]AwsAccount, *sts.Client, error) {
	if awsConfig.Organization.AdminAccount == nil {
		awsConfig.Organization.AdminAccount = &AwsAccount{
			AccountName:		"Default-Admin-AwsAccount",
			SharedConfigProfile:	"",
		}
	}
	awsCfg, err := newAwsConfig(ctx, awsConfig, *awsConfig.Organization.AdminAccount, nil)
	if err != nil {
		return nil, nil, err
	}
	svc := organizations.NewFromConfig(awsCfg)
	accounts, err := loadAccounts(ctx, awsConfig, svc)
	if err != nil {
		return nil, nil, err
	}
	if awsConfig.Organization.MemberCredentials != nil {
		awsCfg, err = newAwsConfig(ctx, awsConfig, *awsConfig.Organization.MemberCredentials, nil)
		if err != nil {
			return nil, nil, err
		}
	}
	return accounts, sts.NewFromConfig(awsCfg), err
}

func loadAccounts(ctx context.Context, awsConfig *AwsProviderConfig, accountsApi *organizations.Client) ([]AwsAccount, error) {
	var rawAccounts []orgTypes.Account
	var err error
	if len(awsConfig.Organization.OrganizationUnits) > 0 {
		rawAccounts, err = getOUAccounts(ctx, accountsApi, awsConfig.Organization.OrganizationUnits)
	} else {
		rawAccounts, err = getAllAccounts(ctx, accountsApi)
	}

	if err != nil {
		return []AwsAccount{}, err
	}
	accounts := make([]AwsAccount, 0)
	for _, account := range rawAccounts {

		if account.Status != orgTypes.AccountStatusActive {
			continue
		}
		roleArn := arn.ARN{
			Partition:	"aws",
			Service:	"iam",
			Region:		"",
			AccountID:	*account.Id,
			Resource:	"role/" + awsConfig.Organization.ChildAccountRoleName,
		}
		if parsed, err := arn.Parse(aws.ToString(account.Arn)); err == nil {
			roleArn.Partition = parsed.Partition
		}

		accounts = append(accounts, AwsAccount{
			AccountName:		*account.Id,
			RoleARN:		roleArn.String(),
			RoleSessionName:	awsConfig.Organization.ChildAccountRoleSessionName,
			ExternalID:		awsConfig.Organization.ChildAccountExternalID,
			SharedConfigProfile:	awsConfig.Organization.AdminAccount.SharedConfigProfile,
			Regions:		awsConfig.Organization.ChildAccountRegions,
			source:			"org",
		})
	}
	return accounts, err
}

func getOUAccounts(ctx context.Context, accountsApi *organizations.Client, ous []string) ([]orgTypes.Account, error) {
	var rawAccounts []orgTypes.Account

	for _, ou := range ous {
		var paginationToken *string
		for {
			resp, err := accountsApi.ListAccountsForParent(ctx, &organizations.ListAccountsForParentInput{
				NextToken:	paginationToken,
				ParentId:	aws.String(ou),
			})
			if err != nil {
				return nil, err
			}
			rawAccounts = append(rawAccounts, resp.Accounts...)
			if resp.NextToken == nil {
				break
			}
			paginationToken = resp.NextToken
		}
	}
	return rawAccounts, nil
}

func getAllAccounts(ctx context.Context, accountsApi *organizations.Client) ([]orgTypes.Account, error) {
	var rawAccounts []orgTypes.Account
	var paginationToken *string

	for {
		resp, err := accountsApi.ListAccounts(ctx, &organizations.ListAccountsInput{
			NextToken: paginationToken,
		})
		if err != nil {
			return nil, err
		}
		rawAccounts = append(rawAccounts, resp.Accounts...)
		if resp.NextToken == nil {
			break
		}
		paginationToken = resp.NextToken
	}
	return rawAccounts, nil
}
