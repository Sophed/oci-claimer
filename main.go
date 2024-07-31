package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
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

	for {
		for n, s := range config.Domains {
			if attempt(config, s) {
				return
			}
			if n != len(config.Domains)-1 {
				time.Sleep(time.Second * time.Duration(config.DomainSwitch))
			}
		}
		log.Println("Retrying in " + strconv.Itoa(int(config.RetryDelay)) + " seconds...")
		time.Sleep(time.Second * time.Duration(config.RetryDelay))
	}

}

func attempt(config Config, domain string) bool {
	config.Instance.Domain = domain
	err := config.Instance.claim()
	if err != nil {

		if strings.Contains(fmt.Sprint(err), "Out of host capacity.") {
			if config.NotifyCapacity {
				alert(config.WebhookURL, 0xff0000, "", "Out of capacity in domain: "+config.Instance.Domain)
			}
			log.Println("Out of capacity in " + config.Instance.Domain)
			return false
		}

		// you're on your own here, good luck
		alert(config.WebhookURL, 0xff0000, "<@"+config.DiscordID+">", "Unknown error occurred, check terminal output for more details.")
		fmt.Println(err)
		return false
	} else {
		alert(config.WebhookURL, 0x00ff00, "<@"+config.DiscordID+">", "Possible instance claimed! Check OCI panel.")
		return true
	}
}
