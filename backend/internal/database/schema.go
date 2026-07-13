package database

type Schema struct {
	Name string
	SQL  string
}

var Schemas = []Schema{
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
				portal_role_id BIGINT NOT NULL,

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

				CONSTRAINT portal_memberships_role_fk
					FOREIGN KEY (portal_role_id)
					REFERENCES portal_roles(id)
					ON DELETE RESTRICT
					ON UPDATE CASCADE,

				CONSTRAINT portal_memberships_unique
					UNIQUE (portal_id, user_id)
			);
		`,
	},
	{
		Name: "Folders",
		SQL: `
			CREATE TABLE IF NOT EXISTS folders (
				id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

				portal_id BIGINT NOT NULL,
				parent_folder_id BIGINT,

				name TEXT NOT NULL,

				created_by BIGINT NOT NULL,
				updated_by BIGINT,

				created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

				CONSTRAINT folders_portal_fk
					FOREIGN KEY (portal_id)
					REFERENCES portals(id)
					ON DELETE CASCADE
					ON UPDATE CASCADE,

				CONSTRAINT folders_parent_fk
					FOREIGN KEY (parent_folder_id)
					REFERENCES folders(id)
					ON DELETE CASCADE
					ON UPDATE CASCADE,

				CONSTRAINT folders_created_by_fk
					FOREIGN KEY (created_by)
					REFERENCES users(id)
					ON DELETE RESTRICT
					ON UPDATE CASCADE,

				CONSTRAINT folders_updated_by_fk
					FOREIGN KEY (updated_by)
					REFERENCES users(id)
					ON DELETE SET NULL
					ON UPDATE CASCADE,

				CONSTRAINT folders_name_check
					CHECK (length(trim(name)) > 0),

				CONSTRAINT folders_root_parent_check
					CHECK (
						parent_folder_id IS NOT NULL
						OR name = '/'
					)
			);
		`,
	},
	{
		Name: "Documents",
		SQL: `
			CREATE TABLE IF NOT EXISTS documents (
				id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

				folder_id BIGINT NOT NULL,

				owner_id BIGINT,
				created_by BIGINT NOT NULL,
				updated_by BIGINT,

				title TEXT NOT NULL,

				latest_version_id BIGINT,

				created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

				CONSTRAINT documents_folder_fk
					FOREIGN KEY (folder_id)
					REFERENCES folders(id)
					ON DELETE CASCADE
					ON UPDATE CASCADE,

				CONSTRAINT documents_owner_fk
					FOREIGN KEY (owner_id)
					REFERENCES users(id)
					ON DELETE SET NULL
					ON UPDATE CASCADE,

				CONSTRAINT documents_created_by_fk
					FOREIGN KEY (created_by)
					REFERENCES users(id)
					ON DELETE RESTRICT
					ON UPDATE CASCADE,

				CONSTRAINT documents_updated_by_fk
					FOREIGN KEY (updated_by)
					REFERENCES users(id)
					ON DELETE SET NULL
					ON UPDATE CASCADE,

				CONSTRAINT documents_title_check
					CHECK (length(trim(title)) > 0)
			);
		`,
	},
	{
		Name: "Document Versions",
		SQL: `
			CREATE TABLE IF NOT EXISTS document_versions (
				id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

				document_id BIGINT NOT NULL,

				version_number BIGINT NOT NULL,

				content JSONB NOT NULL,

				created_by BIGINT NOT NULL,

				created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

				CONSTRAINT document_versions_document_fk
					FOREIGN KEY (document_id)
					REFERENCES documents(id)
					ON DELETE CASCADE
					ON UPDATE CASCADE,

				CONSTRAINT document_versions_created_by_fk
					FOREIGN KEY (created_by)
					REFERENCES users(id)
					ON DELETE RESTRICT
					ON UPDATE CASCADE,

				CONSTRAINT document_versions_version_unique
					UNIQUE (document_id, version_number),

				CONSTRAINT document_versions_version_check
					CHECK (version_number > 0)
			);
		`,
	},
	{
		Name: "Files",
		SQL: `
			CREATE TABLE IF NOT EXISTS files (
				id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

				storage_key TEXT NOT NULL,

				original_filename TEXT NOT NULL,
				mime_type TEXT NOT NULL,
				file_size BIGINT NOT NULL,

				checksum TEXT,

				uploaded_by BIGINT NOT NULL,

				created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

				CONSTRAINT files_uploaded_by_fk
					FOREIGN KEY (uploaded_by)
					REFERENCES users(id)
					ON DELETE RESTRICT
					ON UPDATE CASCADE,

				CONSTRAINT files_storage_key_unique
					UNIQUE (storage_key),

				CONSTRAINT files_original_filename_check
					CHECK (length(trim(original_filename)) > 0),

				CONSTRAINT files_mime_type_check
					CHECK (length(trim(mime_type)) > 0),

				CONSTRAINT files_file_size_check
					CHECK (file_size >= 0)
			);
		`,
	},
	{
		Name: "Document Files",
		SQL: `
			CREATE TABLE IF NOT EXISTS document_files (
				id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

				document_id BIGINT NOT NULL,
				file_id BIGINT NOT NULL,

				attachment_type TEXT NOT NULL DEFAULT 'attachment',

				created_by BIGINT NOT NULL,

				created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

				CONSTRAINT document_files_document_fk
					FOREIGN KEY (document_id)
					REFERENCES documents(id)
					ON DELETE CASCADE
					ON UPDATE CASCADE,

				CONSTRAINT document_files_file_fk
					FOREIGN KEY (file_id)
					REFERENCES files(id)
					ON DELETE RESTRICT
					ON UPDATE CASCADE,

				CONSTRAINT document_files_created_by_fk
					FOREIGN KEY (created_by)
					REFERENCES users(id)
					ON DELETE RESTRICT
					ON UPDATE CASCADE,

				CONSTRAINT document_files_type_check
					CHECK (
						attachment_type IN ('attachment', 'embedded')
					),

				CONSTRAINT document_files_unique
					UNIQUE (document_id, file_id)
			);
		`,
	},
	{
		Name: "Permission Definitions",
		SQL: `
			CREATE TABLE IF NOT EXISTS permission_definitions (
				id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

				name TEXT NOT NULL,
				description TEXT,

				created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

				CONSTRAINT permission_definitions_name_unique
					UNIQUE (name),

				CONSTRAINT permission_definitions_name_check
					CHECK (length(trim(name)) > 0)
			);
		`,
	},
	{
		Name: "Folder Permissions",
		SQL: `
			CREATE TABLE IF NOT EXISTS folder_permissions (
				id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

				folder_id BIGINT NOT NULL,

				portal_role_id BIGINT,
				user_id BIGINT,

				permission_id BIGINT NOT NULL,

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

				CONSTRAINT folder_permissions_permission_fk
					FOREIGN KEY (permission_id)
					REFERENCES permission_definitions(id)
					ON DELETE CASCADE
					ON UPDATE CASCADE,

				CONSTRAINT folder_permissions_created_by_fk
					FOREIGN KEY (created_by)
					REFERENCES users(id)
					ON DELETE RESTRICT
					ON UPDATE CASCADE,

				CONSTRAINT folder_permissions_subject_check
					CHECK (
						portal_role_id IS NOT NULL
						OR user_id IS NOT NULL
					)
			);
		`,
	},
	{
		Name: "Document Links",
		SQL: `
			CREATE TABLE IF NOT EXISTS document_links (
				id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

				source_document_id BIGINT NOT NULL,
				target_document_id BIGINT NOT NULL,

				created_by BIGINT NOT NULL,

				created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

				CONSTRAINT document_links_source_fk
					FOREIGN KEY (source_document_id)
					REFERENCES documents(id)
					ON DELETE CASCADE
					ON UPDATE CASCADE,

				CONSTRAINT document_links_target_fk
					FOREIGN KEY (target_document_id)
					REFERENCES documents(id)
					ON DELETE RESTRICT
					ON UPDATE CASCADE,

				CONSTRAINT document_links_created_by_fk
					FOREIGN KEY (created_by)
					REFERENCES users(id)
					ON DELETE RESTRICT
					ON UPDATE CASCADE,

				CONSTRAINT document_links_not_self_check
					CHECK (
						source_document_id <> target_document_id
					),

				CONSTRAINT document_links_unique
					UNIQUE (source_document_id, target_document_id)
			);
		`,
	},
}