package main

import (
	"github.com/1101137105/dcard/api"
	"github.com/1101137105/dcard/glob"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	err    error
	engine *gin.Engine
)

func main() {
	glob.InitConfig()
	glob.SetupRedis()
	engine = api.SetupAPI()

	err = engine.Run(glob.Config.Www.Addr)
	if err != nil {
		log.Println("Dcard fatal error; server stopped; msg: %v;", err)
		panic(err)
	}
}
