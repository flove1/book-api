CREATE TABLE IF NOT EXISTS suspensions (
    id bigserial PRIMARY KEY,
    reason text NOT NULL,
    user_id bigint NOT NULL REFERENCES users ON DELETE CASCADE,
    moderator_id bigint NOT NULL REFERENCES users ON DELETE CASCADE,
    expires_in interval NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);