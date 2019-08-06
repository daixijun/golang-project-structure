package validator

import (
	"reflect"
	"sync"

	"github.com/gin-gonic/gin/binding"
	zh_locale "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
)

type DefaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

var _ binding.StructValidator = &DefaultValidator{}

var (
	uni   *ut.UniversalTranslator
	trans ut.Translator
)

func (v *DefaultValidator) ValidateStruct(obj interface{}) error {

	if kindOfData(obj) == reflect.Struct {

		v.lazyinit()

		if err := v.validate.Struct(obj); err != nil {
			return error(err)
		}
	}

	return nil
}

func (v *DefaultValidator) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

func (v *DefaultValidator) lazyinit() {
	v.once.Do(func() {
		zh := zh_locale.New()
		uni = ut.New(zh, zh)
		trans, _ = uni.GetTranslator("zh")
		v.validate = validator.New()
		v.validate.SetTagName("binding")
		v.validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			fieldName := fld.Tag.Get("json")
			if fieldName == "-" {
				return ""
			}
			return fieldName
		})
		// add any custom validations etc. here
		_ = zh_translations.RegisterDefaultTranslations(v.validate, trans)
	})
}

func kindOfData(data interface{}) reflect.Kind {

	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}

func GetErrorMessages(err error) map[string]string {
	messages := make(map[string]string)
	if err == nil {
		return messages
	}
	for _, err := range err.(validator.ValidationErrors) {
		messages[err.Field()] = err.Translate(trans)
	}

	return messages
}
