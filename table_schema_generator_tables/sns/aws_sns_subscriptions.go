package sns

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/mitchellh/mapstructure"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsSnsSubscriptionsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsSnsSubscriptionsGenerator{}

func (x *TableAwsSnsSubscriptionsGenerator) GetTableName() string {
	return "aws_sns_subscriptions"
}

func (x *TableAwsSnsSubscriptionsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsSnsSubscriptionsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsSnsSubscriptionsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsSnsSubscriptionsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().SNS
			config := sns.ListSubscriptionsInput{}
			for {
				output, err := svc.ListSubscriptions(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, output.Subscriptions, func(result any) (any, error) {
					c := client.(*aws_client.Client)
					svc := c.AwsServices().SNS
					item := result.(types.Subscription)
					s := Subscription{
						SubscriptionArn:	item.SubscriptionArn,
						Owner:			item.Owner,
						Protocol:		item.Protocol,
						TopicArn:		item.TopicArn,
						Endpoint:		item.Endpoint,
					}

					if aws.ToString(item.SubscriptionArn) == "PendingConfirmation" {
						return s, nil

					}

					attrs, err := svc.GetSubscriptionAttributes(ctx, &sns.GetSubscriptionAttributesInput{SubscriptionArn: item.SubscriptionArn})
					if err != nil {
						return nil, err
					}
					dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{WeaklyTypedInput: true, Result: &s})
					if err != nil {
						return nil, err
					}
					if err := dec.Decode(attrs.Attributes); err != nil {
						return nil, err
					}
					return s, nil

				})
				if aws.ToString(output.NextToken) == "" {
					break
				}
				config.NextToken = output.NextToken
			}
			return nil
		},
	}
}

type Subscription struct {
	Endpoint	*string

	Owner	*string

	Protocol	*string

	SubscriptionArn	*string

	TopicArn	*string

	ConfirmationWasAuthenticated	*bool

	DeliveryPolicy	*string

	EffectiveDeliveryPolicy	*string

	FilterPolicy	*string

	PendingConfirmation	*bool

	RawMessageDelivery	*bool

	RedrivePolicy	*string

	SubscriptionRoleArn	*string

	UnknownFields	map[string]interface{}	`mapstructure:",remain"`
}

func (x *TableAwsSnsSubscriptionsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("sns")
}

func (x *TableAwsSnsSubscriptionsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("delivery_policy").ColumnType(schema.ColumnTypeJSON).
			Extractor(aws_client.MarshaledJsonExtractor("DeliveryPolicy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("redrive_policy").ColumnType(schema.ColumnTypeJSON).
			Extractor(aws_client.MarshaledJsonExtractor("RedrivePolicy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("endpoint").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SubscriptionArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("filter_policy").ColumnType(schema.ColumnTypeJSON).
			Extractor(aws_client.MarshaledJsonExtractor("FilterPolicy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("topic_arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("effective_delivery_policy").ColumnType(schema.ColumnTypeJSON).
			Extractor(aws_client.MarshaledJsonExtractor("EffectiveDeliveryPolicy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("unknown_fields").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owner").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("protocol").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("confirmation_was_authenticated").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pending_confirmation").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("raw_message_delivery").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subscription_role_arn").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableAwsSnsSubscriptionsGenerator) GetSubTables() []*schema.Table {
	return nil
}
