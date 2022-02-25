package casbin

import "testing"

func TestModelRBAC_CheckPermission(t *testing.T) {

	t.Run("valid user permission", func(t *testing.T) {

		policies := Policies{
			{
				Subject: "role-bar-id", // 角色-bar
				Object:  "dept-list",
				Action:  "Get",
			},
		}

		model := NewRBAC(policies)

		uid := "user-foo-1"
		model.AddRoleForUser(uid, "role-bar-id")

		ok, err := model.CheckPermission(uid, "dept-list", "Get")
		if err != nil || !ok {
			t.Fail()
		}
	})
}
