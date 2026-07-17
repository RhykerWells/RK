package errors

import "errors"

var (
	ErrInvalidUserID             = errors.New("invalid user id")
	ErrUserNotFound              = errors.New("user not found")
	ErrInvalidMissingRequestType = errors.New("invalid or missing request type")
)

var (
	ErrInvalidPortalID = errors.New("invalid portal id")
	ErrPortalNotFound  = errors.New("portal not found")
)

var (
	ErrInvalidRoleID = errors.New("invalid role ID")
	ErrRoleNotFound  = errors.New("role not found")
)

var (
	ErrMemberNotFound = errors.New("member not found")
)

var (
	ErrInvalidFolderID = errors.New("invalid folder ID")
	ErrFolderNotFound  = errors.New("folder not found")
)

var (
	ErrInvalidDocumentID = errors.New("invalid document ID")
	ErrDocumentNotFound  = errors.New("document not found")

	ErrNewDocumentInvalidLocation  = errors.New("document requires either a folder or owner")
	ErrUpdateDocumentInvalidFolder = errors.New("document folder must be valid")
)
