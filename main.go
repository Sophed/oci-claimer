package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {

	config := load_config("config.json")

	key, err := os.ReadFile(config.KeyPath)
	if err != nil {
		log.Fatal(err)
	}
	config.Instance.SSHPublicKey = string(key)

	for n, s := range config.Domains {
		attempt(config, s)
		if n != len(config.Domains)-1 {
			time.Sleep(time.Second * 5)
		}
	}

}

func attempt(config Config, domain string) {
	config.Instance.Domain = domain
	err := config.Instance.claim()
	if err != nil {

		if strings.Contains(fmt.Sprint(err), "Out of host capacity.") {
			if config.NotifyCapacity {
				alert(config.WebhookURL, 0xff0000, "", "Out of capacity in domain: "+config.Instance.Domain)
			}
			fmt.Println("Out of capacity.")
			return
		}

		// you're on your own here, good luck
		alert(config.WebhookURL, 0xff0000, "<@"+config.DiscordID+">", "Unknown error occurred, check terminal output for more details.")
		fmt.Println(err)
	} else {
		alert(config.WebhookURL, 0x00ff00, "<@"+config.DiscordID+">", "Possible instance claimed! Check OCI panel.")
	}
}
