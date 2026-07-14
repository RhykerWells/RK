package schemas

func AuthSchemas() []Schema {
	return []Schema{
		{
			Name: "Users",
			SQL: `
				CREATE TABLE IF NOT EXISTS users (

					id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

					auth_type TEXT NOT NULL,

					username TEXT NOT NULL,
					password_hash TEXT,
					discord_id TEXT,

					display_name TEXT NOT NULL,
					email TEXT,
					avatar_url TEXT,

					is_active BOOLEAN NOT NULL DEFAULT TRUE,

					last_login_at TIMESTAMPTZ,

					created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
					updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

					is_administrator BOOLEAN NOT NULL DEFAULT FALSE

					CONSTRAINT users_auth_type_check
						CHECK (auth_type IN ('local', 'discord')),

					CONSTRAINT users_local_auth_check
						CHECK (
							auth_type <> 'local'
							OR (
								username IS NOT NULL
								AND password_hash IS NOT NULL
								AND discord_id IS NULL
							)
						),

					CONSTRAINT users_discord_auth_check
						CHECK (
							auth_type <> 'discord'
							OR (
								discord_id IS NOT NULL
								AND username IS NULL
								AND password_hash IS NULL
								)
						),

					CONSTRAINT users_username_check
						CHECK (
							username IS NULL
							OR length(trim(username)) > 0
						),

					CONSTRAINT users_display_name_check
						CHECK (
							length(trim(display_name)) > 0
						)
				);
			`,
		},
	}
}
