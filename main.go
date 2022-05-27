package main

import (
	_ "github.com/diegoluisi/template-app-golang/config"
	"github.com/diegoluisi/template-app-golang/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.New()
	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  true,
			"message": "Say hello for your new Golang+gin microservise (By backstage template-app-golang)",
		})
	})
	server.GET("/health-check/liveness", controllers.HealthCheckLiveness)
	server.GET("/health-check/readiness", controllers.HealthCheckReadiness)
	server.Run()
}
