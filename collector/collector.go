package collector

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"app/config"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

// JiraCollector initiates the collection of Metrics from a JIRA instance
func JiraCollector() *Metrics {
	return &Metrics{
		issue: prometheus.NewDesc(prometheus.BuildFQName("jira", "cloud", "issue"),
			"Shows the number of issues matching the JQL",
			[]string{"status", "project", "key", "assignee", "location", "priority"}, nil,
		),
	}
}

// Describe writes all descriptors to the prometheus desc channel.
func (collector *Metrics) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.issue
}

//Collect implements required collect function for all prometheus collectors
func (collector *Metrics) Collect(ch chan<- prometheus.Metric) {

	collectedIssues, err := fetchJiraIssues()
	if err != nil {
		log.Error(err)
		return
	}

	for _, issue := range collectedIssues.Issues {
		createdTimestamp := convertToUnixTime(issue.Fields.Created)
		ch <- prometheus.MustNewConstMetric(collector.issue, prometheus.CounterValue, createdTimestamp, issue.Fields.Status.Name, issue.Fields.Project.Name, issue.Key, issue.Fields.Assignee.Name, issue.Fields.Location.Name, issue.Fields.Priority.Name)
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

func fetchJiraIssues() (jiraIssue, error) {

	cfgs, err := config.Init()
	if err != nil {
		log.Error(err)
	}
	var AllIssues jiraIssue

	for _, cfg := range cfgs {
		var ji jiraIssue

		err = validateJiraCfg(cfg)
		if err != nil {
			return ji, err
		}

		url := fmt.Sprintf("%s/rest/api/2/search?jql=%s", cfg.JiraURL, cfg.JiraJql)
		resp, err := fetchAPIResults(url, cfg.JiraUsername, cfg.JiraToken)

		err = json.Unmarshal(resp, &ji)
		if err != nil {
			return ji, err
		}

		AllIssues.Issues = append(AllIssues.Issues, ji.Issues...)

		// Pagination support
		if ji.Total > len(AllIssues.Issues) {
			var startsAt int

			for {

				// we use startsAt to track our process through the pagination
				// here we set it to the lenghth of the intial capture, + 1
				startsAt = len(AllIssues.Issues) + 1

				url := fmt.Sprintf("%s/rest/api/2/search?jql=%s&startAt=%d", cfg.JiraURL, cfg.JiraJql, startsAt)
				resp, err := fetchAPIResults(url, cfg.JiraUsername, cfg.JiraToken)

				err = json.Unmarshal(resp, &ji)
				if err != nil {
					return ji, err
				}

				AllIssues.Issues = append(AllIssues.Issues, ji.Issues...)

				// The API has a funny way of counting, to ensure we get all issues
				// We break when no more issues are returned
				if len(ji.Issues) == 0 {
					log.Debug("No futher issues returned from API")
					break
				}

			}
		}
	}

	return AllIssues, nil
}

func fetchAPIResults(url, user, token string) ([]byte, error) {

	client := http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "jira-cloud-exporter")
	req.SetBasicAuth(user, token)
	log.Infof(fmt.Sprintf("Sending request to %s", url))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

func validateJiraCfg(cfg config.Config) error {

	_, err := url.ParseRequestURI(cfg.JiraURL)
	if err != nil {
		return fmt.Errorf("Error validating URL, please ensure the URL is valid: %v", err)
	}

	if !strings.HasPrefix(cfg.JiraURL, "https://") {
		return fmt.Errorf("The Jira URL: %s is insecure, your API token is being sent in clear text", cfg.JiraURL)
	}

	if cfg.JiraUsername == "" || cfg.JiraToken == "" {
		return fmt.Errorf("Check credentials supplied are set and valid")
	}

	return nil
}
