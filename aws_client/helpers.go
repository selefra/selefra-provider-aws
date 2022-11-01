package aws_client

import (
	"context"
	"errors"
	"reflect"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws/arn"
	smithy "github.com/aws/smithy-go"
)

var notFoundErrorSubstrings = []string{
	"InvalidAMIID.Unavailable",
	"NonExistentQueue",
	"NoSuch",
	"NotFound",
	"ResourceNotFoundException",
	"WAFNonexistentItemException",
	"NoSuchResource",
}

var accessDeniedErrorStrings = map[string]struct{}{
	"AuthorizationError":			{},
	"AccessDenied":				{},
	"AccessDeniedException":		{},
	"InsufficientPrivilegesException":	{},
	"UnauthorizedOperation":		{},
	"Unauthorized":				{},
}

func isNotFoundError(err error) bool {
	var ae smithy.APIError
	if !errors.As(err, &ae) {
		return false
	}
	errorCode := ae.ErrorCode()
	for _, s := range notFoundErrorSubstrings {
		if strings.Contains(errorCode, s) {
			return true
		}
	}
	return false
}

func IgnoreAccessDeniedServiceDisabled(err error) bool {
	var ae smithy.APIError
	if errors.As(err, &ae) {
		switch ae.ErrorCode() {
		case "UnrecognizedClientException":
			return strings.Contains(ae.Error(), "The security token included in the request is invalid")
		case "AWSOrganizationsNotInUseException":
			return true
		case "OptInRequired", "SubscriptionRequiredException", "InvalidClientTokenId":
			return true
		}
	}
	return isAccessDeniedError(err)
}

func isAccessDeniedError(err error) bool {
	var ae smithy.APIError
	if !errors.As(err, &ae) {
		return false
	}
	_, ok := accessDeniedErrorStrings[ae.ErrorCode()]
	return ok
}

func IsAWSError(err error, code ...string) bool {
	var ae smithy.APIError
	if !errors.As(err, &ae) {
		return false
	}
	for _, c := range code {
		if strings.Contains(ae.ErrorCode(), c) {
			return true
		}
	}
	return false
}

func IgnoreNotAvailableRegion(err error) bool {
	var ae smithy.APIError
	if errors.As(err, &ae) {
		if ae.ErrorCode() == "InvalidRequestException" && strings.Contains(ae.ErrorMessage(), "not available in the current Region") {
			return true
		}
	}
	return false
}

func IgnoreWithInvalidAction(err error) bool {
	var ae smithy.APIError
	if errors.As(err, &ae) {
		if ae.ErrorCode() == "InvalidAction" {
			return true
		}
	}
	return false
}

func IsErrorRegex(err error, code string, messageRegex *regexp.Regexp) bool {
	var ae smithy.APIError
	if !errors.As(err, &ae) {
		return false
	}
	if ae.ErrorCode() == code && messageRegex.MatchString(ae.ErrorMessage()) {
		return true
	}
	return false
}

func IsInvalidParameterValueError(err error) bool {
	var apiErr smithy.APIError
	return errors.As(err, &apiErr) && apiErr.ErrorCode() == "InvalidParameterValue"
}

func makeARN(service string, partition, accountID, region string, idParts ...string) arn.ARN {
	return arn.ARN{
		Partition:	partition,
		Service:	string(service),
		Region:		region,
		AccountID:	accountID,
		Resource:	strings.Join(idParts, "/"),
	}
}

type AwsService struct {
	Regions map[string]*map[string]interface{} `json:"regions"`
}

type SupportedServiceRegionsData struct {
	Partitions		map[string]AwsPartition	`json:"partitions"`
	regionVsPartition	map[string]string
}

const MAX_GOROUTINES = 10

const (
	PartitionServiceRegionFile	= "data/partition_service_region.json"
	defaultPartition		= "aws"
)

var (
	readOnce		sync.Once
	supportedServiceRegion	*SupportedServiceRegionsData
)

func IgnoreCommonErrors(err error) bool {
	if IgnoreAccessDeniedServiceDisabled(err) || IgnoreNotAvailableRegion(err) || IgnoreWithInvalidAction(err) || isNotFoundError(err) {
		return true
	}
	return false
}

func TagsIntoMap(tagSlice interface{}, dst map[string]string) {
	stringify := func(v reflect.Value) string {
		vt := v.Type()
		if vt.Kind() == reflect.String {
			return v.String()
		}
		if vt.Kind() != reflect.Ptr || vt.Elem().Kind() != reflect.String {
			panic("field is not string or *string")
		}

		if v.IsNil() {

			return ""
		}

		return v.Elem().String()
	}

	if k := reflect.TypeOf(tagSlice).Kind(); k != reflect.Slice {
		panic("invalid usage: Only slices are supported as input: " + k.String())
	}
	slc := reflect.ValueOf(tagSlice)

	for i := 0; i < slc.Len(); i++ {
		val := slc.Index(i)
		if k := val.Kind(); k != reflect.Struct {
			panic("slice member is not struct: " + k.String())
		}

		keyField, valField := val.FieldByName("Key"), val.FieldByName("Value")
		if keyField.Type().Kind() == reflect.Ptr && keyField.IsNil() {
			continue
		}

		if keyField.IsZero() {
			panic("slice member is missing Key field")
		}

		dst[stringify(keyField)] = stringify(valField)
	}
}

func Sleep(ctx context.Context, dur time.Duration) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(dur):
		return nil
	}
}
