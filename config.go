package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	JiraURL        string `json:"jira_url"`
	JiraAPIKey     string `json:"jira_apikey"`
	JiraUsername   string `json:"jira_username"`
	JiraProjectKey string `json:"jira_project_key"`
}

func loadConfig() (*Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get user home directory: %w", err)
	}

	configDir := filepath.Join(homeDir, ".config", "squirrelcli")
	configPath := filepath.Join(configDir, "config.json")

	// Check if config file exists, if not initilinitialise it.
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return initConfig(configPath)
	}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file %s: %w", configPath, err)
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("failed to decode config file: %w", err)
	}

	return &config, nil
}

func initConfig(configPath string) (*Config, error) {
	configDir := filepath.Dir(configPath)
	// Create config directory if it doesn't exist
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create config directory %s: %w", configDir, err)
	}

	// Prompt user for configuration values
	fmt.Println("🔧 First time setup: Please provide your JIRA configuration")

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("JIRA URL (e.g., https://your-domain.atlassian.net): ")
	scanner.Scan()
	jiraURL := strings.TrimSpace(scanner.Text())

	fmt.Print("JIRA API Key (generate a new one here: https://id.atlassian.com/manage-profile/security/api-tokens): ")
	scanner.Scan()
	jiraAPIKey := strings.TrimSpace(scanner.Text())

	fmt.Print("JIRA Username: ")
	scanner.Scan()
	jiraUsername := strings.TrimSuffix(strings.TrimSpace(scanner.Text()), "/")

	fmt.Print("Project Key (e.g., BOSS): ")
	scanner.Scan()
	projectKey := strings.TrimSpace(scanner.Text())

	// Validate input
	if jiraURL == "" || jiraAPIKey == "" || jiraUsername == "" || projectKey == "" {
		return nil, fmt.Errorf("all configuration values are required")
	}

	config := Config{
		JiraURL:        jiraURL,
		JiraAPIKey:     jiraAPIKey,
		JiraUsername:   jiraUsername,
		JiraProjectKey: projectKey,
	}

	file, err := os.Create(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create config file %s: %w", configPath, err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(config); err != nil {
		return nil, fmt.Errorf("failed to write config: %w", err)
	}

	fmt.Printf("✅ Configuration saved to %s\n\n", configPath)

	return &config, nil
}
