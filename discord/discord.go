package discord

import (
	"encoding/base64"
	"errors"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"

	_logger "github.com/oyakodon/emoji-gen-bot/logger"
)

var logger _logger.Logger = _logger.NewZapLogger()

const (
	CONTENT_TYPE_IMAGE_PNG = "image/png"
)

var (
	MAX_SIZE_EMOJI = 256 * 1024 // 256 KB = 256 * 1024 B

	ErrExceedEmojiSize = errors.New("exceed emoji size: >=256KB")
)

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

func (bot *DiscordBot) MaxEmojiCount(s *discordgo.Session) int {
	m := map[discordgo.PremiumTier]int{
		discordgo.PremiumTierNone: 50,
		discordgo.PremiumTier1:    100,
		discordgo.PremiumTier2:    150,
		discordgo.PremiumTier3:    250,
	}

	g, _ := s.Guild(bot.guildId)
	return m[g.PremiumTier]
}

func (bot *DiscordBot) EmojiCount(s *discordgo.Session) int {
	emojis, _ := s.GuildEmojis(bot.guildId)
	return len(emojis)
}

func (bot *DiscordBot) CreateEmoji(s *discordgo.Session, name string, data []byte) (string, error) {
	image := base64.StdEncoding.EncodeToString(data)
	if len(image) >= MAX_SIZE_EMOJI {
		return "", ErrExceedEmojiSize
	}

	e, err := s.GuildEmojiCreate(bot.guildId, &discordgo.EmojiParams{
		Name:  name,
		Image: "data:" + CONTENT_TYPE_IMAGE_PNG + ";base64," + image,
	})
	if err != nil {
		return "", err
	}

	return e.MessageFormat(), nil
}

// isephemeral: 送信したユーザにのみ表示される
func (bot *DiscordBot) Respond(s *discordgo.Session, i *discordgo.Interaction, message string, isephemeral bool, attachment []*discordgo.File) error {
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

// チャンネルにメッセージを送信する
func (bot *DiscordBot) SendMessage(s *discordgo.Session, channelid, message string) error {
	_, err := s.ChannelMessageSend(channelid, message)
	return err
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
