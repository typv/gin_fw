package route

import (
	"gin_fw/src/modules/home"
	"github.com/gin-gonic/gin"
)

type RouteFacade struct {
	homeRoute *home.HomeRoute
}

func NewRouteFacade() *RouteFacade {
	return &RouteFacade{
		homeRoute: home.NewHomeRoute(),
	}
}

func (f *RouteFacade) SetupRoutes(r *gin.Engine) {
	f.homeRoute.SetupRoute(r)

	// Setup routes for other modules
}
