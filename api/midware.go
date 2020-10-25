package api

import (
	"github.com/1101137105/dcard/glob"
	"github.com/1101137105/dcard/service"
	"github.com/gin-gonic/gin"
)

var (
	IpMap          = make(map[string]int)
	IpTimeMap      = make(map[string]int64)
	IpRedisTimeMap = make(map[string]int64)
)

func CheckRequestLimit(limit int) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		var err error
		var val int
		if val, err = service.DemoService.CheckRequestLimit(ip, limit); err != nil {
			glob.SetCommonResponse(c, nil, err)

			c.Abort()
		}
		c.Set(ip, val)
		c.Next()
	}
}
func CheckRequestRedisLimit(limit int) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		var err error
		var val int
		if val, err = service.DemoRedisService.CheckRequestLimit(ip, limit); err != nil {
			glob.SetCommonResponse(c, nil, err)

			c.Abort()
		}
		c.Set(ip, val)
		c.Next()
	}
}
