package aws_client

import (
	"context"
	"encoding/json"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-utils/pkg/reflect_util"
	"github.com/songzhibin97/go-ognl"
	"reflect"
)

func AwsAccountIDExtractor() schema.ColumnValueExtractor {
	return column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
		return client.(*Client).GetAccount(), nil
	})
}

func AwsRegionIDExtractor() schema.ColumnValueExtractor {
	return column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
		return client.(*Client).GetRegion(), nil
	})
}

func TagsExtractor(tagsFieldName ...string) schema.ColumnValueExtractor {
	if len(tagsFieldName) == 0 {
		tagsFieldName = append(tagsFieldName, "Tags")
	}
	return column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
		for _, tagName := range tagsFieldName {
			value := ognl.Get(result, "."+tagName)
			if !reflect_util.IsNil(value) {
				return TagsToMap(value), nil
			}
		}
		return nil, nil
	})
}

func TagsToMap(value interface{}) map[string]string {
	mp := map[string]string{}
	tv := reflect.ValueOf(value)

	if !tv.IsValid() {
		return mp
	}

	switch tv.Kind() {
	case reflect.Ptr, reflect.Interface:
		return TagsToMap(tv.Elem())
	case reflect.Slice, reflect.Array:
		for i := 0; i < tv.Len(); i++ {
			if !tv.Index(i).IsValid() {
				continue
			}
			switch tv.Index(i).Kind() {
			case reflect.Struct:
				if tv.Index(i).FieldByName("Key").IsValid() && tv.Index(i).FieldByName("Value").IsValid() {
					mp[getTagValue(tv.Index(i).FieldByName("Key"))] = getTagValue(tv.Index(i).FieldByName("Value"))
				}
			}
		}
	}
	return mp
}

func getTagValue(tv reflect.Value) string {
	if !tv.IsValid() {
		return ""
	}

	switch tv.Kind() {
	case reflect.String:
		return tv.String()
	case reflect.Ptr, reflect.Interface:
		return getTagValue(tv.Elem())
	default:
		return ""
	}
}

func WAFScopeExtractor() schema.ColumnValueExtractor {
	return column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
		return client.(*Client).WAFScope, nil
	})
}

func MarshaledJsonExtractor(path string) schema.ColumnValueExtractor {
	return column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
		var j map[string]interface{}

		value := ognl.Get(result, path).Value()

		if value == nil {
			return nil, nil
		}

		var val reflect.Value
		val = reflect.ValueOf(value)
		if reflect.TypeOf(value).Kind() == reflect.Ptr {
			val = val.Elem()
		}

		var b []byte
		switch val.Kind() {
		case reflect.String:
			b = []byte(val.String())
		case reflect.Slice, reflect.Array:
			b = val.Bytes()
		}
		if len(b) == 0 {
			return nil, nil
		}

		j = make(map[string]interface{})
		err := json.Unmarshal(b, &j)
		if err != nil {
			return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
		}

		return j, nil
	})
}

func (c *Client) RegionGlobalARN(service string, idParts ...string) string {
	return makeARN(service, c.Partition, "", c.Region, idParts...).String()
}

func (c *Client) PartitionGlobalARN(service string, idParts ...string) string {
	return makeARN(service, c.Partition, "", "", idParts...).String()
}
