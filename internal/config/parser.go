package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LoadConfig(path string) (Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	config := make(Config)
	var currentSection string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}

		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			currentSection = strings.ToLower(strings.Trim(line, "[]"))
			if _, exists := config[currentSection]; !exists {
				config[currentSection] = make(map[string]string)
			}
			continue
		}

		// Parse key-value
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// If no section defined yet, use a default one
		if currentSection == "" {
			currentSection = "default"
			if _, exists := config[currentSection]; !exists {
				config[currentSection] = make(map[string]string)
			}
		}

		config[currentSection][key] = value
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return config, nil
}
