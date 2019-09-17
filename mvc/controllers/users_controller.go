package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Washington-Ksea/golang-microservices/mvc/utils"

	"github.com/Washington-Ksea/golang-microservices/mvc/services"
)

//GetUser get user info from URL request
func GetUser(resp http.ResponseWriter, req *http.Request) {
	userID, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		//Just return the Bad request to the client
		return
	}

	user, apiErr := services.GetUser(userID)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		//Handle the err and return to the client
		return
	}

	//return user to client
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)

}
