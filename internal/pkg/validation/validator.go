package validation

import (
	"errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	chTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var trans ut.Translator

/**
 * ValidatorLocalInit 初始化语言包
 * local使用的语音zh|en
 */
func ValidatorLocalInit(local string) error {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhT := zh.New() //chinese
		enT := en.New() //english
		uni := ut.New(enT, zhT, enT)

		trans, ok = uni.GetTranslator(local)
		if !ok {
			return errors.New("uni.GetTranslator failed")
		}
		// 注册翻译器
		var err error
		switch local {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		case "zh":
			err = chTranslations.RegisterDefaultTranslations(v, trans)
		default:
			err = chTranslations.RegisterDefaultTranslations(v, trans)
		}
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

// 错误翻译
func Translate(err error) (errMsg string) {
	errs := err.(validator.ValidationErrors)
	for _, err := range errs {
		errMsg = err.Translate(trans)
		break
	}
	return
}
