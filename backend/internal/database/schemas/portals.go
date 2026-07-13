package schemas

func PortalSchemas() []Schema {
	return []Schema{
		{
			Name: "Portals",
			SQL: `
				CREATE TABLE IF NOT EXISTS portals (
					id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

					name TEXT NOT NULL,
					domain TEXT NOT NULL,

					created_by BIGINT NOT NULL,
					updated_by BIGINT,

					created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
					updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

					CONSTRAINT portals_name_check
						CHECK (length(trim(name)) > 0),

					CONSTRAINT portals_domain_check
						CHECK (length(trim(domain)) > 0),

					CONSTRAINT portals_domain_unique
						UNIQUE (domain),

					CONSTRAINT portals_created_by_fk
						FOREIGN KEY (created_by)
						REFERENCES users(id)
						ON DELETE RESTRICT
						ON UPDATE CASCADE,

					CONSTRAINT portals_updated_by_fk
						FOREIGN KEY (updated_by)
						REFERENCES users(id)
						ON DELETE SET NULL
						ON UPDATE CASCADE
				);
			`,
		},
		{
			Name: "Portal Roles",
			SQL: `
				CREATE TABLE IF NOT EXISTS portal_roles (
					id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

					portal_id BIGINT NOT NULL,

					name TEXT NOT NULL,
					description TEXT,

					-- visual/ordering metadata for roles in the UI
					colour TEXT,
					position BIGINT NOT NULL DEFAULT 0,

					-- optional mapping to a Discord role id (string)
					discord_role_id TEXT,

					created_by BIGINT NOT NULL,
					updated_by BIGINT,

					created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
					updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

					CONSTRAINT portal_roles_portal_fk
						FOREIGN KEY (portal_id)
						REFERENCES portals(id)
						ON DELETE CASCADE
						ON UPDATE CASCADE,

					CONSTRAINT portal_roles_created_by_fk
						FOREIGN KEY (created_by)
						REFERENCES users(id)
						ON DELETE RESTRICT
						ON UPDATE CASCADE,

					CONSTRAINT portal_roles_updated_by_fk
						FOREIGN KEY (updated_by)
						REFERENCES users(id)
						ON DELETE SET NULL
						ON UPDATE CASCADE,

					CONSTRAINT portal_roles_name_check
						CHECK (length(trim(name)) > 0),

					CONSTRAINT portal_roles_name_unique
						UNIQUE (portal_id, name)
				);
			`,
		},
		{
			Name: "Portal Memberships",
			SQL: `
				CREATE TABLE IF NOT EXISTS portal_memberships (
					id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

					portal_id BIGINT NOT NULL,
					user_id BIGINT NOT NULL,

					created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
					updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

					CONSTRAINT portal_memberships_portal_fk
						FOREIGN KEY (portal_id)
						REFERENCES portals(id)
						ON DELETE CASCADE
						ON UPDATE CASCADE,

					CONSTRAINT portal_memberships_user_fk
						FOREIGN KEY (user_id)
						REFERENCES users(id)
						ON DELETE CASCADE
						ON UPDATE CASCADE,

					CONSTRAINT portal_memberships_unique
						UNIQUE (portal_id, user_id)
				);
			`,
		},

		{
			Name: "Portal Membership Roles",
			SQL: `
				CREATE TABLE IF NOT EXISTS portal_membership_roles (
					id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

					portal_membership_id BIGINT NOT NULL,
					portal_role_id BIGINT NOT NULL,

					created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

					CONSTRAINT portal_membership_roles_membership_fk
						FOREIGN KEY (portal_membership_id)
						REFERENCES portal_memberships(id)
						ON DELETE CASCADE
						ON UPDATE CASCADE,

					CONSTRAINT portal_membership_roles_role_fk
						FOREIGN KEY (portal_role_id)
						REFERENCES portal_roles(id)
						ON DELETE CASCADE
						ON UPDATE CASCADE,

					CONSTRAINT portal_membership_roles_unique
						UNIQUE (portal_membership_id, portal_role_id)
				);
			`,
		},
	}
}
