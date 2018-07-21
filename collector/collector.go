package collector

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/infinityworks/jira-cloud-exporter/config"
	"github.com/prometheus/client_golang/prometheus"
)

//Define a struct for you collector that contains pointers
//to prometheus descriptors for each metric you wish to expose.
//Note you can also include fields of other types if they provide utility
//but we just won't be exposing them as metrics.
type JiraMetrics struct {
	jiraIssues *prometheus.Desc
}

type JiraIssue struct {
	Fields Fields `json:"fields"`
	ID     string `json:"key"`
}

type Fields struct {
	Project Project `json:"project"`
	Status  Status  `json:"status"`
}

type Status struct {
	Name string `json:"name"`
}

type Project struct {
	Name string `json:"name"`
}

type JiraIssues struct {
	Issues []JiraIssue `json:"issues"`
}

//You must create a constructor for you collector that
//initializes every descriptor and returns a pointer to the collector
func JiraCollector() *JiraMetrics {
	return &JiraMetrics{
		jiraIssues: prometheus.NewDesc(prometheus.BuildFQName("jira", "issue", "status"),
			"Shows the number of issues in the workspace",
			[]string{"status", "project", "id"}, nil,
		),
	}
}

//Each and every collector must implement the Describe function.
//It essentially writes all descriptors to the prometheus desc channel.
func (collector *JiraMetrics) Describe(ch chan<- *prometheus.Desc) {

	//Update this section with the each metric you create for a given collector
	ch <- collector.jiraIssues
}

//Collect implements required collect function for all promehteus collectors
func (collector *JiraMetrics) Collect(ch chan<- prometheus.Metric) {

	collectedIssues := fetchJiraIssues()

	//Write latest value for each metric in the prometheus metric channel.
	//Note that you can pass CounterValue, GaugeValue, or UntypedValue types here.
	for _, issue := range collectedIssues.Issues {
		ch <- prometheus.MustNewConstMetric(collector.jiraIssues, prometheus.GaugeValue, 1, issue.Fields.Status.Name, issue.Fields.Project.Name, issue.ID)
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
