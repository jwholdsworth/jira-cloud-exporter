package collector

import "github.com/prometheus/client_golang/prometheus"

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
