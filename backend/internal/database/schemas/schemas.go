package schemas

type Schema struct {
	Name string
	SQL  string
}

func AllSchemas() [][]Schema {
    return [][]Schema{
        AuthSchemas(),
        PortalSchemas(),
        DocumentSchemas(),
        FileSchemas(),
        PermissionSchemas(),
    }
}