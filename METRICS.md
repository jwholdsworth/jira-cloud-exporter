# Example Metrics

```
# HELP jira_cloud_exporter Shows the number of issues matching the JQL
# TYPE jira_cloud_exporter gauge
jira_cloud_exporter{assignee="",key="SE-52",project="Snakeoil Enterprises",status="To Do"} 1
jira_cloud_exporter{assignee="",key="SE-53",project="Snakeoil Enterprises",status="To Do"} 1
jira_cloud_exporter{assignee="john.smith",key="SE-24",project="Snakeoil Enterprises",status="In Progress"} 1
jira_cloud_exporter{assignee="foo.bar",key="SE-51",project="Snakeoil Enterprises",status="In Progress"} 1
```
