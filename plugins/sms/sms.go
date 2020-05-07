package sms

import (
	"github.com/GoAdminGroup/go-admin/context"
	c "github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/service"
	"github.com/GoAdminGroup/go-admin/plugins"
)

type SMS struct {
	Base *plugins.Base
}

func NewSms() *SMS {
	return &SMS{
		Base: &plugins.Base{PlugName: "sms", URLPrefix: "sms"},
	}
}

func (e *SMS) InitPlugin(srv service.List) {
	e.InitBase(srv)
	e.App = e.initRouter(c.Prefix(), srv)
}

func (e *SMS) initRouter(prefix, srv service.List) *context.App {

	app := context.NewApp()
	route := app.Group(prefix)
	/*// 加入认证中间件
	route.GET("/show/me/something", auth.Middleware(db.GetConnection(srv)), func(ctx *context.Context){
		// 控制器逻辑
	})*/
	route.POST("/setconfig", auth.Middleware(db.GetConnection(srv)), e.setConf)
	route.GET("/configlist", auth.Middleware(db.GetConnection(srv)), e.confList)
	route.GET("/config/:id", auth.Middleware(db.GetConnection(srv)), e.getConf)
	return app
}
