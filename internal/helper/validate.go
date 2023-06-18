package helper

import (
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	uni        *ut.UniversalTranslator
	validate   *validator.Validate
	translator ut.Translator
)

func init() {
	en := en.New()
	uni = ut.New(en, en)

	trans, found := uni.GetTranslator("en")
	if !found {
		panic("english translation not found!")
	}
	translator = trans

	validate = validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		splitTag := strings.SplitN(fld.Tag.Get("json"), ",", 2)
		if len(splitTag) == 0 {
			return ""
		}
		name := splitTag[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

func Validate(value any) (map[string]string, error) {
	validationErrors := make(map[string]string)
	en_translations.RegisterDefaultTranslations(validate, translator)

	err := validate.Struct(value)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return validationErrors, err
		}

		for _, fieldErr := range err.(validator.ValidationErrors) {
			fieldName := strings.ToLower(fieldErr.Field())
			validationErrors[fieldName] = strings.Replace(strings.ToLower(fieldErr.Translate(translator)), fieldName+" ", "", 1)
		}
	}

	if len(validationErrors) == 0 {
		validationErrors = nil
	}

	return validationErrors, nil
}
