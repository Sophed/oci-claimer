package main

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	DiscordID      string   `json:"discord_id"`
	WebhookURL     string   `json:"webhook_url"`
	KeyPath        string   `json:"ssh_public_Key"`
	Instance       Instance `json:"instance"`
	Domains        []string `json:"availability_domains"`
	NotifyCapacity bool     `json:"notify_out_of_capacity"`
	RetryDelay     int32    `json:"retry_delay"`
	DomainSwitch   int32    `json:"availability_domain_switch_delay"`
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
