package http

import (
	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/entity"
	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// VehicleHandler handles HTTP requests for vehicle operations
type VehicleHandler struct {
	UC service.VehicleUsecase
}

// Create - POST /vehicles - handles the creation of a new vehicle
func (h *VehicleHandler) Create(c *gin.Context) {
	var v entity.Vehicle
	if err := c.ShouldBindJSON(&v); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := c.Request.Context()
	if err := h.UC.Create(ctx, &v); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, v)
}

// Update - PUT /vehicles/:vin - handles the update of an existing vehicle
func (h *VehicleHandler) Update(c *gin.Context) {
	vin := c.Param("vin")
	var v entity.Vehicle
	if err := c.ShouldBindJSON(&v); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := c.Request.Context()
	v.VIN = vin
	if err := h.UC.Update(ctx, &v); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, v)
}

// Get - GET /vehicles/:vin - handles fetching a vehicle by VIN
func (h *VehicleHandler) Get(c *gin.Context) {
	ctx := c.Request.Context()
	vin := c.Param("vin")
	v, err := h.UC.Get(ctx, vin)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, v)
}

// Delete - DELETE /vehicles/:vin - handles deleting a vehicle by VIN
func (h *VehicleHandler) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	vin := c.Param("vin")
	if err := h.UC.Delete(ctx, vin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// List - GET /vehicles - handles listing all vehicles
func (h *VehicleHandler) List(c *gin.Context) {
	ctx := c.Request.Context()
	vehicles, err := h.UC.List(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, vehicles)
}
