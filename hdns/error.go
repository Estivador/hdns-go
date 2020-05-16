package hdns

import "fmt"

// ErrorCode represents an error code returned from the API.
type ErrorCode string

// Error codes returned from the API.
const (
	ErrorCodePaginationError     ErrorCode = "400" // Pagination selectors are mutually exclusive
	ErrorCodeUnauthorized        ErrorCode = "401" // Unathorized
	ErrorCodeForbidden           ErrorCode = "403" // Insufficient permissions
	ErrorCodeNotFound            ErrorCode = "404" // Resource not found
	ErrorCodeNotAcceptable       ErrorCode = "406" // Not acceptable
	ErrorCodeConflict            ErrorCode = "409" // Conflict
	ErrorCodeUnprocessableEntity ErrorCode = "422" // Unprocessable entity
)

// Error is an error returned from the API.
type Error struct {
	Code    ErrorCode
	Message string
}

func (e Error) Error() string {
	return fmt.Sprintf("%s (%s)", e.Message, e.Code)
}

// IsError returns whether err is an API error with the given error code.
func IsError(err error, code ErrorCode) bool {
	apiErr, ok := err.(Error)
	return ok && apiErr.Code == code
}
