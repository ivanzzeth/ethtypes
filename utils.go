package ethtypes

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func ToString(data interface{}) string {
	strs := []string{}

	switch data := data.(type) {
	case string:
		return data
	case []string:
		strs = data
	default:
		val := reflect.ValueOf(data)
		switch val.Kind() {
		case reflect.Slice:
			for i := 0; i < val.Len(); i++ {
				elem := val.Index(i)
				if elem.Type().Implements(reflect.TypeOf((*fmt.Stringer)(nil)).Elem()) {
					e := elem.Interface()
					strs = append(strs, e.(fmt.Stringer).String())
				}
			}
		}
		// for _, d := range data {
		// 	strs = append(strs, d.String())
		// }
	}

	res, _ := json.Marshal(strs)

	return string(res)
}
