package logs

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
)

type jsonLog struct {
	Timestamp int64  `json:"timestamp"`
	Message   string `json:"message"`
	RequestID string `json:"requestId"`
}

// ContextKey is used for context.Context value. The value requires a key that is not primitive type.
type ContextKey string // can be unexported

// ContextKeyRequestID is the ContextKey for RequestID
const ContextKeyRequestID ContextKey = "requestID" // can be unexported

// AssignRequestID will attach a brand new request ID to a http request
func AssignRequestID(ctx context.Context) context.Context {

	reqID := uuid.New()

	return context.WithValue(ctx, ContextKeyRequestID, reqID.String())
}

// GetRequestID will get reqID from a http request and return it as a string
func GetRequestID(ctx context.Context) string {

	reqID := ctx.Value(ContextKeyRequestID)

	if id, ok := reqID.(string); ok {
		return id
	}

	return ""
}

// Log - log in JSON along with contextual information
func Log(ctx context.Context, message string) {
	if os.Getenv("SUPPRESS_LOGS") == "true" {
		return
	}

	logObject := jsonLog{
		Timestamp: time.Now().Unix(),
		Message:   message,
		RequestID: GetRequestID(ctx),
	}
	b, _ := json.Marshal(logObject)
	fmt.Println(string(b))
}
