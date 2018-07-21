package config

import "os"

type Config struct {
	JiraToken     string
	JiraUsername  string
	JiraJql       string
	JiraURL       string
	MetricsPath   string
	ListenAddress string
}

// Init populates the Config struct based on environmental runtime configuration
func Init() Config {

	jiraToken := getEnv("JIRA_TOKEN", "")
	jiraUsername := getEnv("JIRA_USERNAME", "")
	jiraJql := getEnv("JIRA_JQL", "")
	jiraURL := getEnv("JIRA_URL", "https://infinityworks.atlassian.net")
	metricsPath := getEnv("METRICS_PATH", "/metrics")
	listenAddress := getEnv("LISTEN_ADDRESS", ":9800")

	appConfig := Config{
		jiraToken,
		jiraUsername,
		jiraJql,
		jiraURL,
		metricsPath,
		listenAddress,
	}

	return appConfig
}

func getEnv(environmentVariable string, defaultValue string) string {
	envVar := os.Getenv(environmentVariable)
	if len(envVar) == 0 {
		return defaultValue
	}

	return envVar
}
