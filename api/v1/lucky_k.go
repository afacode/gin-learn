package v1

import (
	"gin-learn/model/response"
	"gin-learn/service"

	"github.com/gin-gonic/gin"
)

func GetK(c *gin.Context) {
	err, list := service.GetAllK()
	if err != nil {
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithData(list, c)
	}
}
