package utils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var trans ut.Translator

func init() {
	InitTrans("zh")
	fmt.Println("注册翻译器完毕")
}

func InitTrans(locale string) (err error) {
	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validate.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New()          // 中文翻译
		uniT := ut.New(zhT, zhT) //注册中文翻译机

		//enT := en.New() // 英文翻译
		//uniT := ut.New(zhT, enT) //注册英文翻译机

		// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
		trans, ok = uniT.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}
		// 注册翻译器
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(validate, trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(validate, trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(validate, trans)
		}
		return
	}
	return
}
func ParamError(err validator.ValidationErrorsTranslations) map[string]string {
	errMsgs := make(map[string]string)
	for k, v := range err {
		simplifiedFieldName := formatFieldName(k)
		errMsgs[simplifiedFieldName] = v
	}
	return errMsgs
}
func formatFieldName(fieldName string) string {
	// 这里可以根据需求修改字段名称的格式化逻辑
	parts := strings.Split(fieldName, ".")
	return parts[len(parts)-1] // 只取最后一个部分
}
