package collector

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"jira-cloud-exporter/config"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

// JiraCollector initiates the collection of metrics from the JIRA instance
func JiraCollector() *JiraMetrics {
	return &JiraMetrics{
		jiraIssues: prometheus.NewDesc(prometheus.BuildFQName("jira", "cloud", "exporter"),
			"Shows the number of issues matching the JQL",
			[]string{"status", "project", "key", "assignee"}, nil,
		),
	}
}

// Describe writes all descriptors to the prometheus desc channel.
func (collector *JiraMetrics) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.jiraIssues
}

//Collect implements required collect function for all prometheus collectors
func (collector *JiraMetrics) Collect(ch chan<- prometheus.Metric) {

	collectedIssues := fetchJiraIssues()

	for _, issue := range collectedIssues.Issues {
		createdTimestamp := convertToUnixTime(issue.Fields.Created)
		ch <- prometheus.MustNewConstMetric(collector.jiraIssues, prometheus.CounterValue, createdTimestamp, issue.Fields.Status.Name, issue.Fields.Project.Name, issue.Key, issue.Fields.Assignee.Name)
	}
}

func convertToUnixTime(timestamp string) float64 {
	layout := "2006-01-02T15:04:05.000-0700"
	dateTime, err := time.Parse(layout, timestamp)
	if err != nil {
		log.Error(err)
		return 0
	}

	return float64(dateTime.Unix())
}

func fetchJiraIssues() JiraIssues {
	// DI this

	cfgs := config.Init()
	// fmt.Println("Array:", cfgs)

	var AllIssues JiraIssues

	for _, cfg := range cfgs {
		var jiraIssues JiraIssues

		// fmt.Println("Config:", cfg)

		// Confirm Jira URL begins with the http:// or https:// scheme specification
		// Also emit a warning if HTTPS isn't being used
		if !strings.HasPrefix(cfg.JiraURL, "http") {
			log.Error("The Jira URL: ", cfg.JiraURL, " does not begin with 'http'")
			// Return an error to the calling function instead:
			return jiraIssues
		} else if !strings.HasPrefix(cfg.JiraURL, "https://") {
			log.Warn("The Jira URL: ", cfg.JiraURL, " is insecure, your API token is being sent in clear text")
		}

		if len(cfg.JiraUsername) < 6 {
			log.Warn("The Jira username has fewer than 6 characters, are you sure it is valid?")
		}
		if len(cfg.JiraToken) < 10 {
			log.Warn("The Jira token has fewer than 10 characters, are you sure it is valid?")
		}

		client := http.Client{}
		url := fmt.Sprintf("%s/rest/api/2/search?jql=%s", cfg.JiraURL, cfg.JiraJql)
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Error(err)
			return jiraIssues
		}
		req.Header.Set("User-Agent", "jira-cloud-exporter")
		req.SetBasicAuth(cfg.JiraUsername, cfg.JiraToken)
		log.Info(fmt.Sprintf("Sending request to %s", url))
		res, err := client.Do(req)

		if err != nil {
			log.Error(err)
			return jiraIssues
		}

		body, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			log.Error(readErr)
			return jiraIssues
		}

		jsonError := json.Unmarshal(body, &jiraIssues)
		if jsonError != nil {
			log.Error(jsonError)
		}

		AllIssues.Issues = append(AllIssues.Issues, jiraIssues.Issues...)

	}

	return AllIssues
}

type error interface {
	Error() string
}
