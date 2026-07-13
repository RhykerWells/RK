package schemas

func PermissionSchemas() []Schema {
	return []Schema{
		{
			Name: "Portal Permissions",
			SQL: `
				CREATE TABLE IF NOT EXISTS portal_role_permissions (

					id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

					portal_role_id BIGINT NOT NULL,
					permission_key TEXT NOT NULL,

					allow BOOLEAN NOT NULL,

					CONSTRAINT portal_role_permissions_role_fk
						FOREIGN KEY (portal_role_id)
						REFERENCES portal_roles(id)
						ON DELETE CASCADE
						ON UPDATE CASCADE,

					CONSTRAINT portal_role_permissions_unique
						UNIQUE (
							portal_role_id,
							permission_key
						)
				);
			`,
		},
		{
			Name: "Folder Permissions",
			SQL: `
				CREATE TABLE IF NOT EXISTS folder_permission_overrides (
					id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

					folder_id BIGINT NOT NULL,

					portal_role_id BIGINT,
					user_id BIGINT,

					permission_key TEXT NOT NULL,

					allow BOOLEAN NOT NULL,

					created_by BIGINT NOT NULL,

					created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

					CONSTRAINT folder_permissions_folder_fk
						FOREIGN KEY (folder_id)
						REFERENCES folders(id)
						ON DELETE CASCADE
						ON UPDATE CASCADE,

					CONSTRAINT folder_permissions_role_fk
						FOREIGN KEY (portal_role_id)
						REFERENCES portal_roles(id)
						ON DELETE CASCADE
						ON UPDATE CASCADE,

					CONSTRAINT folder_permissions_user_fk
						FOREIGN KEY (user_id)
						REFERENCES users(id)
						ON DELETE CASCADE
						ON UPDATE CASCADE,

					CONSTRAINT folder_permissions_created_by_fk
						FOREIGN KEY (created_by)
						REFERENCES users(id)
						ON DELETE RESTRICT
						ON UPDATE CASCADE,

					CONSTRAINT folder_permissions_subject_check
						CHECK (
							(portal_role_id IS NOT NULL) <>
							(user_id IS NOT NULL)
						)
				);
			`,
		},
		{
			Name: "Document Permissions",
			SQL: `
				CREATE TABLE IF NOT EXISTS document_permission_overrides (

					id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

					document_id BIGINT NOT NULL,

					portal_role_id BIGINT,
					user_id BIGINT,

					permission_key TEXT NOT NULL,

					allow BOOLEAN NOT NULL,

					created_by BIGINT NOT NULL,

					created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

					CONSTRAINT document_permission_overrides_document_fk
						FOREIGN KEY (document_id)
						REFERENCES documents(id)
						ON DELETE CASCADE
						ON UPDATE CASCADE,

					CONSTRAINT document_permission_overrides_role_fk
						FOREIGN KEY (portal_role_id)
						REFERENCES portal_roles(id)
						ON DELETE CASCADE
						ON UPDATE CASCADE,

					CONSTRAINT document_permission_overrides_user_fk
						FOREIGN KEY (user_id)
						REFERENCES users(id)
						ON DELETE CASCADE
						ON UPDATE CASCADE,

					CONSTRAINT document_permission_overrides_created_by_fk
						FOREIGN KEY (created_by)
						REFERENCES users(id)
						ON DELETE RESTRICT
						ON UPDATE CASCADE,

					CONSTRAINT document_permission_overrides_subject_check
						CHECK (
							(portal_role_id IS NOT NULL) <>
							(user_id IS NOT NULL)
						),

					CONSTRAINT document_permission_overrides_unique
						UNIQUE (
							document_id,
							portal_role_id,
							user_id,
							permission_key
						)
				);
			`,
		},
	}
}
