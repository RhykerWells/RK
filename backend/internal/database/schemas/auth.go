package schemas

func AuthSchemas() []Schema {
	return []Schema{
		{
			Name: "Global Roles",
			SQL: `
				CREATE TABLE IF NOT EXISTS global_roles (

					id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

					name TEXT NOT NULL,
					description TEXT,

					created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
					updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

					CONSTRAINT global_roles_name_check
						CHECK (length(trim(name)) > 0),

					CONSTRAINT global_roles_name_unique
						UNIQUE (name)
				);
			`,
		},
		{
			Name: "Users",
			SQL: `
				CREATE TABLE IF NOT EXISTS users (

					id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

					global_role_id BIGINT NOT NULL,

					auth_type TEXT NOT NULL,

					username TEXT,
					password_hash TEXT,
					discord_id TEXT,

					display_name TEXT NOT NULL,
					email TEXT,
					avatar_url TEXT,

					is_active BOOLEAN NOT NULL DEFAULT TRUE,

					last_login_at TIMESTAMPTZ,

					created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
					updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

					CONSTRAINT users_global_role_fk
						FOREIGN KEY (global_role_id)
						REFERENCES global_roles(id)
						ON DELETE RESTRICT
						ON UPDATE CASCADE,

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
