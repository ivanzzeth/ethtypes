package ethtypes

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"gorm.io/gorm/schema"
)

var _ schema.SerializerInterface = GormSlice[*Address]{}
var _ fmt.Stringer = GormSlice[*Address]{}
var _ JsonObj = (*GormSlice[*Address])(nil)

type GormSlice[T JsonObj] []T

func NewGormSlice[T JsonObj](s []T) *GormSlice[T] {
	slice := GormSlice[T](s)
	return &slice
}

func (s GormSlice[T]) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue interface{}) (err error) {
	switch value := dbValue.(type) {
	case string:
		return s.UnmarshalJSON([]byte(value))
	default:
		return fmt.Errorf("unsupported data %#v", dbValue)
	}
}

func (s GormSlice[T]) Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue interface{}) (interface{}, error) {
	data, err := s.MarshalJSON()
	return string(data), err
}

func (s GormSlice[T]) MarshalJSON() ([]byte, error) {
	arr := []string{}
	for _, elem := range s {
		data, err := elem.MarshalJSON()
		if err != nil {
			return nil, err
		}
		arr = append(arr, string(data))
	}

	data, err := json.Marshal(arr)
	if err != nil {
		return nil, err
	}

	fmt.Println("MarshalJSON data", string(data))

	return data, nil
}

func (s *GormSlice[T]) UnmarshalJSON(data []byte) error {
	arr := []string{}
	err := json.Unmarshal(data, &arr)
	if err != nil {
		return err
	}

	newSlice := make(GormSlice[T], len(arr))

	for i, data := range arr {
		elem := newSlice[i]
		elemTypePointer := reflect.TypeOf(elem)
		elemType := elemTypePointer.Elem()
		// fmt.Println("elemTypePointer", elemTypePointer, "elemType", elemType)
		elem = reflect.New(elemType).Interface().(T)
		newSlice[i] = elem

		err = elem.UnmarshalJSON([]byte(data))
		if err != nil {
			return err
		}
		// fmt.Println("elem ", i, elem, newSlice[i], reflect.TypeOf(newSlice[i]))
	}

	// for i, elem := range newSlice {
	// 	fmt.Println("slice ", i, elem)
	// }

	*s = newSlice

	// for i, elem := range *s {
	// 	fmt.Println("raw slice ", i, elem)
	// }
	fmt.Println("Unmarshal data", s)

	return nil
}

func (s GormSlice[T]) String() string {
	js, err := s.MarshalJSON()
	if err != nil {
		panic(err)
	}
	return string(js)
}

func (s GormSlice[T]) Unwrap() []T {
	return s
}
