package commands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func ExecuteCommand(s *discordgo.Session, m *discordgo.Message, prefix string) {
	msg := strings.Split(strings.TrimSpace(m.Content), prefix)[1]
	if len(msg) > 2 {
		msg = strings.Split(strings.Split(m.Content, " ")[0], prefix)[1]
	}
	switch msg {
	case "info":
		HandleInfoCommand(s, m)
	default:
		return
	}
}
