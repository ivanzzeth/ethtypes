package ethtypes

import (
	"encoding/json"
	"fmt"
)

func ToString[T fmt.Stringer](data []T) string {
	strs := []string{}
	for _, v := range data {
		strs = append(strs, v.String())
	}

	res, _ := json.Marshal(strs)

	return string(res)
}
