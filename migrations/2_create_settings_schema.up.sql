CREATE TABLE guild_settings (
    guild_id VARCHAR(20) PRIMARY KEY NOT NULL,
    channel_lock BOOLEAN NOT NULL DEFAULT FALSE,
    channel_lock_id VARCHAR(20),
    temporary_lock BOOLEAN NOT NULL DEFAULT FALSE,
    require_reason BOOLEAN NOT NULL DEFAULT TRUE,
    min_role_id VARCHAR(20),
    log_channel VARCHAR(20)
);