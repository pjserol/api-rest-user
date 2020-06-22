package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/pjserol/api-rest-user/common/logs"
	"github.com/pjserol/api-rest-user/config"
	"github.com/pjserol/api-rest-user/db"
	"github.com/pjserol/api-rest-user/handler"
)

func main() {
	env := config.InitEnvironment()
	ctx := context.Background()

	if env.IsTestMode {
		// display the line and file in the logs
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	}

	// run in a goroutine to don't be a blocking call
	go db.Initialise(ctx)

	// See http://www.gorillatoolkit.org/pkg/mux
	r := mux.NewRouter()

	r.HandleFunc("/health-check/", handler.HealthCheckHandler).Methods("GET")

	r.HandleFunc("/v1/user/{userId:.*?}/", handler.GetUserHandler).Methods("GET")

	r.HandleFunc("/v1/user/", handler.PostUserHandler).Methods("POST")

	r.HandleFunc("/v1/user/{userId:.*?}/", handler.PutUserHandler).Methods("PUT")

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("0.0.0.0:%d", env.PortNumber),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logs.Log(ctx, fmt.Sprintf("environment:%s::Ready to serve requests on 0.0.0.0:%d", env.AppEnvironment, env.PortNumber))
	logs.Log(ctx, "API started!")

	log.Fatal(srv.ListenAndServe())
}
