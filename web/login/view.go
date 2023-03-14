package login

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Account  string `json:"account" binding:"required,min=2,max=100"`
	Password string `json:"password" binding:"required,min=2,max=100"`
}

type WebResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Login(c *gin.Context) {
	req := LoginRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		resp := WebResponse{Code: 0, Message: err.Error(), Data: nil}
		c.JSON(http.StatusOK, resp)
	} else {
		// query user info
		userQuery, err := QueryOneUsersByAccount(req.Account)
		if err != nil {
			panic(err)
		}
		if userQuery.Password == req.Password {
			resp := WebResponse{Code: 1, Message: "ok", Data: userQuery.ID}
			c.JSON(http.StatusOK, resp)
		} else {
			resp := WebResponse{Code: 0, Message: "failed", Data: nil}
			c.JSON(http.StatusOK, resp)
		}
	}
}
