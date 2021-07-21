package middleware

import (
	"database/sql"
	"log"

	"github.com/Necroforger/dgrouter/exrouter"
	"github.com/wpi-25/scribe/db"
)

func AdminOnly(fn exrouter.HandlerFunc) exrouter.HandlerFunc {
	return func(ctx *exrouter.Context) {
		log.Println("Attempting to authenticate an admin command")

		row := db.DB.QueryRow("SELECT min_role_id FROM guild_settings WHERE guild_id = ?", ctx.Msg.GuildID)
		var min_role_id sql.NullString
		err := row.Scan(min_role_id)
		if err != nil {
			ctx.Reply("Could not Scan: %s", err)
		}
		err = row.Err()
		if err != nil {
			guild, _ := ctx.Guild(ctx.Msg.GuildID)
			owner := guild.OwnerID
			if owner == ctx.Msg.Author.ID {
				fn(ctx)
			} else {
				ctx.Reply("Could not authenticate you! Have the guild owner set an admin role.")
			}
		}
	}
}
