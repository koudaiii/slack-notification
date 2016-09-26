package main

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type SlackConfiguration struct {
	SlackWebhookURL string `envconfig:"url" default:"http://localhost/"`
}

const (
	ConfigPrefix = "slack"
)

func LoadConfig() (*SlackConfiguration, error) {
	var config SlackConfiguration
	err := envconfig.Process(ConfigPrefix, &config)

	if err != nil {
		return nil, errors.Wrap(err, "Failed to load config from envs.")
	}

	return &config, nil
}
