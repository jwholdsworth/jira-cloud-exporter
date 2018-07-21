# jira-cloud-exporter

Exposes basic JIRA metrics to a Prometheus compatible endpoint.

## Configuration

Configuration is provided in the form of environment variables.

### Required

* `JIRA_TOKEN` should be self-explanatory. Create one under your user account settings in JIRA Cloud.
* `JIRA_USERNAME` is the username associated with the token.
* `JIRA_URL` is the URL to your organisation's JIRA application.

### Optional
* `JIRA_JQL` is the JIRA query language search filter (defaults to empty, so you'll get everything)
* `METRICS_PATH` is the endpoint Prometheus should scrape on this exporter. Defaults to `/metrics`
* `LISTEN_ADDRESS` is the IP and port to bind this exporter to. Defaults to `:9800`.
