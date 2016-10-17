package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	if _, err := os.Stat(".env"); err == nil {
		if err2 := godotenv.Load(); err2 != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	//Incoming Webhook URL
	if os.Getenv("SLACK_WEBHOOK_URL") == "" {
		fmt.Fprintln(os.Stderr, "Please provide 'SLACK_WEBHOOK_URL' through environment")
		os.Exit(1)
	}

	// input send text
	args := strings.Join(os.Args[1:], " ")
	if args == "" {
		fmt.Fprintln(os.Stdout, "usage: slack-notifier <TEXT>")
		os.Exit(0)
	}

	//Form JSON payload to send to Slack
	json := `{"text": "` + args + `"}`
	//Post JSON payload to the Webhook URL
	client := http.Client{}
	req, err := http.NewRequest("POST", os.Getenv("SLACK_WEBHOOK_URL"), bytes.NewBufferString(json))
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
