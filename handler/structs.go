package handler

import (
	"errors"
	"strings"
	"time"

	"github.com/pjserol/api-rest-user/common/utils"
	"github.com/pjserol/api-rest-user/config"
	"github.com/pjserol/api-rest-user/model"
)

// ContextKey is used for context.Context value. The value requires a key that is not primitive type.
type ContextKey string

// ContextKeyRequestID is the ContextKey for RequestID
const ContextKeyRequestID ContextKey = "requestID"

type errorResponse struct {
	Success  bool      `json:"success"`
	Messages []string  `json:"messages"`
	Time     time.Time `json:"time"`
	Timing   []timing  `json:"timing"`
}

type healthCheckResponse struct {
	Success           bool               `json:"success"`
	Messages          []string           `json:"messages"`
	Time              time.Time          `json:"time"`
	Timing            []timing           `json:"timing"`
	Environment       config.Environment `json:"environment"`
	DatabaseConnected bool               `json:"databaseConnected"`
}

type timing struct {
	TimeMillis int64  `json:"timeMillis"`
	Source     string `json:"source"`
}

type userResponse struct {
	Success  bool       `json:"success"`
	Messages []string   `json:"messages"`
	Time     time.Time  `json:"time"`
	Timing   []timing   `json:"timing"`
	Response model.User `json:"response"`
}

type userRequest struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (u *userRequest) validate() ([]string, error) {
	errorDetails := make([]string, 0)

	if !utils.ValidateEmail(u.Email) {
		errorDetails = append(errorDetails, "the field 'email'is not valid")
	}

	if strings.TrimSpace(u.FirstName) == "" {
		errorDetails = append(errorDetails, "the field 'firstName' cannot be empty")
	}

	if strings.TrimSpace(u.LastName) == "" {
		errorDetails = append(errorDetails, "the field 'lastName' cannot be empty")
	}

	if len(errorDetails) > 0 {
		return errorDetails, errors.New("ill-formed user")
	}

	return errorDetails, nil
}

type getUserRequest struct {
	ID string `json:"id"`
}
