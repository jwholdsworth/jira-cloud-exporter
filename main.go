package main

import (
	"fmt"
	"net/http"

	"github.com/infinityworks/jira-cloud-exporter/collector"
	"github.com/infinityworks/jira-cloud-exporter/config"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

func main() {
	cfg := config.Init()
	jiraCollector := collector.JiraCollector()
	prometheus.MustRegister(jiraCollector)

	http.Handle(cfg.MetricsPath, promhttp.Handler())
	http.Handle("/", http.RedirectHandler(cfg.MetricsPath, http.StatusMovedPermanently))
	log.Info(fmt.Sprintf("Listening on %s", cfg.ListenAddress))
	log.Fatal(http.ListenAndServe(cfg.ListenAddress, nil))
}
