package httputil

import (
	"net/http"

	"github.com/rapidashorg/temp-shared-libraries/errors"
)

var (
	errInvalidRequestBody = errors.NewMaskedError(101, "ErrInvalidRequestBody", intptr(http.StatusBadRequest), nil)
	errInternalServer     = errors.NewMaskedError(102, "ErrInternalServer", nil, nil)
)

// intptr convert int to int pointer.
func intptr(val int) *int {
	return &val
}

// intptr convert string to string pointer.
func strptr(val string) *string {
	return &val
}
