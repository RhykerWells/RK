package middleware

type ServerCtxKey int

const (
	ContextSessionKey ServerCtxKey = iota
	ContextUserKey
	ContextPortalKey
	ContextMemberKey
	ContextFolderKey
	ContextDocumentKey
)
