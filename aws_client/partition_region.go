package aws_client

import (
	_ "embed"
	"encoding/json"

	"github.com/selefra/selefra-utils/pkg/if_expression"
)

var (
	//go:embed endpoint/endpoint.json
	endpoint                   []byte
	serviceRegionDataTransport *ServiceRegionDataTransport
)

func init() {
	regionData := ServiceRegionData{}
	_ = json.Unmarshal(endpoint, &regionData)
	serviceRegionDataTransport = &ServiceRegionDataTransport{
		Partitions: map[string]AwsPartition{},
		region:     make(map[string]string),
		regionSet:  make(map[string]struct{}),
	}

	for _, partition := range regionData.Partitions {
		serviceRegionDataTransport.Partitions[partition.Partition] = partition
		for _, services := range partition.Services {
			for region := range services.Regions {
				serviceRegionDataTransport.region[region] = partition.Partition
				serviceRegionDataTransport.regionSet[region] = struct{}{}
			}
		}
	}
}

type ServiceRegionDataTransport struct {
	Partitions map[string]AwsPartition
	region     map[string]string
	regionSet  map[string]struct{}
}

func isRegionService(service string, region string) bool {
	if serviceRegionDataTransport == nil || serviceRegionDataTransport.Partitions == nil {
		return false
	}
	reg, ok := serviceRegionDataTransport.region[region]
	cur := serviceRegionDataTransport.Partitions[if_expression.ReturnString(ok, reg, "aws")]
	if cur.Services[service] == nil {
		return false
	}
	if _, ok := cur.Services[service].Regions[region]; !ok {
		return false
	}
	return true
}

type ServiceRegionData struct {
	Partitions []AwsPartition
}

type AwsPartition struct {
	Partition     string `json:"Partition"`
	PartitionName string `json:"partitionName"`
	Services      map[string]*Endpoints
}

type Endpoints struct {
	Regions map[string]struct{} `json:"endpoints"`
}
