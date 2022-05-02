package commands

import (
	"github.com/bwmarrin/discordgo"
)

func HandleInfoCommand(s *discordgo.Session, m *discordgo.Message) {

	var (
		message = "sdasd"
	)

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "Selam",
		Description: message,
	})
}
