package main

import (
	"fmt"
	"net/http"
	"os"

	"app/collector"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

func main() {
	metricsPath := getEnv("METRICS_PATH", "/metrics")
	listenAddress := getEnv("LISTEN_ADDRESS", ":9800")
	jiraCollector := collector.JiraCollector()
	prometheus.MustRegister(jiraCollector)

	http.Handle(metricsPath, promhttp.Handler())
	if metricsPath != "/" {
		http.Handle("/", http.RedirectHandler(metricsPath, http.StatusMovedPermanently))
	}
	log.Info(fmt.Sprintf("Listening on %s", listenAddress))
	log.Fatal(http.ListenAndServe(listenAddress, nil))
}

func getEnv(environmentVariable, defaultValue string) string {
	envVar := os.Getenv(environmentVariable)
	if len(envVar) == 0 {
		return defaultValue
	}

	return envVar
}
