package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	//load the env to get the token
	err := godotenv.Load()
	if err != nil {
		log.Printf("No .env file found")
	}
	botToken := os.Getenv("DISCORD_BOT_TOKEN")

	discord, err := discordgo.New("Bot " + (botToken))
	if err != nil {
		log.Fatal(err)
	}

	discord.AddHandler(func(session *discordgo.Session, message *discordgo.MessageCreate) {

		//checks if bot is the author so it doesnt get stuck in a loop
		if message.Author.ID == session.State.User.ID {
			return
		}

		if message.Content == "Eat cement" {
			session.ChannelMessageSend(message.ChannelID, "You have eaten some cement!")
		}
	})

	discord.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = discord.Open()
	if err != nil {
		log.Fatal(err)
	}

	logger := log.New(os.Stdout, "", log.Flags())

	logger.Printf("We <3 Cement")
	logger.Printf("Bot is on and working! CTRL-C to exit.")

	//makes it so you can shutdown the bot (sad and mean and evil)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	discord.Close()
	logger.Printf("Shutting down... goodnight...")
}
