package main

import (
	"fmt"
	"log"
	"myapp/internal/config"
)

func main() {
	cfg, err := config.LoadConfig("configs/config.ini")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database Host:", cfg["database"]["host"])
	fmt.Println("Server Port:", cfg["server"]["port"])

	for section, kv := range cfg {
		fmt.Println("[" + section + "]")
		for k, v := range kv {
			fmt.Printf("%s = %s\n", k, v)
		}
	}
}
