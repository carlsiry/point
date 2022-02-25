package casbin

import "github.com/casbin/casbin/v2"

type ModelRBAC struct {
	*casbin.Enforcer
}

func (m *ModelRBAC) literal() string {
	const model = `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
	`

	return model
}

func (m *ModelRBAC) set(enforcer *casbin.Enforcer) {
	m.Enforcer = enforcer
}

func (m *ModelRBAC) load(lst Policies) {
	lst.Batch(m.AddPermissionsForUser)
}

func (m *ModelRBAC) DeleteRoles(lst []string) {
	for _, id := range lst {
		m.DeletePermissionsForUser(id)
		m.DeleteRole(id)
	}
}

func (m *ModelRBAC) SetPermissionByRole(id string) {
	// enforcer.DeletePermissionsForUser(PrefixRoleID + id)
}

func (m *ModelRBAC) CheckPermission(userID, url, method string) (bool, error) {
	return m.Enforce(userID, url, method)
}

func (m *ModelRBAC) AddRolesForUser(id string, lstRoleID []string) {
	for _, role := range lstRoleID {
		m.AddRoleForUser(id, role)
	}
}

func NewRBAC(policies Policies) *ModelRBAC {
	model := new(ModelRBAC)
	initEnforcer(model, policies)

	return model
}
