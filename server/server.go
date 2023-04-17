package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
)
func StartServer(port string)  {
	r:=SetupRouter()
	r.Run(fmt.Sprintf("%s:%s", "127.0.0.1", port))
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/home", Home)
	r.GET("/students", FindAll)
	r.GET("/students/:id", FindOneById)
	r.POST("/students", Create)
	return r
}