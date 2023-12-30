package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

func main() {
	sess, err := discordgo.New("Bot MTE4MDUwNjA2NTAyMDkxMTczNw.GB8a0L.CQDRbsixbj6ThKWNFDJfZpCr6cIMs1KOp0-f1Y")
	if err != nil {
		log.Fatal(err)
	}
	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}
		if m.Content == "!hello" {
			s.ChannelMessageSend(m.ChannelID, "Hello, I can go.")
		}
		if m.Content == "!ping" {
			before := time.Now()
			s.ChannelMessageSend(m.ChannelID, "Pinging!")
			after := time.Now()
			duration := after.Sub(before)
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Bot Latency: %v", duration))
		}
		if m.Content == "!help" {
			em := &discordgo.MessageEmbed{
				Title:       "Help Panel",
				Description: "**!hello\n!ping\n!help**",
				Color:       0x2e2e2e,
				Footer: &discordgo.MessageEmbedFooter{
					Text: "Made with DiscordGO",
				},
			}
			s.ChannelMessageSendEmbed(m.ChannelID, em)
		}
	})
	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
	err = sess.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()
	fmt.Println("Bot is on")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
