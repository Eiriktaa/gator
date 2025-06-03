package main

import (
	"log"

	"example.com/eiriktaa/gator/internal/config"
)

func main() {
	cfg := config.LoadConfiguration()
	cfg.PrintSelf()
	cfg.Current_user_name = "test"
	err := cfg.WriteToFile()
	if err != nil {
		log.Fatalf("Failed to write config %v", err)
	}
}
