package commands

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Necroforger/dgrouter/exrouter"
	a "github.com/wpi-25/scribe/args"
	"github.com/wpi-25/scribe/db"
	"github.com/wpi-25/scribe/middleware"
)

func AdminCommands(r *exrouter.Route) {
	log.Println("Initializing Admin Commands")
	r.Use(middleware.AdminOnly)
	r.On("void", voidInvite).Desc("Deletes an invite")
	r.Group(func(r *exrouter.Route) {
		r.Cat("Owner")
		//r.Use(middleware.GuildOwnerOnly)
		r.On("minperms", minPerms).Desc("Sets the minimum role needed for admin commands. Server Owner Only")
	})
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

func minPerms(ctx *exrouter.Context) {
	args := exrouter.ParseArgs(ctx.Msg.Content)
	roleId, err := a.ParseRoleIDFromMention(args.Get(1))
	if err != nil {
		ctx.Reply(err)
		return
	}

	ctx.Reply(roleId)

	row := db.DB.QueryRowx("SELECT guild_id FROM guild_settings WHERE guild_id = $1", ctx.Msg.GuildID)
	var guild_id string
	err = row.Scan(&guild_id)
	if err != nil {
		if err == sql.ErrNoRows {
			settings := db.GuildSettings{
				GuildId:   ctx.Msg.GuildID,
				MinRoleId: roleId,
			}
			_, err := db.DB.Exec("INSERT INTO guild_settings (guild_id, min_role_id) VALUES ($1, $2)", settings.GuildId, settings.MinRoleId)
			if err != nil {
				ctx.Reply(fmt.Sprintf("Could not update settings: %s", err))
				log.Println(err)
			}
		}
	}

}
