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
	logger := log.New(os.Stdout, "", log.Flags())
	logger.Printf("We <3 Cement")

	//load the env to get the token
	err := godotenv.Load()
	if err != nil {
		logger.Printf("No .env file found")
	}

	token, found := os.LookupEnv("DISCORD_BOT_TOKEN")
	if !found {
		logger.Printf("Please provide the bot token")
		os.Exit(1)
	}

	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		logger.Printf("Failed to create discord bot: %s", err)
		os.Exit(1)
	}

	logger.Printf("Bot is waking up!")

	discord.AddHandler(func(session *discordgo.Session, message *discordgo.MessageCreate) {

		//checks if bot is the author so it doesnt get stuck in a loop
		if message.Author.ID == session.State.User.ID {
			return
		}

		session.ChannelMessageSend(message.ChannelID, "You have eaten some cement!")
	})

	discord.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = discord.Open()
	if err != nil {
		logger.Printf("Failed to create start bot: %s", err)
		os.Exit(1)
	}
	defer discord.Close()

	logger.Printf("The bot drank its coffe and is now awake :)")
	logger.Printf("Press CTRL+C to put him to bed")

	// Makes it so you can shutdown the bot (sad and mean and evil)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	logger.Printf("The bot went to bed zZzZzZ")
}
