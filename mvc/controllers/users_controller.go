package controllers

import (
	"net/http"
	"strconv"

	"github.com/Washington-Ksea/golang-microservices/mvc/utils"
	"github.com/gin-gonic/gin"

	"github.com/Washington-Ksea/golang-microservices/mvc/services"
)

//GetUser get user info from URL request
func GetUser(c *gin.Context) {
	//userID, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	//c.Query("caller") // query caller
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.RespondErr(c, apiErr)
		//c.JSON(apiErr.StatusCode, apiErr)
		//jsonValue, _ := json.Marshal(apiErr)
		//resp.WriteHeader(apiErr.StatusCode)
		//resp.Write(jsonValue)
		//Just return the Bad request to the client
		return
	}

	user, apiErr := services.UsersService.GetUser(userID)
	if apiErr != nil {
		utils.RespondErr(c, apiErr)
		//c.JSON(apiErr.StatusCode, apiErr)
		//jsonValue, _ := json.Marshal(apiErr)
		//resp.WriteHeader(apiErr.StatusCode)
		//resp.Write(jsonValue)
		//Handle the err and return to the client
		return
	}
	utils.Respond(c, http.StatusOK, user)
	//c.JSON(http.StatusOK, user)
	//return user to client
	//jsonValue, _ := json.Marshal(user)
	//resp.Write(jsonValue)

}
