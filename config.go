package main

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	DiscordID  string   `json:"discord_id"`
	WebhookURL string   `json:"webhook_url"`
	KeyPath    string   `json:"ssh_public_Key"`
	Instance   Instance `json:"instance"`
}

func load_config(path string) Config {
	var config Config
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}
