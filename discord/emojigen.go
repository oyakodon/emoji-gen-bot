package discord

import (
	"bytes"

	"github.com/bwmarrin/discordgo"

	"github.com/oyakodon/emoji-gen-bot/emoji"
)

const (
	CommandName        = "emoji"
	CommandDescription = "送信されたテキストから絵文字を生成して返します。"

	MsgUnexpectedOperation = "このメッセージは表示されないはずだよ！"
	MsgRestrictedChannel   = "使用できるチャンネルが設定で制限されています... :man_bowing:"

	SubcommandTest = "test"
)

type EmojiGenBot struct {
	*DiscordBot

	channels []string
}

// isephemeral: 送信したユーザにのみ表示される
func (b EmojiGenBot) respond(s *discordgo.Session, i *discordgo.Interaction, message string, isephemeral bool, attachment []*discordgo.File) error {
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

func (b EmojiGenBot) interaction(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	respond := func(message string) error {
		return b.respond(s, i.Interaction, message, true, nil)
	}

	respondWithPNGImage := func(message string, data []byte) error {
		return b.respond(s, i.Interaction, message, true, []*discordgo.File{
			{
				Name:        "emoji.png",
				ContentType: "image/png",
				Reader:      bytes.NewReader(data),
			},
		})
	}

	// 別サーバの際は実行しない
	if i.GuildID != b.guildId {
		return respond(MsgUnexpectedOperation)
	}

	// configのChannelIDに含まれていなくても実行しない
	ok := false
	for _, c := range b.channels {
		if ok = ok || c == i.ChannelID; ok {
			break
		}
	}
	if !ok {
		return respond(MsgRestrictedChannel)
	}

	// サブコマンドで分岐
	options := i.ApplicationCommandData().Options[0]
	switch options.Name {
	case SubcommandTest:
		b, err := emoji.GenerateEmoji(
			"絵文\n字！",
			0xFF008080,
			emoji.EmojiColorTransparent,
			emoji.EmojiAlignCenter,
		)
		if err != nil {
			return err
		}

		return respondWithPNGImage("OK!", b)
	}

	return nil
}

func NewEmojiGenBot(
	bottoken, guildId, nickname string,
	channels []string,
) *EmojiGenBot {
	bot := &EmojiGenBot{
		DiscordBot: NewDiscordBot(
			bottoken,
			guildId,
			discordgo.IntentGuildMessages, // サーバメッセージ送信
		),
		channels: channels,
	}

	bot.SetNickname(nickname)
	bot.AddCommand(&discordgo.ApplicationCommand{
		Name:        CommandName,
		Description: CommandDescription,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "test",
				Description: "(デバッグ用)",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
			},
		},
	})
	bot.AddInteraction(bot.interaction)

	return bot
}
