package middleware

type ServerCtxKey int

const (
	ContextPortalKey ServerCtxKey = iota
	ContextSessionKey
	ContextUserKey
	ContextFolderKey
	ContextDocumentKey
)
