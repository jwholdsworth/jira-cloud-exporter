package collector

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jwholdsworth/jira-cloud-exporter/config"
	"github.com/prometheus/client_golang/prometheus"
)

// JiraCollector initiates the collection of metrics from the JIRA instance
func JiraCollector() *JiraMetrics {
	return &JiraMetrics{
		jiraIssues: prometheus.NewDesc(prometheus.BuildFQName("jira", "issue", "status"),
			"Shows the number of issues in the workspace",
			[]string{"status", "project", "id", "assignee"}, nil,
		),
	}
}

// Describe writes all descriptors to the prometheus desc channel.
func (collector *JiraMetrics) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.jiraIssues
}

//Collect implements required collect function for all promehteus collectors
func (collector *JiraMetrics) Collect(ch chan<- prometheus.Metric) {

	collectedIssues := fetchJiraIssues()

	for _, issue := range collectedIssues.Issues {
		ch <- prometheus.MustNewConstMetric(collector.jiraIssues, prometheus.GaugeValue, 1, issue.Fields.Status.Name, issue.Fields.Project.Name, issue.ID, issue.Fields.Assignee.Name)
	}
}

func fetchJiraIssues() JiraIssues {
	// DI this
	cfg := config.Init()
	var jiraIssues JiraIssues

	client := http.Client{}
	url := fmt.Sprintf("%s/rest/api/2/search?jql=%s", cfg.JiraURL, cfg.JiraJql)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "jira-cloud-exporter")
	req.SetBasicAuth(cfg.JiraUsername, cfg.JiraToken)
	res, err := client.Do(req)

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	jsonError := json.Unmarshal(body, &jiraIssues)
	if jsonError != nil {
		log.Fatal(jsonError)
	}

	return jiraIssues
}
