package casbin

import "testing"

func TestModelRBAC_CheckPermission(t *testing.T) {

	t.Run("validate user permission", func(t *testing.T) {

		policies := Policies{
			{
				Subject: "role-bar-id", // role-bar
				Object:  "dept-list",
				Action:  "get",
			},
		}

		model := NewRBAC(policies)

		uid := "user-foo-1"
		model.AddRoleForUser(uid, "role-bar-id")

		ok, err := model.CheckPermission(uid, "dept-list", "get")
		if err != nil || !ok {
			t.Fail()
		}
	})
}
