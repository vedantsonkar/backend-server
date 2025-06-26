package utils

import (
	"time"

	"github.com/gin-gonic/gin"
)

func JSONWithOptionalDebug(c *gin.Context, status int, payload gin.H) {
	if debug, _ := c.Get("debugEnabled"); debug == true {
		debugData := gin.H{}

		if rt, ok := c.Get("requestTime"); ok {
			debugData["requestTime"] = rt.(time.Duration).String()
		}
		if dt, ok := c.Get("dbTime"); ok {
			debugData["dbTime"] = dt.(time.Duration).String()
		}

		payload["debug"] = debugData
	}
	c.JSON(status, payload)
}
