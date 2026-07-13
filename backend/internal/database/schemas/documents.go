package schemas

func DocumentSchemas() []Schema {
	return []Schema{
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
}