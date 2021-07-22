package commands

import (
	"log"

	"github.com/lus/dgc"
)

func MetaCommands(router *dgc.Router) {
	log.Println("Initializing Meta Commands")
	router.RegisterCmd(&dgc.Command{
		Name:    "ping",
		Handler: PingCommand,
	})
}

func PingCommand(ctx *dgc.Ctx) {
	ctx.RespondText("Pong!")
}
