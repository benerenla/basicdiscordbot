package main

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"discordbot.com/commands"
	"github.com/bwmarrin/discordgo"
)

var (
	Token   = "YOUR TOKEN"
	BotID   string
	prefix  = "YOUR PREFİX"
	connStr = "postgresql://postgres::root@postgres/postgres?sslmode=disable"
)

func main() {

	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	fmt.Println("Başarıyla Açıldı ")
	u, err := dg.User("@me")
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	BotID = u.ID

	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	<-sc
}
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "" {
		return
	}
	if strings.Count(m.Content, prefix) < 2 {
		commands.ExecuteCommand(s, m.Message, prefix)
	}
}
