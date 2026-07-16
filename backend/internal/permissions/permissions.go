package permissions

type Permission string

const (
	PermissionAdministrator Permission = "administrator"

	// Global Portal permissions
	PermissionPortalsManage Permission = "portals.manage"

	// Individual Portal permissions
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

type Definition struct {
	HumanReadableName string
	Description       string
}

var Definitions = map[Permission]Definition{
	PermissionAdministrator: {
		HumanReadableName: "Administrator",
		Description:       "Grants all permissions and access to all portals",
	},
	// Global Portal permissions
	PermissionPortalsManage: {
		HumanReadableName: "Manage Portals",
		Description:       "Allows management of site portals",
	},
	// Individual Portal permissions
	PermissionPortalManageRoles: {
		HumanReadableName: "Manage Portal Roles",
		Description:       "Allows management of portal roles",
	},
	PermissionPortalManageMembers: {
		HumanReadableName: "Manage Portal Members",
		Description:       "Allows management of portal members",
	},
	PermissionPortalManageMemberRoles: {
		HumanReadableName: "Manage Member Roles",
		Description:       "Allows management of member roles",
	},

	// Folder permissions
	PermissionFolderView: {
		HumanReadableName: "View Folder",
		Description:       "Allows viewing of folder contents",
	},
	PermissionFolderCreate: {
		HumanReadableName: "Create Folder",
		Description:       "Allows creation of new folders",
	},
	PermissionFolderRename: {
		HumanReadableName: "Rename Folder",
		Description:       "Allows renaming of folders",
	},
	PermissionFolderDelete: {
		HumanReadableName: "Delete Folder",
		Description:       "Allows deletion of folders",
	},

	// Document permissions
	PermissionDocumentView: {
		HumanReadableName: "View Document",
		Description:       "Allows viewing of documents",
	},
	PermissionDocumentCreate: {
		HumanReadableName: "Create Document",
		Description:       "Allows creation of new documents",
	},
	PermissionDocumentEdit: {
		HumanReadableName: "Edit Document",
		Description:       "Allows editing of documents",
	},
	PermissionDocumentDelete: {
		HumanReadableName: "Delete Document",
		Description:       "Allows deletion of documents",
	},
	PermissionDocumentShare: {
		HumanReadableName: "Share Document",
		Description:       "Allows sharing of documents",
	},

	// File permissions
	PermissionFileUpload: {
		HumanReadableName: "Upload File",
		Description:       "Allows uploading of files",
	},
	PermissionFileDelete: {
		HumanReadableName: "Delete File",
		Description:       "Allows deletion of files",
	},
}

var AllPermissions = []Permission{
	PermissionPortalsManage,
	PermissionPortalManageRoles,
	PermissionPortalManageMembers,
	PermissionPortalManageMemberRoles,

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

func PermissionName(permission Permission) string {
	if def, ok := Definitions[permission]; ok {
		return def.HumanReadableName
	}
	return string(permission)
}
