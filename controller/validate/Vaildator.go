package vaildate

import (
	"gopkg.in/go-playground/validator.v8"
	"reflect"
	"regexp"
)

func CheckVersion(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,) bool{
		if version, ok := field.Interface().(string); ok {
			reg := regexp.MustCompile("^\\d+\\.\\d+\\.\\d+$")
			return reg.MatchString(version)
		}
		return false
}
