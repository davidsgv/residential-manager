package entities

const (
	CreateAdmin          = 1 << iota // 1 << 0 = 1
	CreateWatchman                   // 1 << 1 = 2
	CreateApartmentAdmin             // 1 << 2 = 4
	CreateResident                   // 1 << 3 = 8

	QueryUsers
	QueryPermissions
	QueryRoles

	UpdateUsers

	DeleteUser
)

const (
	NotAllowedQueryUsers           = "You dont have enought permissions to query users"
	NotAllowedCreateAdmin          = ""
	NotAllowedCreateWatchman       = ""
	NotAllowedCreateApartmentAdmin = ""
	NotAllowedCreateResident       = ""
	NotAllowedQueryPermissions     = ""
	NotAllowedQueryRoles           = ""
	NotAllowedUpdateUsers          = ""
	NotAllowedDeleteUser           = ""
)

// return all the permissions in the app
func GetAllPermissions() map[string]int {
	permissions := map[string]int{
		"Create Admin":          int(CreateAdmin),
		"Create Watchman":       int(CreateWatchman),
		"Create ApartmentAdmin": int(CreateApartmentAdmin),
		"Create Resident":       int(CreateResident),
		"Query Users":           int(QueryUsers),
		"Query Permissions":     int(QueryPermissions),
		"Query Roles":           int(QueryRoles),
		"Update User":           int(QueryRoles),
		"Delete User":           int(DeleteUser),
	}
	return permissions
}
