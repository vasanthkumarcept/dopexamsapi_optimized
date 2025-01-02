// errors.go
package apierrors

import (
	"time"
)

type ErrorDetails struct {
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
	Details   []string  `json:"details"`
}

func NewErrorDetails(message string, details []string) *ErrorDetails {
	return &ErrorDetails{
		Timestamp: time.Now(),
		Message:   message,
		Details:   details,
	}
}

type APIError struct {
	Status  int             `json:"status"`
	Message string          `json:"message"`
	Details []*ErrorDetails `json:"details"`
}

func NewAPIError(status int, message string, details []*ErrorDetails) *APIError {
	return &APIError{
		Status:  status,
		Message: message,
		Details: details,
	}
}
