package router

import (
	"github.com/Andy4433/Repositorio_go.git/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.PostOpeningHandler )
		v1.DELETE("/opening", handler.DeleteOpeningHandler )
		v1.PUT("/opening",handler.PutOpeningHandler)
		v1.GET("/openings", handler.ListOpeningHandler )
	}
}
