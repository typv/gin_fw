package home

import "github.com/gin-gonic/gin"

type HomeController struct {
	homeService *HomeService
}

func NewHomeController() *HomeController {
	return &HomeController{
		homeService: NewHomeService(),
	}
}

func (ctrl *HomeController) Index(c *gin.Context) {
	msg := ctrl.homeService.GetHello()
	res := gin.H{
		"message": msg,
	}
	c.JSON(200, res)
}

func (ctrl *HomeController) GetUsers(c *gin.Context) {
	msg := ctrl.homeService.GetUsers()
	res := gin.H{
		"message": msg,
	}
	c.JSON(200, res)
}

func (ctrl *HomeController) GetDepartments(c *gin.Context) {
	msg := ctrl.homeService.GetDepartments()
	res := gin.H{
		"message": msg,
	}
	c.JSON(200, res)
}
