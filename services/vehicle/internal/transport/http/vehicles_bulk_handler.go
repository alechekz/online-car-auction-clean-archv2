package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/model"
	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/service"
)

// VehiclesBulkHandler handles HTTP requests for bulk vehicle operations
type VehiclesBulkHandler struct {
	UC service.VehiclesBulkUsecase
}

// CreateBulk - POST /vehicles/bulk - handles the creation of multiple vehicles in bulk
func (h *VehiclesBulkHandler) CreateBulk(c *gin.Context) {
	var vb model.VehiclesBulk
	if err := c.ShouldBindJSON(&vb); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.UC.Create(&vb); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, vb)
}

// UpdateBulk - PUT /vehicles/bulk - handles the update of multiple vehicles in bulk
func (h *VehiclesBulkHandler) UpdateBulk(c *gin.Context) {
	var vb model.VehiclesBulk
	if err := c.ShouldBindJSON(&vb); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.UC.Update(&vb); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, vb)
}
