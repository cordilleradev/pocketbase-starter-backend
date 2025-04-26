package common

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	// Pocketbase
	AppName     string
	FrontendURL string

	// SMTP Settings
	SMTPEnabled    bool
	SMTPHost       string
	SMTPPort       int
	SMTPUsername   string
	SMTPPassword   string
	SendingName    string
	SendingAddress string
}

func NewConfig(devMode bool) (Config, error) {
	if devMode {
		// Load .env file in development mode
		err := godotenv.Load()
		if err != nil {
			log.Printf("Warning: Error loading .env file: %v", err)
		}
	}

	// Default SMTP port if not specified
	smtpPort := 587
	if portStr := os.Getenv("SMTP_PORT"); portStr != "" {
		port, err := strconv.Atoi(portStr)
		if err != nil {
			return Config{}, err
		}
		smtpPort = port
	}

	// Parse SMTP enabled flag
	smtpEnabled := false
	if enabledStr := os.Getenv("SMTP_ENABLED"); enabledStr != "" {
		enabled, err := strconv.Atoi(enabledStr)
		if err != nil {
			return Config{}, err
		}
		smtpEnabled = enabled > 0
	}

	// Get required environment variables or return error if missing/blank
	appName, err := getEnvRequired("APP_NAME")
	if err != nil {
		return Config{}, err
	}

	frontendURL, err := getEnvRequired("FRONTEND_URL")
	if err != nil {
		return Config{}, err
	}

	smtpHost, err := getEnvRequired("SMTP_HOST")
	if err != nil {
		return Config{}, err
	}

	smtpUsername, err := getEnvRequired("SMTP_USERNAME")
	if err != nil {
		return Config{}, err
	}

	smtpPassword, err := getEnvRequired("SMTP_PASSWORD")
	if err != nil {
		return Config{}, err
	}

	sendingName, err := getEnvRequired("SENDING_NAME")
	if err != nil {
		return Config{}, err
	}

	sendingAddress, err := getEnvRequired("SENDING_ADDRESS")
	if err != nil {
		return Config{}, err
	}

	return Config{
		AppName:     appName,
		FrontendURL: frontendURL,

		SMTPEnabled:    smtpEnabled,
		SMTPHost:       smtpHost,
		SMTPPort:       smtpPort,
		SMTPUsername:   smtpUsername,
		SMTPPassword:   smtpPassword,
		SendingName:    sendingName,
		SendingAddress: sendingAddress,
	}, nil
}

// Helper function to get required environment variable or return error
func getEnvRequired(key string) (string, error) {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		return "", fmt.Errorf("required environment variable %s is not set or blank", key)
	}
	return value, nil
}
