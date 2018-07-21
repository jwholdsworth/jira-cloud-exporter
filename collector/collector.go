package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

//Define a struct for you collector that contains pointers
//to prometheus descriptors for each metric you wish to expose.
//Note you can also include fields of other types if they provide utility
//but we just won't be exposing them as metrics.
type JiraMetrics struct {
	jiraNumberOfIssues *prometheus.Desc
}

//You must create a constructor for you collector that
//initializes every descriptor and returns a pointer to the collector
func JiraCollector() *JiraMetrics {
	return &JiraMetrics{
		jiraNumberOfIssues: prometheus.NewDesc(prometheus.BuildFQName("jira", "issues", "status"),
			"Shows the number of issues in the workspace",
			[]string{"status", "workspace"}, nil,
		),
	}
}

//Each and every collector must implement the Describe function.
//It essentially writes all descriptors to the prometheus desc channel.
func (collector *JiraMetrics) Describe(ch chan<- *prometheus.Desc) {

	//Update this section with the each metric you create for a given collector
	ch <- collector.jiraNumberOfIssues
}

//Collect implements required collect function for all promehteus collectors
func (collector *JiraMetrics) Collect(ch chan<- prometheus.Metric) {

	var issues float64

	//Write latest value for each metric in the prometheus metric channel.
	//Note that you can pass CounterValue, GaugeValue, or UntypedValue types here.
	ch <- prometheus.MustNewConstMetric(collector.jiraNumberOfIssues, prometheus.GaugeValue, issues, "unassigned", "SB")
}
