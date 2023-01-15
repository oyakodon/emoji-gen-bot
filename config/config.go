package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type BotConfigFont struct {
	Id   string `yaml:"id"`
	Name string `yaml:"name"`
	Path string `yaml:"path"`
}

type BotConfig struct {
	BotToken string `yaml:"bottoken"`

	GuildId  string   `yaml:"guildid"`
	Channels []string `yaml:"channels"`

	NoticeChannelId string `yaml:"noticechannelid"`
	Nickname        string `yaml:"nickname"`

	Fonts []BotConfigFont `yaml:"fonts"`
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
