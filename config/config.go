package config

import (
	"os"
	"strings"

	"github.com/prometheus/common/log"
)

type Config struct {
	JiraToken    string
	JiraUsername string
	JiraJql      string
	JiraURL      string
}

// Init populates the Config struct based on environmental runtime configuration
func Init() []Config {

	// Should read these from a config file (viper module):
	jiraToken := getEnv("JIRA_TOKEN", "")
	jiraUsername := getEnv("JIRA_USERNAME", "")
	jiraJql := getEnv("JIRA_JQL", "")
	jiraURL := getEnv("JIRA_URL", "")

	tokens := strings.Split(jiraToken, ",")
	usernames := strings.Split(jiraUsername, ",")
	jqls := strings.Split(jiraJql, ",")
	urls := strings.Split(jiraURL, ",")

	if len(urls) != len(usernames) {
		log.Fatal("The number of Jira URLs doesn't match the number of Usernames")
		// Return an error to the calling function instead:

	} else if len(usernames) != len(tokens) {
		log.Fatal("The number of Jira Usernames doesn't match the number of Jira Tokens")
		// Return an error to the calling function instead:
	}

	// The line below shouldn't be needed
	//var appConfig []Config
	appConfig := make([]Config, 0)

	for i, items := range urls {
		temp := Config{tokens[i], usernames[i], jqls[i], items}
		appConfig = append(appConfig, temp)
	}

	return appConfig
}

func getEnv(environmentVariable, defaultValue string) string {
	envVar := os.Getenv(environmentVariable)
	if len(envVar) == 0 {
		return defaultValue
	}

	return envVar
}
