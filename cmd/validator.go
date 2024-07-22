package cmd

import (
	"xwya/utils"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

func InitValidate() {
	utils.Validate = validator.New()
	uni := ut.New(zh.New())
	utils.Trans, _ = uni.GetTranslator("zh")

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册中文翻译
		_ = zh_translations.RegisterDefaultTranslations(v, utils.Trans)

		// 自定义错误消息模板
		v.RegisterTranslation("required", utils.Trans, func(ut ut.Translator) error {
			return ut.Add("required", "{0} 是必填字段", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("required", fe.Field())
			return t
		})
		// 自定义邮箱错误模板
		v.RegisterTranslation("email", utils.Trans, func(ut ut.Translator) error {
			return ut.Add("email", "{0} 不是有效的邮箱地址", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("email", fe.Field())
			return t
		})
		// 自定义最小和最大长度错误模板
		v.RegisterTranslation("min", utils.Trans, func(ut ut.Translator) error {
			return ut.Add("min", "{0} 的长度不能小于 {1}", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("min", fe.Field(), fe.Param())
			return t
		})

		v.RegisterTranslation("max", utils.Trans, func(ut ut.Translator) error {
			return ut.Add("max", "{0} 的长度不能大于 {1}", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("max", fe.Field(), fe.Param())
			return t
		})
		v.RegisterTranslation("alphanum", utils.Trans, func(ut ut.Translator) error {
			return ut.Add("alphanum", "{0} 只能包含字母和数字", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("max", fe.Field())
			return t
		})
	}
}
