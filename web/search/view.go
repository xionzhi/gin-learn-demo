package search

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	KeyWord string `form:"keyword" binding:"required,min=2,max=30"`
}

type WebResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Search(c *gin.Context) {
	req := LoginRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		resp := WebResponse{Code: 0, Message: err.Error(), Data: ""}
		c.JSON(http.StatusOK, resp)
	} else {
		// query user info
		if req.KeyWord == "hello" {
			resp := WebResponse{Code: 1, Message: "ok", Data: "hello world!"}
			c.JSON(http.StatusOK, resp)
		} else {
			resp := WebResponse{Code: 0, Message: "failed", Data: ""}
			c.JSON(http.StatusOK, resp)
		}
	}
}
