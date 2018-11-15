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

	jiraToken := getEnv("JIRA_TOKEN", "")
	jiraUsername := getEnv("JIRA_USERNAME", "")
	jiraJql := getEnv("JIRA_JQL", "")
	jiraURL := getEnv("JIRA_URL", "")

	tokens := strings.Split(jiraToken, ",")
	usernames := strings.Split(jiraUsername, ",")
	jqls := strings.Split(jiraJql, ",")
	urls := strings.Split(jiraURL, ",")
	// fmt.Println("URLs:", urls)
	// fmt.Println("Tokens:", tokens)
	// fmt.Println("Usernames:", usernames)
	// fmt.Println("JQLs:", jqls)

	if len(urls) != len(usernames) {
		log.Error("The number of Jira URLs doesn't match the number of Usernames")
		os.Exit(1)
	} else if len(usernames) != len(tokens) {
		log.Error("The number of Jira Usernames doesn't match the number of Jira Tokens")
		os.Exit(1)
	}

	var appConfig []Config
	appConfig = make([]Config, 0)

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
