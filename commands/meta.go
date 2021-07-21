package commands

import (
	"log"

	"github.com/Necroforger/dgrouter/exrouter"
)

func MetaCommands(router *exrouter.Route) {
	log.Println("Initializing Meta Commands")
	router.Group(func(r *exrouter.Route) {
		router.Cat("Meta")
		router.On("ping", PingCommand).Desc("Pings the bot")

		// Help
		router.Default = router.On("help", func(ctx *exrouter.Context) {
			var text = ""

			for _, v := range router.Routes {
				text += v.Name + ": \t" + v.Description + "\n"
			}
			ctx.Reply(text)
		}).Desc("Displays this help message")
	})
}

func PingCommand(ctx *exrouter.Context) {
	ctx.Reply("Pong!")
}
