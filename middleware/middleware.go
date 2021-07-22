package middleware

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/lus/dgc"
	"github.com/wpi-25/scribe/db"
)

func AdminOnly(next dgc.ExecutionHandler) dgc.ExecutionHandler {
	return func(c *dgc.Ctx) {
		is_admin := false
		for _, flag := range c.Command.Flags {
			if flag == "admin" {
				is_admin = true
			}
		}
		if is_admin {
			log.Println("Attempting to authenticate an admin command")

			row := db.DB.QueryRow("SELECT min_role_id FROM guild_settings WHERE guild_id = ?", c.Event.GuildID)
			var min_role_id sql.NullString
			err := row.Scan(min_role_id)
			if err != nil {
				c.RespondText(fmt.Sprintf("Could not Scan: %s", err))
				return
			}
			err = row.Err()
			if err != nil {
				guild, _ := c.Session.Guild(c.Event.GuildID)
				owner := guild.OwnerID
				if owner == c.Event.Author.ID {
					next(c)
				} else {
					c.RespondText("Could not authenticate you! Have the server owner set an admin role.")
				}
			} else {
				memberRoles := c.Event.Member.Roles
				for _, role := range memberRoles {
					if role == min_role_id.String {
						next(c)
						return
					}
				}
				c.RespondText("You don't have permissions!")
			}
		} else {
			next(c)
		}
	}

}

func GuildOwnerOnly(next dgc.ExecutionHandler) dgc.ExecutionHandler {
	return func(c *dgc.Ctx) {
		is_owner := false
		for _, flag := range c.Command.Flags {
			if flag == "owner" {
				is_owner = true
			}
		}

		if is_owner {
			log.Println("Attempting to authenticate a guild owner command")
			guild, _ := c.Session.Guild(c.Event.GuildID)
			owner := guild.OwnerID
			if owner == c.Event.Author.ID {
				next(c)
			} else {
				c.RespondText("You aren't the owner! Please ask them to run this command for you.")
			}
		} else {
			next(c)
		}
	}
}
