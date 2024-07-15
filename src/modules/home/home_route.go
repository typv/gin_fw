package home

import (
	"github.com/gin-gonic/gin"
)

type HomeRoute struct {
	homeCtrl *HomeController
}

func NewHomeRoute() *HomeRoute {
	return &HomeRoute{
		homeCtrl: NewHomeController(),
	}
}

func (f *HomeRoute) SetupRoute(r *gin.Engine) {
	group := r.Group("home")
	{
		group.GET("/", f.homeCtrl.Index)
		group.GET("/test", f.homeCtrl.Test)
	}
}
