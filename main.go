package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

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
			log.Fatal("Out of capacity.")
		}

		// you're on your own here, good luck
		alert(config.WebhookURL, 0xff0000, "<@"+config.DiscordID+">", "Unknown error occurred, check terminal output for more details.")
		log.Fatal(err)
	}

	alert(config.WebhookURL, 0x00ff00, "<@"+config.DiscordID+">", "Possible instance claimed! Check OCI panel.")

}
