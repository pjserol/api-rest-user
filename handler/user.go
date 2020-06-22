package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/pjserol/api-rest-user/common/logs"
	"github.com/pjserol/api-rest-user/common/utils"
	"github.com/pjserol/api-rest-user/model"
	"github.com/pjserol/api-rest-user/serviceapi"
)

const (
	getUserHandler  = "GetUserHandler"
	postUserHandler = "PostUserHandler"
	putUserHandler  = "PutUserHandler"
)

// GetUserHandler return a user from the id
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := logs.AssignRequestID(r.Context())
	start := utils.MakeTimestampMilli()
	w.Header().Set("Content-Type", jsonContentType)

	vars := mux.Vars(r)
	userID := vars["userId"]

	s, err := serviceapi.InitService(ctx)
	if err != nil {
		returnError(w, r, getUserHandler, []string{err.Error()}, start, http.StatusInternalServerError)
		return
	}

	user, err := s.GetUser(ctx, userID)
	if err != nil {
		returnError(w, r, getUserHandler, []string{err.Error()}, start, http.StatusInternalServerError)
		return
	}

	response, _ := json.MarshalIndent(userResponse{
		Success:  true,
		Messages: []string{},
		Time:     time.Now().UTC(),
		Timing: []timing{
			{
				Source:     getUserHandler,
				TimeMillis: utils.MakeTimestampMilli() - start,
			},
		},
		Response: user,
	}, "", jsonIndent)

	logs.Log(ctx, fmt.Sprintf("useraction::%s", getUserHandler))
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(response))
}

// PostUserHandler create a user
func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := logs.AssignRequestID(r.Context())
	start := utils.MakeTimestampMilli()
	w.Header().Set("Content-Type", jsonContentType)

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		returnError(w, r, postUserHandler, []string{err.Error()}, start, http.StatusInternalServerError)
		return
	}

	var request userRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		returnError(w, r, postUserHandler, []string{err.Error()}, start, http.StatusBadRequest)
		return
	}

	if errDetails, err := request.validate(); err != nil {
		returnError(w, r, postUserHandler, errDetails, start, http.StatusBadRequest)
		return
	}

	s, err := serviceapi.InitService(ctx)
	if err != nil {
		returnError(w, r, postUserHandler, []string{err.Error()}, start, http.StatusInternalServerError)
		return
	}

	input := model.CreateUserInput{
		Email:     request.Email,
		FirstName: request.FirstName,
		LastName:  request.LastName,
	}

	user, err := s.CreateUser(ctx, input)
	if err != nil {
		returnError(w, r, postUserHandler, []string{err.Error()}, start, http.StatusInternalServerError)
		return
	}

	response, _ := json.MarshalIndent(userResponse{
		Success:  true,
		Messages: []string{},
		Time:     time.Now().UTC(),
		Timing: []timing{
			{
				Source:     postUserHandler,
				TimeMillis: utils.MakeTimestampMilli() - start,
			},
		},
		Response: user,
	}, "", jsonIndent)

	logs.Log(ctx, fmt.Sprintf("useraction::%s", postUserHandler))
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(response))
}

// PutUserHandler update a user from the id
func PutUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := logs.AssignRequestID(r.Context())
	start := utils.MakeTimestampMilli()
	w.Header().Set("Content-Type", jsonContentType)

	vars := mux.Vars(r)
	userID := vars["userId"]

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		returnError(w, r, putUserHandler, []string{err.Error()}, start, http.StatusInternalServerError)
		return
	}

	var request userRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		returnError(w, r, putUserHandler, []string{err.Error()}, start, http.StatusBadRequest)
		return
	}

	if errDetails, err := request.validate(); err != nil {
		returnError(w, r, putUserHandler, errDetails, start, http.StatusBadRequest)
		return
	}

	s, err := serviceapi.InitService(ctx)
	if err != nil {
		returnError(w, r, putUserHandler, []string{err.Error()}, start, http.StatusInternalServerError)
		return
	}

	input := model.UpdateUserInput{
		ID:        userID,
		Email:     request.Email,
		FirstName: request.FirstName,
		LastName:  request.LastName,
	}

	user, err := s.UpdateUser(ctx, input)
	if err != nil {
		returnError(w, r, putUserHandler, []string{err.Error()}, start, http.StatusInternalServerError)
		return
	}

	response, _ := json.MarshalIndent(userResponse{
		Success:  true,
		Messages: []string{},
		Time:     time.Now().UTC(),
		Timing: []timing{
			{
				Source:     putUserHandler,
				TimeMillis: utils.MakeTimestampMilli() - start,
			},
		},
		Response: user,
	}, "", jsonIndent)

	logs.Log(ctx, fmt.Sprintf("useraction::%s", putUserHandler))
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(response))
}
