package handler

import (
	"fmt"
	"reflect"

	"github.com/Tomoki108/burny/handler/io"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/labstack/echo/v4"
)

var Validator *validator.Validate
var Trans ut.Translator

func init() {
	en := en.New()
	uni := ut.New(en, en)

	Trans, _ = uni.GetTranslator("en")
	Validator = validator.New()
	Validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		label := fld.Tag.Get("json")
		if label == "" {
			return fld.Name
		}
		return label
	})
	en_translations.RegisterDefaultTranslations(Validator, Trans)
}

// handleReq binds and validates request
func handleReq[T any](c echo.Context, req *T) *io.ErrorResponse {
	// bind
	if err := c.Bind(req); err != nil {
		resp := io.NewErrResp(fmt.Sprintf("Failed to bind request: %s", err.Error()))
		return &resp
	}

	// validate
	if err := Validator.Struct(req); err != nil {
		resp := io.NewErrResp("Validation error")
		validationErrors := err.(validator.ValidationErrors)
		for _, e := range validationErrors {
			resp.Details = append(resp.Details, io.ErrorDetail{
				Field:   e.Field(),
				Message: e.Translate(Trans),
			})
		}

		return &resp
	}
	return nil
}
