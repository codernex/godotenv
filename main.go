// env/env.go
package env

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Load reads the .env file and loads the variables into the environment
func Load(filename string) error {
	if filename == "" {
		filename = ".env"
	}

	absPath, err := filepath.Abs(filename)
	if err != nil {
		return fmt.Errorf("error getting absolute path: %v", err)
	}

	file, err := os.Open(absPath)
	if err != nil {
		return fmt.Errorf("error opening .env file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Split into key-value pair
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue // or return error if you prefer
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Remove surrounding quotes if present
		if len(value) > 0 && (value[0] == '"' || value[0] == '\'') {
			value = value[1 : len(value)-1]
		}

		// Set the environment variable
		if err := os.Setenv(key, value); err != nil {
			return fmt.Errorf("error setting env var %s: %v", key, err)
		}
	}

	return scanner.Err()
}

// Get returns the value of the environment variable or a default if not found
func Get(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// MustGet returns the value of the environment variable or panics if not found
func MustGet(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	panic(fmt.Sprintf("environment variable %s not found", key))
}
