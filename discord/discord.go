package discord

import (
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"

	_logger "github.com/oyakodon/emoji-gen-bot/logger"
)

var logger _logger.Logger = _logger.NewZapLogger()

type DiscordBot struct {
	token   string
	guildId string

	intent discordgo.Intent

	nickname string

	interactions        []func(s *discordgo.Session, i *discordgo.InteractionCreate) error
	commands            []*discordgo.ApplicationCommand
	registerdCommandIds []string
}

func (bot *DiscordBot) AddInteraction(handler func(s *discordgo.Session, i *discordgo.InteractionCreate) error) {
	bot.interactions = append(bot.interactions, handler)
}

func (bot *DiscordBot) AddCommand(cmd *discordgo.ApplicationCommand) {
	bot.commands = append(bot.commands, cmd)
}

func (bot *DiscordBot) SetNickname(nickname string) {
	bot.nickname = nickname
}

// isephemeral: 送信したユーザにのみ表示される
func (bot *DiscordBot) respond(s *discordgo.Session, i *discordgo.Interaction, message string, isephemeral bool, attachment []*discordgo.File) error {
	data := &discordgo.InteractionResponseData{
		Content: message,
	}

	if isephemeral {
		data.Flags = discordgo.MessageFlagsEphemeral
	}

	if len(attachment) > 0 {
		data.Files = attachment
	}

	return s.InteractionRespond(i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: data,
	})
}

func (bot *DiscordBot) Run() error {
	logger.Info("Connecting...")

	s, err := discordgo.New("Bot " + bot.token)
	if err != nil {
		return err
	}

	s.Identify.Intents = bot.intent

	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		logger.Infof("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	// インタラクションを登録
	for _, h := range bot.interactions {
		s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if err := h(s, i); err != nil {
				logger.Warn(err)
			}
		})
	}

	// 接続
	if err := s.Open(); err != nil {
		return err
	}
	defer s.Close()

	// 表示名を変更
	if err := s.GuildMemberNickname(bot.guildId, "@me", bot.nickname); err != nil {
		logger.Infof("failed to set own nickname: %v", err)
	}

	// コマンドの登録
	for _, cmd := range bot.commands {
		_cmd, err := s.ApplicationCommandCreate(s.State.User.ID, bot.guildId, cmd)
		if err != nil {
			return err
		}

		bot.registerdCommandIds = append(bot.registerdCommandIds, _cmd.ID)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	logger.Info("Press Ctrl+C to exit")
	<-stop

	logger.Info("Shutting down...")

	// 表示名をリセット
	if err := s.GuildMemberNickname(bot.guildId, "@me", ""); err != nil {
		logger.Infof("failed to set own nickname: %v", err)
	}

	// コマンドの登録解除
	for _, id := range bot.registerdCommandIds {
		if err := s.ApplicationCommandDelete(s.State.User.ID, bot.guildId, id); err != nil {
			return err
		}
	}

	logger.Info("done!")

	return nil
}

func NewDiscordBot(bottoken, guildId string, intent discordgo.Intent) *DiscordBot {
	return &DiscordBot{
		token:   bottoken,
		guildId: guildId,
		intent:  intent,
	}
}
