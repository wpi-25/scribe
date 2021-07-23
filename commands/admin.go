package commands

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/lus/dgc"
	"github.com/wpi-25/scribe/db"
)

func AdminCommands(r *dgc.Router) {
	log.Println("Initializing Admin Commands")
	r.RegisterCmd(&dgc.Command{
		Name: "void",

		Description: "Voids an invite",
		Usage:       "void <code>",
		Example:     "void djgHW23",
		Handler:     voidInvite,

		Flags: []string{
			"admin",
		},
	})

	r.RegisterCmd(&dgc.Command{
		Name: "minperms",

		Description: "Sets the minimum permissions for admin actions",
		Usage:       "minperms <@role>",
		Example:     "minperms @Admin",
		Handler:     minPerms,

		Flags: []string{
			"owner",
		},
	})

	r.RegisterCmd(&dgc.Command{
		Name: "listinvites",

		Description: "Lists managed invites",
		Usage:       "listinvites",
		Handler:     listInvites,

		Flags: []string{
			"admin",
		},
	})
}

func voidInvite(ctx *dgc.Ctx) {

	args := ctx.Arguments

	code := args.Get(0)
	ctx.RespondText(code.Raw())

	_, err := ctx.Session.InviteDelete(code.Raw())
	if err != nil {
		ctx.RespondText(fmt.Sprintf("Could not delete invite: %s", err))
	} else {
		ctx.RespondText("Done")
	}
}

func minPerms(ctx *dgc.Ctx) {
	args := ctx.Arguments

	roleId := args.Get(1).AsRoleMentionID()

	ctx.RespondText(roleId)

	row := db.DB.QueryRowx("SELECT guild_id FROM guild_settings WHERE guild_id = $1", ctx.Event.GuildID)
	var guild_id string
	err := row.Scan(&guild_id)
	if err != nil {
		if err == sql.ErrNoRows {
			settings := db.GuildSettings{
				GuildId:   ctx.Event.GuildID,
				MinRoleId: roleId,
			}
			_, err := db.DB.Exec("INSERT INTO guild_settings (guild_id, min_role_id) VALUES ($1, $2)", settings.GuildId, settings.MinRoleId)
			if err != nil {
				ctx.RespondText(fmt.Sprintf("Could not update settings: %s", err))
				log.Println(err)
				return
			}
		}
	} else {
		_, err := db.DB.Exec("UPDATE guild_settings SET min_role_id=$1 WHERE guild_id=$2", roleId, ctx.Event.GuildID)
		if err != nil {
			ctx.RespondText(fmt.Sprintf("Could not update settings: %s", err))
			return
		}
	}
	ctx.RespondText("Done.")
}

func listInvites(c *dgc.Ctx) {
	rows, err := db.DB.Queryx("SELECT * FROM invites WHERE guild_id = $1", c.Event.GuildID)
	if err != nil {
		c.RespondText(fmt.Sprintf("Could not get a list of invite reasons."))
	}

	var invites []InviteWReason

	for rows.Next() {
		var invite InviteWReason
		err = rows.StructScan(&invite)
	}

	text := "**Invites**"

	for _, i := range invites {
		log.Printf("%s\t <@%s>, <#%s>, %s", i.Code, i.UserCreated, i.Channel, i.Reason)
		text += fmt.Sprintf("%s\t <@%s>, <#%s>, %s", i.Code, i.UserCreated, i.Channel, i.Reason)
	}
	c.RespondText(text)
}

type InviteWReason struct {
	Code        string
	GuildID     string `db:"guild_id"`
	CreatedAt   string `db:"created_at"`
	UserCreated string `db:"user_created"`
	Channel     string `db:"target_channel"`
	MaxUses     int    `db:"max_uses"`
	Reason      string
}
