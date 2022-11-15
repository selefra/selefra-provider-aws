module github.com/selefra/selefra-provider-aws

go 1.19

require (
	github.com/aws/aws-sdk-go-v2 v1.17.1
	github.com/aws/aws-sdk-go-v2/config v1.17.10
	github.com/aws/aws-sdk-go-v2/credentials v1.12.23
	github.com/aws/aws-sdk-go-v2/feature/s3/manager v1.11.37
	github.com/aws/aws-sdk-go-v2/service/accessanalyzer v1.17.0
	github.com/aws/aws-sdk-go-v2/service/acm v1.15.2
	github.com/aws/aws-sdk-go-v2/service/apigateway v1.15.22
	github.com/aws/aws-sdk-go-v2/service/apigatewayv2 v1.12.20
	github.com/aws/aws-sdk-go-v2/service/applicationautoscaling v1.15.20
	github.com/aws/aws-sdk-go-v2/service/apprunner v1.15.0
	github.com/aws/aws-sdk-go-v2/service/appsync v1.15.12
	github.com/aws/aws-sdk-go-v2/service/athena v1.18.12
	github.com/aws/aws-sdk-go-v2/service/autoscaling v1.23.18
	github.com/aws/aws-sdk-go-v2/service/backup v1.17.11
	github.com/aws/aws-sdk-go-v2/service/cloudformation v1.23.0
	github.com/aws/aws-sdk-go-v2/service/cloudfront v1.20.7
	github.com/aws/aws-sdk-go-v2/service/cloudhsmv2 v1.13.21
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v1.19.2
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.21.8
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.16.0
	github.com/aws/aws-sdk-go-v2/service/codebuild v1.19.19
	github.com/aws/aws-sdk-go-v2/service/codepipeline v1.13.19
	github.com/aws/aws-sdk-go-v2/service/cognitoidentity v1.14.3
	github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider v1.21.1
	github.com/aws/aws-sdk-go-v2/service/configservice v1.27.2
	github.com/aws/aws-sdk-go-v2/service/databasemigrationservice v1.21.14
	github.com/aws/aws-sdk-go-v2/service/dax v1.11.19
	github.com/aws/aws-sdk-go-v2/service/directconnect v1.17.20
	github.com/aws/aws-sdk-go-v2/service/docdb v1.19.13
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.17.3
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.65.0
	github.com/aws/aws-sdk-go-v2/service/ecr v1.17.20
	github.com/aws/aws-sdk-go-v2/service/ecrpublic v1.13.19
	github.com/aws/aws-sdk-go-v2/service/ecs v1.18.26
	github.com/aws/aws-sdk-go-v2/service/efs v1.17.18
	github.com/aws/aws-sdk-go-v2/service/eks v1.22.3
	github.com/aws/aws-sdk-go-v2/service/elasticache v1.22.12
	github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk v1.14.20
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing v1.14.20
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.18.22
	github.com/aws/aws-sdk-go-v2/service/elasticsearchservice v1.16.12
	github.com/aws/aws-sdk-go-v2/service/emr v1.20.13
	github.com/aws/aws-sdk-go-v2/service/eventbridge v1.16.17
	github.com/aws/aws-sdk-go-v2/service/firehose v1.14.21
	github.com/aws/aws-sdk-go-v2/service/fsx v1.25.4
	github.com/aws/aws-sdk-go-v2/service/glacier v1.13.19
	github.com/aws/aws-sdk-go-v2/service/glue v1.34.0
	github.com/aws/aws-sdk-go-v2/service/guardduty v1.16.2
	github.com/aws/aws-sdk-go-v2/service/iam v1.18.23
	github.com/aws/aws-sdk-go-v2/service/inspector v1.12.19
	github.com/aws/aws-sdk-go-v2/service/inspector2 v1.8.3
	github.com/aws/aws-sdk-go-v2/service/iot v1.30.0
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.15.21
	github.com/aws/aws-sdk-go-v2/service/kms v1.18.15
	github.com/aws/aws-sdk-go-v2/service/lambda v1.24.8
	github.com/aws/aws-sdk-go-v2/service/lightsail v1.23.2
	github.com/aws/aws-sdk-go-v2/service/mq v1.13.15
	github.com/aws/aws-sdk-go-v2/service/neptune v1.18.0
	github.com/aws/aws-sdk-go-v2/service/organizations v1.16.15
	github.com/aws/aws-sdk-go-v2/service/qldb v1.14.20
	github.com/aws/aws-sdk-go-v2/service/rds v1.27.0
	github.com/aws/aws-sdk-go-v2/service/redshift v1.26.13
	github.com/aws/aws-sdk-go-v2/service/resourcegroups v1.12.20
	github.com/aws/aws-sdk-go-v2/service/route53 v1.22.4
	github.com/aws/aws-sdk-go-v2/service/route53domains v1.12.19
	github.com/aws/aws-sdk-go-v2/service/s3 v1.29.1
	github.com/aws/aws-sdk-go-v2/service/s3control v1.24.2
	github.com/aws/aws-sdk-go-v2/service/sagemaker v1.52.0
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.16.4
	github.com/aws/aws-sdk-go-v2/service/sesv2 v1.15.0
	github.com/aws/aws-sdk-go-v2/service/shield v1.17.11
	github.com/aws/aws-sdk-go-v2/service/sns v1.18.3
	github.com/aws/aws-sdk-go-v2/service/sqs v1.19.12
	github.com/aws/aws-sdk-go-v2/service/ssm v1.31.3
	github.com/aws/aws-sdk-go-v2/service/sts v1.17.1
	github.com/aws/aws-sdk-go-v2/service/transfer v1.23.2
	github.com/aws/aws-sdk-go-v2/service/waf v1.11.19
	github.com/aws/aws-sdk-go-v2/service/wafregional v1.12.20
	github.com/aws/aws-sdk-go-v2/service/wafv2 v1.23.0
	github.com/aws/aws-sdk-go-v2/service/workspaces v1.24.0
	github.com/aws/aws-sdk-go-v2/service/xray v1.13.21
	github.com/aws/smithy-go v1.13.4
	github.com/basgys/goxml2json v1.1.0
	github.com/gocarina/gocsv v0.0.0-20220927221512-ad3251f9fa25
	github.com/golang/mock v1.6.0
	github.com/mitchellh/mapstructure v1.5.0
	github.com/pkg/errors v0.9.1
	github.com/rs/zerolog v1.28.0
	github.com/selefra/selefra-provider-sdk v0.0.10
	github.com/selefra/selefra-utils v0.0.2
	github.com/songzhibin97/go-ognl v0.0.2
	github.com/spf13/viper v1.13.0
	github.com/stretchr/testify v1.8.1
	golang.org/x/sync v0.1.0
)

require (
	github.com/Masterminds/squirrel v1.5.3 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.4.9 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.12.19 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.25 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.19 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.26 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.0.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.9.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.1.20 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.7.19 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.19 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.13.19 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.11.25 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.13.8 // indirect
	github.com/bitly/go-simplejson v0.5.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/doug-martin/goqu/v9 v9.18.0 // indirect
	github.com/emirpasic/gods v1.18.1 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/hashicorp/go-hclog v1.2.0 // indirect
	github.com/hashicorp/go-plugin v1.4.5 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/yamux v0.0.0-20180604194846-3520598351bb // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.13.0 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.1 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.12.0 // indirect
	github.com/jackc/pgx/v4 v4.17.2 // indirect
	github.com/jackc/puddle v1.3.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/lann/builder v0.0.0-20180802200727-47ae307949d0 // indirect
	github.com/lann/ps v0.0.0-20150810152359-62de8c46ede0 // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mitchellh/go-testing-interface v0.0.0-20171004221916-a61a99592b77 // indirect
	github.com/natefinch/lumberjack v2.0.0+incompatible // indirect
	github.com/oklog/run v1.0.0 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pelletier/go-toml/v2 v2.0.5 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/spf13/afero v1.8.2 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.4.1 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.23.0 // indirect
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa // indirect
	golang.org/x/net v0.0.0-20220520000938-2e3eb7b945c2 // indirect
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220519153652-3a47de7e79bd // indirect
	google.golang.org/grpc v1.49.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
