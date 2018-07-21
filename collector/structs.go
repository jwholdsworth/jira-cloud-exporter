package collector

import "github.com/prometheus/client_golang/prometheus"

type JiraMetrics struct {
	jiraIssues *prometheus.Desc
}

type JiraIssue struct {
	Fields Fields `json:"fields"`
	Key    string `json:"key"`
}

type Fields struct {
	Assignee Assignee `json:"assignee"`
	Project  Project  `json:"project"`
	Status   Status   `json:"status"`
}

type Assignee struct {
	Name string `json:"name"`
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
