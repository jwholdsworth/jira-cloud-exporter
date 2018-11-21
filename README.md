# jira-cloud-exporter

Exposes basic [JIRA](https://www.atlassian.com/software/jira) metrics (from one or more servers) to a [Prometheus](https://prometheus.io) compatible endpoint using the metric name: *jira_cloud_exporter*.

## Building

`make build`

## Configuration

Configuration is provided in the form of environment variables. If multiple Jira servers are being queried, each variable value should be a comma separated list.

### Required

* `JIRA_TOKEN` should be self-explanatory. Create one under your user account settings in JIRA Cloud.
* `JIRA_USERNAME` is the username associated with the token.
* `JIRA_URL` is the URL to your organisation's JIRA application.

### Optional

* `JIRA_JQL` is the JIRA query language search filter (defaults to empty, so you'll get everything). If querying multiple Jira servers this variable is *not* optional. You must delimit null values if not specifying a filter for a server. For instance, if querying two servers:
  * Filter for first, no filter for second: `JIRA_JQL="project=test,"`
  * No filter for first, filter for second: `JIRA_JQL=",project=test"`
  * No filter for either: `JIRA_JQL=","`
* `METRICS_PATH` is the endpoint Prometheus should scrape on this exporter. Defaults to `/metrics`
* `LISTEN_ADDRESS` is the IP and port to bind this exporter to. Defaults to `:9800`.

## Metrics

See the example metrics in `METRICS.md`.
