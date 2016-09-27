package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if os.Getenv("SLACK_WEBHOOK_URL") == "" {
		fmt.Fprintln(os.Stderr, "Please provide 'SLACK_WEBHOOK_URL' through environment")
		os.Exit(1)
	}
}

func main() {
	//Incoming Webhook URL
	LoadConfig()

	// input send text
	args := strings.Join(os.Args[1:], " ")
	if args == "" {
		fmt.Fprintln(os.Stdout, "usage: slack-notifier <TEXT>")
		os.Exit(0)
	}

	//Incoming Webhook URL
	config, err := LoadConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	//Form JSON payload to send to Slack
	json := `{"text": "` + args + `"}`
	//Post JSON payload to the Webhook URL
	client := http.Client{}
	req, err := http.NewRequest("POST", config.SlackWebhookURL, bytes.NewBufferString(json))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/json")
	_, err = client.Do(req)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	req.Body.Close()
}
