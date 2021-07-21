package commands

import (
	"log"

	"github.com/Necroforger/dgrouter/exrouter"
)

func MetaCommands(router *exrouter.Route) {
	log.Println("Initializing Meta Commands")
	router.On("ping", PingCommand).Desc("Pings the bot")
}

func PingCommand(ctx *exrouter.Context) {
	ctx.Reply("Pong!")
}
