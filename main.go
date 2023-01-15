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

	fonts := make([]*discord.EmojiGenBotFont, 0)
	for _, f := range conf.Fonts {
		fonts = append(fonts, &discord.EmojiGenBotFont{
			Id:       f.Id,
			Name:     f.Name,
			FontPath: f.Path,
		})
	}

	bot := discord.NewEmojiGenBot(
		conf.BotToken,
		conf.GuildId,
		conf.Nickname,
		conf.NoticeChannelId,
		conf.Channels,
		fonts,
	)

	if err := bot.Run(); err != nil {
		logger.Warn(err)
	}
}
