package infrastructure

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	entadapter "github.com/casbin/ent-adapter"
)

func Initmodel() model.Model {
	// Initialize the model from Go code.
	m := model.NewModel()
	m.AddDef("r", "r", "sub, obj, act")
	m.AddDef("p", "p", "sub, obj, act")
	m.AddDef("g", "g", "_, _")
	m.AddDef("e", "e", "some(where (p.eft == allow))")
	m.AddDef("m", "m", "g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act")

	return m
}

func CasbinLoad(driver string, conndsn string) *casbin.Enforcer {
	a, _ := entadapter.NewAdapter(driver, conndsn)
	m := Initmodel()

	e, _ := casbin.NewEnforcer(m, a)

	e.LoadPolicy()
	return e
}
