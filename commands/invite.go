package commands

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/lus/dgc"
)

func InviteCommands(r *dgc.Router) {
	log.Println("Initializing Invite Commands")
	r.RegisterCmd(&dgc.Command{
		Name: "invite",

		Description: "Creates an invite to the server",

		Handler: createInviteCommand,
	})
}

func createInviteCommand(ctx *dgc.Ctx) {
	invite_settings := discordgo.Invite{
		Temporary: true,
	}

	invite, err := ctx.Session.ChannelInviteCreate(ctx.Event.ChannelID, invite_settings)

	if err != nil {
		ctx.RespondText(fmt.Sprintf("Could not create invite: %s", err))
	} else {
		ctx.RespondText(fmt.Sprintf("https://discord.gg/%s", invite.Code))
	}
}
