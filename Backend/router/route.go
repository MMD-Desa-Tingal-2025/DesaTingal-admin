package router

import (
	"github.com/gin-gonic/gin"
	"github.com/fachru/backend/controllers"
	// Import all your handler packages here
	// Example:
	// "your_project/handler/studio"
)

func StudioRouter(r *gin.Engine) {
	studio := r.Group("/studio")
	{
		// Tambahkan route sesuai handler yang sudah Anda miliki
		// Contoh:
		// studio.GET("/", studio.GetAllStudios)
		// studio.POST("/", studio.CreateStudio)
		// studio.GET("/:id", studio.GetStudioByID)
		// studio.PUT("/:id", studio.UpdateStudio)
		// studio.DELETE("/:id", studio.DeleteStudio)
	}
}