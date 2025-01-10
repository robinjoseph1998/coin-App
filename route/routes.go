package route

import (
	"coin-App/src/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, ctrl *controllers.Controller) {

	router.GET("/ping", controllers.HealthPing)

	router.POST("/addcoin", ctrl.AddCoin)
	router.POST("/updatecoin", ctrl.UpdateCoin)
	router.GET("/view/all", ctrl.ListAll)
	router.GET("/viewbyname/or/id", ctrl.GetByNameOrId)
	router.GET("/view/expiredcoins/log", ctrl.ViewExpiredLogs)
}
