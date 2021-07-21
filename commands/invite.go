package commands

import (
	"fmt"
	"log"

	"github.com/Necroforger/dgrouter/exrouter"
	"github.com/bwmarrin/discordgo"
)

func InviteCommands(r *exrouter.Route) {
	log.Println("Initializing Invite Commands")
	r.On("createinvite", createInviteCommand).Desc("Creates an invite with the specified parameters")
}

func createInviteCommand(ctx *exrouter.Context) {
	invite_settings := discordgo.Invite{
		Temporary: true,
	}

	invite, err := ctx.Ses.ChannelInviteCreate(ctx.Msg.ChannelID, invite_settings)

	if err != nil {
		ctx.Reply(fmt.Sprintf("Could not create invite: %s", err))
	} else {
		ctx.Reply(fmt.Sprintf("https://discord.gg/%s", invite.Code))
	}
}
