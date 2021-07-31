package commands

import (
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/lus/dgc"
	"github.com/wpi-25/scribe/db"
)

func InviteCommands(r *dgc.Router) {
	log.Println("Initializing Invite Commands")
	r.RegisterCmd(&dgc.Command{
		Name: "invite",

		Description: "Creates an invite to the server",
		Usage:       "invite <#channel> <duration> <uses> <reason>",
		Example:     "invite #welcome 1h30m 1 Single-use for new member",

		Handler: createInviteCommand,
	})
}

func createInviteCommand(ctx *dgc.Ctx) {
	args := ctx.Arguments

	channel, err := ctx.Session.Channel(args.Get(0).AsChannelMentionID())
	if err != nil {
		ctx.RespondText("Invalid channel!")
		return
	}
	args.Remove(0)
	duration, err := time.ParseDuration(args.Get(0).Raw())
	if err != nil {
		ctx.RespondText("Invalid duration! Should be of the format 1h30m (1 hour 30 minutes)")
	}
	args.Remove(0)
	uses, err := args.Get(0).AsInt()
	if err != nil {
		ctx.RespondText("Invalid number of uses, should be a number.")
	}
	args.Remove(0)
	reason := args.Raw()

	invite_settings := discordgo.Invite{
		Channel:   channel,
		MaxAge:    int(duration.Seconds()),
		MaxUses:   uses,
		Temporary: false,
	}

	invite, err := ctx.Session.ChannelInviteCreate(ctx.Event.ChannelID, invite_settings)

	if err != nil {
		ctx.RespondText(fmt.Sprintf("Could not create invite: %s", err))
	} else {
		ctx.RespondText(fmt.Sprintf("https://discord.gg/%s", invite.Code))
	}

	created_at, err := invite.CreatedAt.Parse()
	if err != nil {
		created_at = time.Now()
	}

	_, err = db.DB.Exec("INSERT INTO invites (code, guild_id, created_at, user_created, target_channel, max_uses, reason) VALUES ($1, $2, $3, $4, $5, $6, $7)", invite.Code, invite.Guild.ID, created_at, invite.Inviter.ID, invite.Channel.ID, invite.MaxUses, reason)
	if err != nil {
		ctx.RespondText(fmt.Sprintf("Could not store the invite in the database: %s", err))
	}
}
