package db

type GuildSettings struct {
	GuildId        string `db:"guild_id"`
	ChannelLock    bool   `db:"channel_lock"`
	ChannelLockId  string `db:"channel_lock_id"`
	TempLock       bool   `db:"temporary_lock"`
	ReasonRequired bool   `db:"require_reason"`
	MinRoleId      string `db:"min_role_id"`
	LogChannel     string `db:"min_role_id"`
}
