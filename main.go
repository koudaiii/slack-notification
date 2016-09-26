package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	// input send text
	args := strings.Join(os.Args[1:], " ")
	if args == "" {
		fmt.Println("no text")
		os.Exit(0)
	}

	//Incoming Webhook URL
	config, err := LoadConfig()
	if err != nil {
		fmt.Println("%v", err)
		os.Exit(1)
	}

	//Form JSON payload to send to Slack
	json := `{"text": "` + args + `"}`
	//Post JSON payload to the Webhook URL
	client := http.Client{}
	req, err := http.NewRequest("POST", config.SlackWebhookURL, bytes.NewBufferString(json))
	if err != nil {
		fmt.Println("Unable to parse slack webhook url.")
	}
	req.Header.Set("Content-Type", "application/json")
	_, err = client.Do(req)
	if err != nil {
		fmt.Println("Unable to reach the server.")
	}
	req.Body.Close()
}
