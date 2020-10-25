package api

import "github.com/gin-gonic/gin"

func SetupAPI() *gin.Engine {

	engine := gin.Default()
	initAPI(engine)
	return engine

}
func initAPI(engine *gin.Engine) {
	demoAPIs(engine)

}
