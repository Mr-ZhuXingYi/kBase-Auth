package common

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/gorm-adapter/v3"
	"github.com/shenyisyn/goft-gin/goft"
)

var E *casbin.Enforcer

type Casbinconfig struct{}

func NewCasbinconfig() *Casbinconfig {
	return &Casbinconfig{}
}

func (this *Casbinconfig) InitCasbin() *casbin.Enforcer {
	adapter, err := gormadapter.NewAdapterByDB(DB)

	e, err := casbin.NewEnforcer("resources/model.conf", adapter)
	goft.Error(err)
	goft.Error(e.LoadPolicy())
	E = e
	return e
}
