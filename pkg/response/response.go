package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gopi/internal/helper"
	"gopi/internal/logger"
	apperrors "gopi/pkg/app_errors"

	"go.uber.org/zap"
)

type Response struct {
	ctx    echo.Context
	logger *zap.Logger
}

type Data struct {
	Ok      bool   `json:"ok"`
	Code    int    `json:"status_code"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
	Errors  any    `json:"errors,omitempty"`
	Data    any    `json:"data,omitempty"`
	Meta    any    `json:"meta,omitempty"`
	Stack   error  `json:"-"`
}

func New(c echo.Context) *Response {
	return &Response{
		ctx:    c,
		logger: logger.Logger,
	}
}

func (r *Response) Success(data Data) error {
	data.Ok = true

	if data.Code == 0 {
		data.Code = 200
	}

	data.Message = helper.CapitalizeFirst(data.Message)

	// reset all error fields in case they are accidentally present
	data.Error = ""
	data.Stack = nil
	data.Errors = nil

	return r.ctx.JSON(data.Code, data)
}

func (r *Response) Error(data Data) error {
	data.Ok = false

	if data.Error == "" {
		switch data.Code {
		case http.StatusBadRequest:
			data.Error = apperrors.ErrBadRequest

		case http.StatusPaymentRequired:
			data.Error = apperrors.ErrUpgradeRequired

		case http.StatusUnauthorized:
			data.Error = apperrors.ErrUnauthorized

		case http.StatusForbidden:
			data.Error = apperrors.ErrForbidden

		default:
			data.Error = "Internal Server Error"
		}
	}

	if data.Code == http.StatusBadRequest && data.Errors != nil {
		data.Error = apperrors.ErrValidation
	}

	if data.Message != "" {
		data.Error = data.Message
		data.Message = ""
	}

	if data.Code == 0 {
		data.Code = 500
	}

	if data.Stack != nil {
		if _, ok := data.Stack.(*apperrors.CustomError); ok {
			data.Error = data.Stack.Error()
		}
		if r.logger != nil {
			r.logger.Error(data.Stack.Error())
		}
	}

	data.Error = helper.CapitalizeFirst(data.Error)

	return r.ctx.JSON(data.Code, data)
}
