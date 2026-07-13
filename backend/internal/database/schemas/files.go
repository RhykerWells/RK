package schemas

func FileSchemas() []Schema {
	return []Schema{
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
	}
}
