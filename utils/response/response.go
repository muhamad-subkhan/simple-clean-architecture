package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type Respon struct {
	Status  int      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}


func Succes(c *gin.Context, data interface{}) {
	response := Respon{
		Status: http.StatusOK,
		Message: "Success",
		Data: data,
	}
	c.JSON(http.StatusOK, response)
}

func Error(c *gin.Context, Status int, message string) {
	response := Respon{
		Status: Status,
		Message: message,
	}

	c.JSON(Status, response)
}