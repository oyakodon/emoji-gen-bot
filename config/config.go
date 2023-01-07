package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type BotConfig struct {
	BotToken string `yaml:"bottoken"`

	GuildId  string   `yaml:"guildid"`
	Channels []string `yaml:"channels"`

	Nickname string `yaml:"nickname"`
}

func LoadConfig(path string) (*BotConfig, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var conf BotConfig
	if err := yaml.Unmarshal(b, &conf); err != nil {
		return nil, err
	}

	return &conf, err
}
