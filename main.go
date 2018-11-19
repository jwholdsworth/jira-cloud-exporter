package main

import (
	"fmt"
	"jira-cloud-exporter/collector"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Should read these from a config file (viper module):
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
