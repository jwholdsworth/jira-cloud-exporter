# Example Metrics

```
# HELP jira_cloud_issue Shows the number of issues matching the JQL
# TYPE jira_cloud_issue gauge
jira_cloud_issue{assignee="",key="SE-52",project="Snakeoil Enterprises",status="To Do"} 1
jira_cloud_issue{assignee="",key="SE-53",project="Snakeoil Enterprises",status="To Do"} 1
jira_cloud_issue{assignee="john.smith",key="SE-24",project="Snakeoil Enterprises",status="In Progress"} 1
jira_cloud_issue{assignee="foo.bar",key="SE-51",project="Snakeoil Enterprises",status="In Progress"} 1
```
