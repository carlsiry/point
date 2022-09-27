package casbin

import "testing"

func TestModelRBAC_CheckPermission(t *testing.T) {

	// t.Run("validate user permission with api", func(t *testing.T) {

	// 	policies := Policies{
	// 		{
	// 			Subject: "role-bar-id", // role-bar
	// 			Object:  "dept-list",
	// 			Action:  "get",
	// 		},
	// 	}

	// 	model := NewRBAC(policies)

	// 	uid := "user-foo-1"
	// 	model.AddRoleForUser(uid, "role-bar-id")

	// 	ok, err := model.CheckPermission(uid, "dept-list", "get")
	// 	if err != nil || !ok {
	// 		t.Fail()
	// 	}
	// })

	t.Run("validate user permission with etc", func(t *testing.T) {

		lst := [][]string{
			{"role-bar-id", "dept-list", "get", "allow"},
			{"role-bar-id", "dept-list", "create", "allow"},
			{"role-bar-id", "dept-list", "delete", "allow"},
			{"role-foo-id", "dept-list", "delete", "deny"},
		}

		model := NewRBAC(nil)
		model.loadPolicy(lst)
		model.AddGroupingPolicy("role-foo-id", "role-bar-id")

		uid := "user-foo-1"
		model.AddRoleForUser(uid, "role-foo-id")

		ok, err := model.CheckPermission(uid, "dept-list", "get")
		if err != nil || !ok {
			t.Fail()
		}

		ok, err = model.CheckPermission(uid, "dept-list", "delete")
		if err != nil || ok {
			t.Fail()
		}
	})
}
