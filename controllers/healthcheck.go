package controllers

import (
	"io/ioutil"
	"net/http"

	"github.com/diegoluisi/${{values.component_id}}/pkd/github"
	"github.com/gin-gonic/gin"
	"github.com/gritzkoo/golang-health-checker-lw/pkg/healthchecker"
	"github.com/sirupsen/logrus"
)

var version, _ = ioutil.ReadFile("rev.txt")
var checker = healthchecker.New(healthchecker.Config{
	Name:    "${{values.component_id}}",
	Version: string(version),
	Integrations: []healthchecker.Check{
		{
			Name:   "GitHub Api Integration",
			Handle: github.Status, // you should write your own tests to pass here!
		},
	},
})

// HealthCheckLiveness show a simple check
func HealthCheckLiveness(c *gin.Context) {
	c.JSON(http.StatusOK, checker.Liveness())
}

// HealthCheckReadiness return a detailed status of all integrations in the list
func HealthCheckReadiness(c *gin.Context) {
	result := checker.Readiness()
	if !result.Status {
		// the line below will print the error on logs, so you need to
		// configure grafana or instana alarms with this context
		logrus.WithField("error", result).Warning("Health check fails")
	}
	// the pod will always return a status 200 to prevent shotdown
	c.JSON(http.StatusOK, result)
}
