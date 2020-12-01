package core

import (
	"reflect"
)

// Serialize Method for serializing instance fields according to provided serializer
func Serialize(serializerClass interface{}, instance interface{}) interface{} {
	t := reflect.TypeOf(serializerClass)
	serializerValue := reflect.New(t).Elem()
	elemValue := reflect.ValueOf(instance).Elem()
	for i := 0; i < t.NumField(); i++ {
		fieldName := t.Field(i).Name
		val := elemValue.FieldByName(fieldName)
		if !val.IsValid() {
			continue
		}
		serializerValue.Field(i).Set(val)
	}
	return serializerValue.Interface()
}
