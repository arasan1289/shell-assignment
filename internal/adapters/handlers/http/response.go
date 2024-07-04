package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// response represents a response body format
type response struct {
	Success bool              `json:"success"`
	Data    any               `json:"data,omitempty"`
	Errors  map[string]string `json:"errors,omitempty"`
}

// newResponse is a helper function to create a response body
func newResponse(success bool, data any, errors map[string]string) response {
	return response{
		Success: success,
		Data:    data,
		Errors:  errors,
	}
}

// parseError parses error messages from the error object and returns a map of error messages with field name
func parseError(ctx *gin.Context, err error) map[string]string {
	var errMsgs = make(map[string]string)

	// Check if the error is a validation error
	var validationErr validator.ValidationErrors
	if errors.As(err, &validationErr) {
		// Log each validation error
		for _, err := range validationErr {
			ctx.Error(err)
		}
		// Generate error messages based on error type
		jsonFormatter := NewJSONFormatter()
		errMsgs = jsonFormatter.Simple(validationErr)
	} else {
		// Log the error
		ctx.Error(err)
		// Add a generic error message
		errMsgs["message"] = err.Error()
	}

	return errMsgs
}

// handleSuccess sends a success response with the specified status code and optional data
func handleSuccess(ctx *gin.Context, data any) {
	rsp := newResponse(true, data, nil)
	ctx.JSON(http.StatusOK, rsp)
}

// handleError sends a error response with the specified status code and error message
func handleError(ctx *gin.Context, err error) {
	statusCode := http.StatusInternalServerError
	errMsg := parseError(ctx, err)
	errRsp := newResponse(false, nil, errMsg)
	ctx.JSON(statusCode, errRsp)
}

// validationError sends a error response with the specified status code and error message
func validationError(ctx *gin.Context, err error) {
	errs := parseError(ctx, err)
	errRsp := newResponse(false, nil, errs)
	ctx.JSON(http.StatusBadRequest, errRsp)
}
