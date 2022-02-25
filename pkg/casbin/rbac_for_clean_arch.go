package casbin

type IAccess interface {
	EnableUser(ID string, forRoleListID []string)
	DisableUser(ID string)
}

type Access = ModelRBAC

func (a *Access) EnableUser(ID string, forRoleListID []string) {
	a.AddRolesForUser(ID, forRoleListID)
}

func (a *Access) DisableUser(ID string) {
	a.DeleteUser(ID)
}

func NewAccess(policies Policies) (res *Access) {
	res = NewRBAC(policies)

	return
}
