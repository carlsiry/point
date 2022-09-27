package casbin

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
)

type loader func(subject string, permissions ...[]string) (bool, error)

type Policy struct {
	Subject string
	Object  string
	Action  string
}

type Policies []Policy

func (lst Policies) Batch(load loader) {
	for _, it := range lst {
		load(it.Subject, []string{it.Object, it.Action})
	}
}

type Model interface {
	literal() string
	set(*casbin.Enforcer)
	load(policies Policies)
}

func initEnforcer(from Model, policies Policies) (*casbin.Enforcer, error) {
	m, err := model.NewModelFromString(from.literal())
	if err != nil {
		return nil, err
	}

	enforcer, err := casbin.NewEnforcer(m)
	if err != nil {
		return nil, err
	}

	from.set(enforcer)
	from.load(policies)

	return enforcer, nil
}
