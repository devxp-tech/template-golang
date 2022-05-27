package controllers

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/gritzkoo/golang-health-checker/pkg/healthcheck"
	"github.com/sirupsen/logrus"
)

// HealthCheckLiveness show a simple check
func HealthCheckLiveness(c *gin.Context) {
	c.JSON(200, healthcheck.HealthCheckerSimple())
}

// HealthCheckReadiness return a detailed status of all integrations in the list
func HealthCheckReadiness(c *gin.Context) {
	version, _ := ioutil.ReadFile("rev.txt")
	result := healthcheck.HealthCheckerDetailed(healthcheck.ApplicationConfig{
		Name:    "MY-APPLICATION",
		Version: string(version),
		Integrations: []healthcheck.IntegrationConfig{
			{
				Type: healthcheck.Web,
				Name: "Just a simple test",
				Host: "https://github.com/status",
			},
		},
	})
	if !result.Status {
		// the line below will print the error on logs, so you need to
		// configure grafana or instana alarms with this context
		logrus.WithField("error", result).Warning("Health check fails")
	}
	// the pod will always return a status 200 to prevent shotdown
	c.JSON(200, result)
}
