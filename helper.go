package datacache

import (
	"fmt"
	"reflect"
)

func InspectStruct(v interface{}) {
	inspectStructV(reflect.ValueOf(v))
}

func inspectStructV(val reflect.Value) {
	if val.Kind() == reflect.Interface && !val.IsNil() {
		elm := val.Elem()
		if elm.Kind() == reflect.Ptr && !elm.IsNil() && elm.Elem().Kind() == reflect.Ptr {
			val = elm
		}
	}
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		fmt.Println("####  val:", val)
	}

	//if val.Elem().Elem().Kind() == reflect.Struct {
	//if reflect.TypeOf(val).Kind() == reflect.Struct {
	if val.Kind() == reflect.Struct {
		for i := 0; i < val.NumField(); i++ {
			valueField := val.Field(i)
			typeField := val.Type().Field(i)
			address := "not-addressable"

			if valueField.Kind() == reflect.Interface && !valueField.IsNil() {
				elm := valueField.Elem()
				if elm.Kind() == reflect.Ptr && !elm.IsNil() && elm.Elem().Kind() == reflect.Ptr {
					valueField = elm
				}
			}

			if valueField.Kind() == reflect.Ptr {
				valueField = valueField.Elem()
			}
			if valueField.CanAddr() {
				address = fmt.Sprintf("0x%X", valueField.Addr().Pointer())
			}

			fmt.Printf("Field Name: %s,\t Field Value: %v,\t Address: %v\t, Field type: %v\t, Field kind: %v\n", typeField.Name,
				valueField.Interface(), address, typeField.Type, valueField.Kind())

			if valueField.Kind() == reflect.Struct {
				inspectStructV(valueField)
			}
		}
	}
}
