package http

import (
	"net/http"
	"time"

	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/service"
	"github.com/gin-gonic/gin"
)

// RegisterVehicleRoutes registers vehicle-related routes to the Gin engine
func RegisterVehicleRoutes(r *gin.Engine, uc service.VehicleUsecase, bulkUc service.VehiclesBulkUsecase) {

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"time":    time.Now().Format(time.RFC3339),
			"service": "vehicle-service",
		})
	})

	// Vehicles routers
	h := &VehicleHandler{UC: uc}
	v := r.Group("/vehicles")
	{
		v.POST("", h.Create)
		v.GET("/:vin", h.Get)
		v.PUT("/:vin", h.Update)
		v.DELETE("/:vin", h.Delete)
		v.GET("", h.List)
	}

	// Vehicles bulk routers
	bulkH := &VehiclesBulkHandler{UC: bulkUc}
	bulk := r.Group("/vehicles/bulk")
	{
		bulk.POST("", bulkH.CreateBulk)
		bulk.PUT("", bulkH.UpdateBulk)
	}
}
