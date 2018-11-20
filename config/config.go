package config

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	JiraToken    string
	JiraUsername string
	JiraJql      string
	JiraURL      string
}

// Init populates the Config struct based on environmental runtime configuration
func Init() ([]Config, error) {

	jiraToken := getEnv("JIRA_TOKEN", "")
	jiraUsername := getEnv("JIRA_USERNAME", "")
	jiraJql := getEnv("JIRA_JQL", "")
	jiraURL := getEnv("JIRA_URL", "")

	tokens := strings.Split(jiraToken, ",")
	usernames := strings.Split(jiraUsername, ",")
	jqls := strings.Split(jiraJql, ",")
	urls := strings.Split(jiraURL, ",")

	appConfig := make([]Config, 0)

	if len(urls) != len(usernames) {
		err := fmt.Errorf("The number of Jira URLs doesn't match the number of Jira Usernames")
		return appConfig, err
	} else if len(usernames) != len(tokens) {
		err := fmt.Errorf("The number of Jira Usernames doesn't match the number of Jira Tokens")
		return appConfig, err
	}

	for i, items := range urls {
		temp := Config{tokens[i], usernames[i], jqls[i], items}
		appConfig = append(appConfig, temp)
	}

	return appConfig, nil
}

func getEnv(environmentVariable, defaultValue string) string {
	envVar := os.Getenv(environmentVariable)
	if len(envVar) == 0 {
		return defaultValue
	}

	return envVar
}
