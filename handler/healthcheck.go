package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/pjserol/api-rest-user/common/logs"
	"github.com/pjserol/api-rest-user/common/utils"
	"github.com/pjserol/api-rest-user/config"
	"github.com/pjserol/api-rest-user/db"
)

const (
	healthCheckHandler = "HealthCheckHandler"
)

// HealthCheckHandler report the health of the application
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	start := utils.MakeTimestampMilli()
	w.Header().Set("Content-Type", jsonContentType)
	w.Header().Set("Access-Control-Allow-Origin", "*")

	messages := make([]string, 0)

	mem := utils.MemoryUsage()
	messages = append(messages, mem)

	env := config.InitEnvironment()

	t, _ := json.MarshalIndent(healthCheckResponse{
		Success:  true,
		Messages: messages,
		Time:     time.Now().UTC(),
		Timing: []timing{
			{
				Source:     healthCheckHandler,
				TimeMillis: utils.MakeTimestampMilli() - start,
			},
		},
		Environment:       env,
		DatabaseConnected: db.Connected,
	}, "", jsonIndent)

	logs.Log(r.Context(), fmt.Sprintf("useraction::%s", healthCheckHandler))
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(t))
}
