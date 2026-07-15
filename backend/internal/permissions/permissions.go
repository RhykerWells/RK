package permissions

const (
	// Portal permissions
	PermissionPortalManage        = "portal.manage"
	PermissionPortalManageMembers = "portal.manage.members"
	PermissionPortalManageRoles   = "portal.manage.roles"

	// Folder permissions
	PermissionFolderView   = "folder.view"
	PermissionFolderCreate = "folder.create"
	PermissionFolderRename = "folder.rename"
	PermissionFolderDelete = "folder.delete"

	// Document permissions
	PermissionDocumentView   = "document.view"
	PermissionDocumentCreate = "document.create"
	PermissionDocumentEdit   = "document.edit"
	PermissionDocumentDelete = "document.delete"
	PermissionDocumentShare  = "document.share"

	// File permissions
	PermissionFileUpload = "file.upload"
	PermissionFileDelete = "file.delete"
)

var AllPermissions = []string{
	PermissionPortalManage,
	PermissionPortalManageMembers,
	PermissionPortalManageRoles,

	PermissionFolderView,
	PermissionFolderCreate,
	PermissionFolderRename,
	PermissionFolderDelete,

	PermissionDocumentView,
	PermissionDocumentCreate,
	PermissionDocumentEdit,
	PermissionDocumentDelete,
	PermissionDocumentShare,

	PermissionFileUpload,
	PermissionFileDelete,
}
