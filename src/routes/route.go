package route

import (
	"gin_fw/src/modules/home"
	"github.com/gin-gonic/gin"
)

type RouteFacade struct {
	homeCtrl *home.HomeController
}

func NewRouteFacade() *RouteFacade {
	return &RouteFacade{
		homeCtrl: home.NewHomeController(),
	}
}

func (f *RouteFacade) SetupRoutes(r *gin.Engine) {
	homeGroup := r.Group("home")
	{
		homeGroup.GET("/", f.homeCtrl.Index)
		homeGroup.GET("/users", f.homeCtrl.GetUsers)
		homeGroup.GET("/departments", f.homeCtrl.GetDepartments)
	}
}
