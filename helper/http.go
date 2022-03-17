package helper

import (
	"net/http"
	"strconv"

	"kumparan/model"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

var GoCache *cache.Cache

func CheckCache(c *gin.Context) {
	if c.Request.Method == "GET" {
		url := GetFullReqUrl(c)
		data, found := GoCache.Get(url)
		if found {
			HandleSuccess(c, data)
			c.Abort()
			return
		}
	}
	c.Next()
}

func HandleSuccess(c *gin.Context, data interface{}) {
	responData := model.Response{
		Status:  "200",
		Message: "Success",
		Data:    data,
	}
	c.JSON(http.StatusOK, responData)
}

func HandleError(c *gin.Context, status int, message string) {
	responData := model.Response{
		Status:  strconv.Itoa(status),
		Message: message,
	}
	c.JSON(status, responData)
}

func GetFullReqUrl(c *gin.Context) string {
	url := c.FullPath()
	rawQuery := c.Request.URL.RawQuery
	if rawQuery != "" {
		url = url + "?" + rawQuery
	}
	return url
}
