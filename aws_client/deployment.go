package aws_client

import (
	"context"
	"strings"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

type DataSource func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics

type WarpDataSource func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) error

type ManyError struct {
	errs []error
}

func (m ManyError) Error() string {

	s := strings.Builder{}

	for i := len(m.errs) - 1; i > 0; i-- {
		s.WriteString(m.errs[i].Error())
	}
	return s.String()
}

func NewManyError(errs []error) error {
	return ManyError{errs: errs}
}

func listBucketRegion(partition string) string {
	switch partition {
	case "aws-cn":
		return "cn-north-1"
	case "aws-us-gov":
		return "us-gov-west-1"
	default:
		return "us-east-1"
	}
}

func FilterError(fn WarpDataSource) DataSource {
	return func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
		err := fn(ctx, clientMeta, client, task, resultChannel)
		if !ignoreCommonErrors(err) {
			return schema.NewDiagnostics().AddErrorPullTable(task.Table, err)
		}
		return nil
	}
}

func ignoreCommonErrors(err error) bool {
	if err == nil {
		return true
	}
	switch e := err.(type) {
	case ManyError:
		for _, err2 := range e.errs {
			if !ignoreCommonErrors(err2) {
				return false
			}
		}
		return true
	default:
		if IgnoreAccessDeniedServiceDisabled(err) || IgnoreNotAvailableRegion(err) || IgnoreWithInvalidAction(err) || isNotFoundError(err) {
			return true
		}
		return false
	}
}
