// Package model holds the business types shared by every layer. It imports nothing but the standard
// library and internal/constant.
package model

import (
	"fmt"
	"strings"
)

// Error codes crossing the service boundary. Whoever exposes the service (HTTP handler, CLI) owns the
// mapping from these to its own status codes; the service layer never knows what an HTTP status is.
const (
	ErrBadRequest = "ERR_BAD_REQUEST" // caller sent a message the FIN format cannot carry
	ErrInternal   = "ERR_INTERNAL"    // our fault
)

// CustomError is the only error type allowed to leave the service layer.
type CustomError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Error implements the error interface.
func (e *CustomError) Error() (message string) {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// NewError builds a CustomError. Several messages are joined with ":" - pass the summary first and
// the cause after it.
func NewError(code string, messages ...string) (err *CustomError) {
	return &CustomError{Code: code, Message: strings.Join(messages, ":")}
}
