package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
}

type ServerConfig struct {
	Host string
	Port string
}

func main() {
	file, err := os.Open("config.ini")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var config Config
	var currentSection string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}

		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			currentSection = strings.ToLower(strings.Trim(line, "[]"))
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch currentSection {
		case "database":
			switch key {
			case "host":
				config.Database.Host = value
			case "port":
				config.Database.Port = value
			case "user":
				config.Database.User = value
			case "password":
				config.Database.Password = value
			}
		case "server":
			switch key {
			case "host":
				config.Server.Host = value
			case "port":
				config.Server.Port = value
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Printf("Parsed Config: %+v\n", config)
	fmt.Println("Database User:", config.Database.User)
	fmt.Println("Server Port:", config.Server.Port)
}
