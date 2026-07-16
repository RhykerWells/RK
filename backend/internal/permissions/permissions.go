package permissions

type Permission string

const (
	// Portal permissions
	PermissionPortalManage            Permission = "portal.manage"
	PermissionPortalManageRoles       Permission = "portal.manage.roles"
	PermissionPortalManageMembers     Permission = "portal.manage.members"
	PermissionPortalManageMemberRoles Permission = "portal.manage.member.roles"

	// Folder permissions
	PermissionFolderView   Permission = "folder.view"
	PermissionFolderCreate Permission = "folder.create"
	PermissionFolderRename Permission = "folder.rename"
	PermissionFolderDelete Permission = "folder.delete"

	// Document permissions
	PermissionDocumentView   Permission = "document.view"
	PermissionDocumentCreate Permission = "document.create"
	PermissionDocumentEdit   Permission = "document.edit"
	PermissionDocumentDelete Permission = "document.delete"
	PermissionDocumentShare  Permission = "document.share"

	// File permissions
	PermissionFileUpload Permission = "file.upload"
	PermissionFileDelete Permission = "file.delete"
)

var AllPermissions = []Permission{
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
