package converter

import (
	"errors"
	"reflect"
)

func Convert(to any, from any) error {
	//
	pointerToResult := reflect.ValueOf(to)
	reflectedValueOfResult := pointerToResult.Elem()
	if reflectedValueOfResult.Kind() != reflect.Struct {
		return errors.New("To is not a struct")
	}
	valueOfData := reflect.ValueOf(from)
	if valueOfData.Kind() != reflect.Struct {
		return errors.New("From is not a struct")
	}
	convert(reflectedValueOfResult, valueOfData)
	return nil
}

func convert(to reflect.Value, from reflect.Value) {
	typeOfData := from.Type()
	for i := 0; i < from.NumField(); i++ {
		resultField := to.FieldByName(typeOfData.Field(i).Name)
		fromField := from.Field(i)
		setValue(resultField, fromField)
	}
}

func setValue(setTo reflect.Value, value reflect.Value) bool {
	if setTo.IsValid() && setTo.CanSet() {
		if setTo.Kind() == value.Kind() {
			return valueSeter(setTo, value)
		} else if setTo.Kind() == reflect.Pointer {
			setToFieldType := setTo.Type().Elem()
			if setToFieldType.Kind() == value.Kind() {
				pointerValue := reflect.New(setToFieldType).Elem()
				valueSeter(pointerValue, value)
				setTo.Set(pointerValue.Addr())
				return true
			}
		} else if value.Kind() == reflect.Pointer {
			pointerValue := value.Elem()
			if setTo.Kind() == pointerValue.Kind() {
				return valueSeter(setTo, pointerValue)
			}
		}
	}
	return false
}

func valueSeter(setTo reflect.Value, value reflect.Value) bool {
	if setTo.Type() == value.Type() {
		setTo.Set(value)
		return true
	}
	if setTo.Kind() == reflect.Struct {
		convert(setTo, value)
		return true
	} else if setTo.Kind() == reflect.Array || setTo.Kind() == reflect.Slice {
		setToFieldType := setTo.Type().Elem()
		arrayValue := reflect.New(setToFieldType).Elem()
		for i := 0; i < value.Len(); i++ {
			item := value.Index(i)
			result := setValue(arrayValue, item)
			if result {
				setTo.Set(reflect.Append(setTo, arrayValue))
			}
		}
		return true
	} else {
		setToFieldType := setTo.Type().Elem()
		pointerValue := reflect.New(setToFieldType).Elem()
		valueSeter(pointerValue, value.Elem())
		setTo.Set(pointerValue.Addr())
		return true
	}
}
