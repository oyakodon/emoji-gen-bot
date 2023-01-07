package main

import (
	"github.com/oyakodon/emoji-gen-bot/config"
	"github.com/oyakodon/emoji-gen-bot/discord"
	_logger "github.com/oyakodon/emoji-gen-bot/logger"
)

const (
	CONFIG_PATH = `./config/config.yml`
)

var logger _logger.Logger = _logger.NewZapLogger()

func main() {
	conf, err := config.LoadConfig(CONFIG_PATH)
	if err != nil {
		logger.Fatal(err)
		return
	}

	bot := discord.NewEmojiGenBot(
		conf.BotToken,
		conf.GuildId,
		conf.Nickname,
		conf.Channels,
	)

	if err := bot.Run(); err != nil {
		logger.Warn(err)
	}
}
