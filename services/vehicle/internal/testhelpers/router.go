package testhelpers

import (
	"github.com/gin-gonic/gin"

	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/transport/http"
)

// NewTestVehicleRouter sets up a Gin router with vehicle routes for testing
func NewTestVehicleRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	h := &http.VehicleHandler{UC: NewTestVehicleUC()}

	v := r.Group("/vehicles")
	{
		v.POST("", h.Create)
		v.GET("/:vin", h.Get)
		v.PUT("/:vin", h.Update)
		v.DELETE("/:vin", h.Delete)
		v.GET("", h.List)
	}
	return r
}
