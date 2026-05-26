package core_http_request

import (
	"encoding/json"
	"fmt"
	"net/http"

	core_errors "github.com/glebateee/taskapp/internal/core/errors"
	"github.com/go-playground/validator/v10"
)

var requestValidator = validator.New()

type validatable interface {
	Validate() error
}

func DecodeAndValidate(r *http.Request, dto any) error {
	if err := json.NewDecoder(r.Body).Decode(dto); err != nil {
		return fmt.Errorf(
			"decode json: %v: %w",
			err,
			core_errors.ErrInvalidArgument,
		)
	}
	var err error
	if v, ok := dto.(validatable); ok {
		err = v.Validate()
	} else {
		err = requestValidator.Struct(dto)
	}
	if err != nil {
		return fmt.Errorf(
			"request validation: %v: %w",
			err,
			core_errors.ErrInvalidArgument,
		)
	}
	return nil
}
