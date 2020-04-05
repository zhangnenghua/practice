package routers

import (
	"beegodemo/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"net/http"
)

func init() {
	// 基础路由
	beego.Get("/basicrouter/get", func(context *context.Context) {
		context.WriteString("success")
	})
	beego.Post("/basicrouter/post", func(context *context.Context) {
		context.WriteString("success")
	})
	beego.Any("/basicrouter/any", func(context *context.Context) {
		context.WriteString("success")
	})
	beego.Handler("/basicrouter/handler", &BasicRouterHandler{})

	// 固定路由
	beego.Router("/", &controllers.MainController{})
	beego.Router("/custom/config", &controllers.CustomConfigController{})
	beego.Router("/app/config", &controllers.AppConfigController{})

	// 正则路由
	beego.Router("/regexprouter/?:id", &controllers.RegexpRouterController{})
	//beego.Router("/regexprouter/:id", &controllers.RegexpRouterController{})
	//beego.Router("/regexprouter/:id([0-9]+)", &controllers.RegexpRouterController{})
	//beego.Router("/regexprouter/:id(\\w+)", &controllers.RegexpRouterController{})

	// 自定义方法及 RESTful 规则
	beego.Router("/custommethodrouter", &controllers.CustomMethodRouterController{}, "get,delete:CustomMethod;post:CustomPostMethod")

	// 自动匹配
	beego.AutoRouter(&controllers.AutoRouterController{})
}

type BasicRouterHandler struct{}

func (p *BasicRouterHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("success"))
}
