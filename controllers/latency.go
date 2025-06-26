package controllers

import (
	"net/http"
	"time"

	"backend-server/utils"

	"github.com/gin-gonic/gin"
)

func LatencyHandler(c *gin.Context) {
	start := time.Now()
	time.Sleep(10 * time.Millisecond)
	compute := time.Since(start)

	utils.JSONWithOptionalDebug(c, http.StatusOK, gin.H{
		"compute_time_ms":  compute.Milliseconds(),
		"response_time_ms": compute.Milliseconds(),
		"expected_latency": "network latency + server compute time",
	})
}
