package v1

import (
	"gin-learn/model/response"

	"github.com/gin-gonic/gin"
)

func JsonInBlacklist(c *gin.Context) {
	response.FailWithMessage("token黑名单", c)
}
