package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Embed struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Color       int    `json:"color"`
}

type Webhook struct {
	Content string  `json:"content,omitempty"`
	Embeds  []Embed `json:"embeds,omitempty"`
}

func alert(url string, color int, data string) {
	embed := Embed{
		Title:       "Claim Attempt",
		Description: data,
		Color:       color,
	}
	message := Webhook{
		Embeds:  []Embed{embed},
	}
	err := message.send(url)
	if err != nil {
		log.Fatal(err)
	}
}

func (w *Webhook) send(url string) error {

	data, err := json.Marshal(w)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		log.Fatal("Unexpected response: ", resp.Status)
		os.Exit(1)
	}

	return nil
}
