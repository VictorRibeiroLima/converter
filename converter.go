package converter

import (
	"errors"
	"reflect"
)

func Convert(to any, from any) error {
	pointerToResult := reflect.ValueOf(to)
	reflectedValueOfResult := pointerToResult.Elem()
	toKind := reflectedValueOfResult.Kind()
	valueOfData := reflect.ValueOf(from)
	fromKind := valueOfData.Kind()
	if toKind != reflect.Struct && toKind != reflect.Array && toKind != reflect.Slice {
		return errors.New("to is not conversable")
	}
	if fromKind != reflect.Struct && fromKind != reflect.Array && fromKind != reflect.Slice {
		return errors.New("from is not conversable")
	}
	if toKind == reflect.Struct {
		if fromKind != reflect.Struct {
			return errors.New("incompatible types")
		}
		convert(reflectedValueOfResult, valueOfData)
	} else {
		if fromKind == reflect.Struct {
			return errors.New("incompatible types")
		}
		convertArray(reflectedValueOfResult, valueOfData)
	}
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
			return valueSetter(setTo, value)
		} else if setTo.Kind() == reflect.Pointer {
			setToFieldType := setTo.Type().Elem()
			if setToFieldType.Kind() == value.Kind() {
				pointerValue := reflect.New(setToFieldType).Elem()
				valueSetter(pointerValue, value)
				setTo.Set(pointerValue.Addr())
				return true
			}
		} else if value.Kind() == reflect.Pointer {
			pointerValue := value.Elem()
			if setTo.Kind() == pointerValue.Kind() {
				return valueSetter(setTo, pointerValue)
			}
		}
	}
	return false
}

func valueSetter(setTo reflect.Value, value reflect.Value) bool {
	if setTo.Type() == value.Type() {
		setTo.Set(value)
		return true
	}
	if setTo.Kind() == reflect.Struct {
		convert(setTo, value)
		return true
	} else if setTo.Kind() == reflect.Array || setTo.Kind() == reflect.Slice {
		return convertArray(setTo, value)
	} else if value.Kind() == reflect.Pointer {
		if value.IsNil() {
			return true
		}
		setToFieldType := setTo.Type().Elem()
		pointerValue := reflect.New(setToFieldType).Elem()
		valueSetter(pointerValue, value.Elem())
		setTo.Set(pointerValue.Addr())
		return true
	}
	return false
}

func convertArray(setTo reflect.Value, value reflect.Value) bool {
	setToFieldType := setTo.Type().Elem()
	for i := 0; i < value.Len(); i++ {
		arrayValue := reflect.New(setToFieldType).Elem()
		item := value.Index(i)
		result := setValue(arrayValue, item)
		if result {
			setTo.Set(reflect.Append(setTo, arrayValue))
		}
	}
	return true
}
