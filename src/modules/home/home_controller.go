package home

import "github.com/gin-gonic/gin"

type HomeController struct {
	homeService *HomeService
}

func NewHomeController() *HomeController {
	return &HomeController{}
}

func (ctrl *HomeController) Index(c *gin.Context) {
	msg := ctrl.homeService.GetHello()
	res := gin.H{
		"message": msg,
	}
	c.JSON(200, res)
}

func (ctrl *HomeController) Test(c *gin.Context) {
	msg := ctrl.homeService.Test()
	res := gin.H{
		"message": msg,
	}
	c.JSON(200, res)
}
