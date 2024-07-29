package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type Config struct {
	DiscordID string `json:"discord_id"`
	WebhookURL string `json:"webhook_url"`
	KeyPath string `json:"ssh_public_Key"`
	Instance Instance `json:"instance"`
}

func main() {

	config := load_config("config.json")

	key, err := os.ReadFile(config.KeyPath)
    if err != nil {
        log.Fatal(err)
    }
	config.Instance.SSHPublicKey = string(key)

	err = config.Instance.claim()
	if err != nil {

		if strings.Contains(fmt.Sprint(err), "Out of host capacity.") {
			alert(config.WebhookURL, 0xff0000, "", "Out of capacity in domain: "+config.Instance.Domain)
			return
		}

		// you're on your own here, good luck
		fmt.Println(err)
		return
	}

	alert(config.WebhookURL, 0x00ff00, "<@"+config.DiscordID+">", "Possible instance claimed! Check OCI panel.")
	
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