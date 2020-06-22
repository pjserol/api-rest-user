package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/pjserol/api-rest-user/common/utils"
)

const (
	jsonContentType = "application/json; charset=utf-8"
	jsonIndent      = "    "
)

func returnError(w http.ResponseWriter, r *http.Request, source string, messages []string, start int64, status int) {
	t, _ := json.MarshalIndent(errorResponse{
		Success:  false,
		Messages: messages,
		Time:     time.Now().UTC(),
		Timing: []timing{
			{
				Source:     source,
				TimeMillis: utils.MakeTimestampMilli() - start,
			},
		},
	}, "", jsonIndent)

	w.WriteHeader(status)
	fmt.Fprint(w, string(t))
}
