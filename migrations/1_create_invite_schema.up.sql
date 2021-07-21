CREATE TABLE invites (
    id SERIAL PRIMARY KEY,
    code TEXT NOT NULL,
    guild_id TEXT NOT NULL,
    created_at TIME NOT NULL,
    user_created TEXT NOT NULL,
    target_channel TEXT NOT NULL,
    max_uses INT NOT NULL
);