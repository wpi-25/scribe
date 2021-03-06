package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/lus/dgc"
	"github.com/wpi-25/scribe/commands"
	"github.com/wpi-25/scribe/db"
	"github.com/wpi-25/scribe/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("WARN: Could not load environment")
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
	router := dgc.Create(&dgc.Router{
		Prefixes: []string{
			os.Getenv("DISCORD_PREFIX"),
		},
		BotsAllowed: false,
		PingHandler: func(c *dgc.Ctx) {
			c.RespondText(fmt.Sprintf("Try `%shelp`!", os.Getenv("DISCORD_PREFIX")))
		},
	})

	router.RegisterMiddleware(middleware.AdminOnly)
	router.RegisterMiddleware(middleware.GuildOwnerOnly)

	router.RegisterDefaultHelpCommand(s, nil)

	commands.MetaCommands(router)
	commands.InviteCommands(router)
	commands.AdminCommands(router)

	// Add command handler to message listener
	router.Initialize(s)

	err = s.Open()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Bot running")
	// Keep the bot running
	<-make(chan struct{})
}
