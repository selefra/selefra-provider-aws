package aws_client

import (
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-utils/pkg/if_expression"
	"math/rand"
)

import (
	"context"
)

var AllNamespaces = []string{
	"comprehend", "rds", "sagemaker", "appstream", "elasticmapreduce", "dynamodb", "lambda", "ecs", "cassandra", "ec2", "neptune", "kafka", "custom-resource", "elasticache",
}

func getRegion(regionalMap map[string]*AwsServices) string {
	if len(regionalMap) == 0 {
		return ""
	}
	regions := make([]string, 0)
	for i := range regionalMap {
		regions = append(regions, i)
	}
	randomIndex := rand.Intn(len(regions))
	return regions[randomIndex]
}

func ExpandByPartition() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
		awsClient := client.(*Client)
		clientTaskContextSlice := make([]*schema.ClientTaskContext, 0)
		for partition := range awsClient.accountAwsServiceManager.AwsServicesManagerMap {
			clientTaskContextSlice = append(clientTaskContextSlice, &schema.ClientTaskContext{
				Client: awsClient.Copy(func(client *Client) {
					client.Partition = partition
				}),
				Task:	task.Clone(),
			})
		}
		return clientTaskContextSlice
	}
}

func ExpandByPartitionAndRegion(awsServiceName string) func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
		awsClient := client.(*Client)
		clientTaskContextSlice := make([]*schema.ClientTaskContext, 0)
		for partition, partitionMap := range awsClient.accountAwsServiceManager.AwsServicesManagerMap {
			for region := range partitionMap {
				if !isSupportedServiceForRegion(awsServiceName, region) {
					continue
				}
				clientTaskContextSlice = append(clientTaskContextSlice, &schema.ClientTaskContext{
					Client: awsClient.Copy(func(client *Client) {
						client.Partition = partition
						client.Region = region
					}),
					Task:	task.Clone(),
				})
			}
		}
		return clientTaskContextSlice
	}
}

func ExpandByPartitionAndRegionAndScope(awsServiceName string) func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
		awsClient := client.(*Client)
		clientTaskContextSlice := make([]*schema.ClientTaskContext, 0)
		for partition, partitionMap := range awsClient.accountAwsServiceManager.AwsServicesManagerMap {
			for region := range partitionMap {

				clientTaskContextSlice = append(clientTaskContextSlice, &schema.ClientTaskContext{
					Client: awsClient.Copy(func(client *Client) {
						client.Partition = partition
						client.Region = region
						client.WAFScope = types.ScopeCloudfront
					}),
					Task:	task,
				})

				if !isSupportedServiceForRegion(awsServiceName, region) {
					continue
				}
				clientTaskContextSlice = append(clientTaskContextSlice, &schema.ClientTaskContext{
					Client: awsClient.Copy(func(client *Client) {
						client.Partition = partition
						client.Region = region
					}),
					Task:	task.Clone(),
				})
			}
		}
		return clientTaskContextSlice
	}
}

func isSupportedServiceForRegion(service string, region string) bool {

	if serviceRegionDataTransport == nil {
		return false
	}

	if serviceRegionDataTransport.Partitions == nil {
		return false
	}

	prt := if_expression.ReturnString(serviceRegionDataTransport.region[region] != "", serviceRegionDataTransport.region[region], "aws")

	currentPartition := serviceRegionDataTransport.Partitions[prt]

	if currentPartition.Services[service] == nil {
		return false
	}

	_, ok := currentPartition.Services[service].Regions[region]
	return ok
}
