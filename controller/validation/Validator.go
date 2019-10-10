package validation

import (
	"gopkg.in/go-playground/validator.v8"
	"reflect"
	"regexp"
	"spm-serv/core"
	"strings"
)

//注册验证器
func Register(v *validator.Validate, name string, fn validator.Func){
	err := v.RegisterValidation(name, fn)
	if err!=nil {
		core.Log.Panicln(err)
	}
}




func CheckVersion(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,) bool{
		if version, ok := field.Interface().(string); ok {
			reg := regexp.MustCompile("^(latest)|(\\d+\\.\\d+\\.\\d+)$")
			return reg.MatchString(version)
		}
		return false
}

//检查是否在指定的多个值中，示例：[1,2,3]
func CheckValues(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,) bool{
		if param=="" {
			return false
		}
		param = strings.TrimSpace(param)
		if len(param) <= 2 {
			return false
		}
		if !(strings.HasPrefix(param, "[") && strings.HasSuffix(param, "]")) {
			return false
		}
		values := strings.Split(param[1:len(param)-1], ",")
		if len(values)==0 {
			return false
		}
		if val, ok := field.Interface().(string); ok {
			for _, v := range values {
				if strings.TrimSpace(v)==val {
					return true
				}
			}
		}
		return false
}
