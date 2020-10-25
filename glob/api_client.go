package glob

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type CommonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SetCommonResponse(c *gin.Context, item interface{}, err error) {

	var statusCode int
	responseData := CommonResponse{}

	if err != nil {

		responseData.Message = "操作失敗"
		responseData.Data = err.Error()
		statusCode = http.StatusBadRequest

		log.Println("apis %s", err.Error())
	} else {
		responseData.Message = "操作成功"
		responseData.Data = item
		statusCode = http.StatusOK

	}
	responseData.Code = statusCode

	c.JSON(statusCode, responseData)

}
