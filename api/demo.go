package api

import (
	"github.com/1101137105/dcard/glob"
	"github.com/gin-gonic/gin"
)

type demoAPI struct{}

var _demoAPI = demoAPI{}

func demoAPIs(engine *gin.Engine) {
	g := engine.Group("/demo/")
	{
		_demoAPI.demo(g)
		_demoAPI.demoRedis(g)

	}
}

func (a demoAPI) demo(g *gin.RouterGroup) {
	limit := glob.Config.Www.Limit

	g.Use(CheckRequestLimit(limit))
	g.GET("/", a.CheckRequestLimit)

}
func (a demoAPI) demoRedis(g *gin.RouterGroup) {
	limit := glob.Config.Www.Limit

	g.Use(CheckRequestRedisLimit(limit))
	g.GET("/redis", a.CheckRequestLimitToRedis)

}

func (a demoAPI) CheckRequestLimit(c *gin.Context) {
	ip := c.ClientIP()
	val, _ := c.Get(ip)
	glob.SetCommonResponse(c, val, nil)

}
func (a demoAPI) CheckRequestLimitToRedis(c *gin.Context) {
	ip := c.ClientIP()
	val, _ := c.Get(ip)
	glob.SetCommonResponse(c, val, nil)

}
