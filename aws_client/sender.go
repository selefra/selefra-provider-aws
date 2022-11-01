package aws_client

import (
	"reflect"
)

func SendResults(resultChannel chan<- any, results any, resultProcessFunc func(result any) (any, error)) {
	reflectValue := reflect.ValueOf(results)
	switch reflectValue.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < reflectValue.Len(); i++ {
			indexValue := reflectValue.Index(i)
			newResult, err := resultProcessFunc(indexValue.Interface())
			if err != nil {
				continue
			}
			resultChannel <- newResult
		}
	default:
		newResult, err := resultProcessFunc(results)
		if err != nil {
			return
		}
		resultChannel <- newResult
	}
}
