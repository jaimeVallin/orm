package routes

import (
	"inventario/go-backend/controllers"

	"github.com/gin-gonic/gin"
)

func PersonRouter(router *gin.Engine) {

	routes := router.Group("api/v1/persons")
	routes.POST("", controllers.PersonCreate)
	routes.GET("/:id", controllers.PersonGetById)
	routes.GET("", controllers.PersonGet)
	routes.PUT("/:id", controllers.PersonUpdate)
	routes.DELETE("/:id", controllers.PersonDelete)

}
