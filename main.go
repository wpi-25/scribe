package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/wpi-25/scribe/commands"
	"github.com/wpi-25/scribe/db"

	"github.com/Necroforger/dgrouter/exrouter"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Could not load environment")
	}
	// Initialize Bot
	s, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Fatalln(err)
	}

	// Initialize Database
	err = db.Connect()
	if err != nil {
		log.Fatalln(fmt.Sprintf("Could not connect to database %s", err))
	}

	// Initialize Command Router
	router := exrouter.New()

	commands.MetaCommands(router)

	// Add command handler to message listener
	s.AddHandler(func(_ *discordgo.Session, m *discordgo.MessageCreate) {
		router.FindAndExecute(s, os.Getenv("DISCORD_PREFIX"), s.State.User.ID, m.Message)
	})

	err = s.Open()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Bot running")
	// Keep the bot running
	<-make(chan struct{})
}
