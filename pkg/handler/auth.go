package handler

import (
	"core"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

func (h *Handler) signUp(c *gin.Context) {
	var input core.User

	if err := c.BindJSON(&input); err != nil {
		h.services.Logging.WriteLog("invalid input body")
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	_, err := h.services.Registration.CreateUser(input)
	if err != nil {
		h.services.Logging.WriteLog(err.Error())
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	rand.Seed(time.Now().UTC().UnixNano())
	code := 1000 + rand.Intn(9999-1000)
	authUser := core.AuthUser{Phone: input.Phone, Code: code}
	_, err = h.services.Authorization.CreateAuthUser(authUser)
	if err != nil {
		h.services.Logging.WriteLog(err.Error())
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.Logging.WriteLog(fmt.Sprintf("User %s with phone number %d is sign-up", input.Name, input.Phone))

	SendSMS(input.Phone, code)

	c.JSON(http.StatusOK, map[string]interface{}{
		"code": code,
	})
}

type signInInput struct {
	Phone int64 `json:"phone" binding:"required"`
	Code  int   `json:"code" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		h.services.Logging.WriteLog("invalid input body")
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Phone, input.Code)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		h.services.Logging.WriteLog(err.Error())
		return
	}

	err = h.services.Logging.WriteLog(fmt.Sprintf("User with phone number %d is sign-in", input.Phone))

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
