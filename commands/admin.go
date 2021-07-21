package commands

import (
	"fmt"
	"log"

	"github.com/Necroforger/dgrouter/exrouter"
	"github.com/wpi-25/scribe/middleware"
)

func AdminCommands(r *exrouter.Route) {
	log.Println("Initializing Admin Commands")
	r.Use(middleware.AdminOnly)
	r.On("void", voidInvite).Desc("Deletes an invite")
}

func voidInvite(ctx *exrouter.Context) {
	args := exrouter.ParseArgs(ctx.Msg.Content)

	code := args.Get(1)

	_, err := ctx.Ses.InviteDelete(code)
	if err != nil {
		ctx.Reply(fmt.Sprintf("Could not delete invite: %s", err))
	} else {
		ctx.Reply("Done")
	}
}
