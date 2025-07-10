package router

import (
	"github.com/gin-gonic/gin"
	// Import all your handler packages here
	// Example:
	// "your_project/handler/studio"
)

// StudioRouter sets up the studio routes.
func StudioRouter(r *gin.Engine) {
	studio := r.Group("/studio")
	_ = studio // Remove this line if you add actual routes below

	{
		// Tambahkan route sesuai handler yang sudah Anda miliki
		// Contoh:
		// studio.GET("/", controllers.GetAllStudios)
		// studio.POST("/", controllers.CreateStudio)
		// studio.GET("/:id", controllers.GetStudioByID)
		// studio.PUT("/:id", controllers.UpdateStudio)
		// studio.DELETE("/:id", controllers.DeleteStudio)
	}
}