package http_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/entity"
	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/platform/logger"
	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/testhelpers"
)

// TestMain sets up the testing environment
func TestMain(m *testing.M) {
	logger.Init()
	os.Exit(m.Run())
}

// TestVehicleHandler_Create tests the Create Vehicle HTTP handler
func TestVehicleHandler_Create(t *testing.T) {
	router := testhelpers.NewTestVehicleRouter()
	v := testhelpers.NewTestVehicle()

	// Valid case
	t.Run("valid request", func(t *testing.T) {
		body, _ := json.Marshal(v)

		req := httptest.NewRequest(http.MethodPost, "/vehicles", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusCreated, rec.Code)

		var v entity.Vehicle
		err := json.Unmarshal(rec.Body.Bytes(), &v)
		assert.NoError(t, err)
		assert.Equal(t, "1HGBH41JXMN109186", v.VIN)
	})

	// Invalid case
	t.Run("invalid request body", func(t *testing.T) {
		v.VIN = "123"
		body, _ := json.Marshal(v)
		req := httptest.NewRequest(http.MethodPost, "/vehicles", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}

// TestVehicleHandler_Update tests the Update Vehicle HTTP handler
func TestVehicleHandler_Update(t *testing.T) {

	// Prepare router with a vehicle
	router := testhelpers.NewTestVehicleRouter()
	v := testhelpers.NewTestVehicle()

	// Valid case
	t.Run("update existing vehicle", func(t *testing.T) {
		body, _ := json.Marshal(v)
		req := httptest.NewRequest(http.MethodPut, "/vehicles/"+v.VIN, bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusOK, rec.Code)

		var got entity.Vehicle
		err := json.Unmarshal(rec.Body.Bytes(), &got)
		assert.NoError(t, err)
		assert.Equal(t, "1HGBH41JXMN109186", got.VIN)
	})

	// Invalid case
	t.Run("update with invalid VIN", func(t *testing.T) {
		v.VIN = "123"
		body, _ := json.Marshal(v)
		req := httptest.NewRequest(http.MethodPut, "/vehicles/"+v.VIN, bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}

// TestVehicleHandler_Get tests the Get Vehicle HTTP handler
func TestVehicleHandler_Get(t *testing.T) {

	// Prepare router with a vehicle
	router := testhelpers.NewTestVehicleRouter()
	v := testhelpers.NewTestVehicle()

	// Valid case
	t.Run("existing vehicle", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/vehicles/"+v.VIN, nil)
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		var got entity.Vehicle
		err := json.Unmarshal(rec.Body.Bytes(), &got)
		assert.NoError(t, err)
		assert.Equal(t, v.VIN, got.VIN)
	})
}

// TestVehicleHandler_Delete tests the Delete Vehicle HTTP handler
func TestVehicleHandler_Delete(t *testing.T) {

	// Prepare router with a vehicle
	router := testhelpers.NewTestVehicleRouter()
	v := testhelpers.NewTestVehicle()

	// Valid case
	t.Run("delete existing vehicle", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/vehicles/"+v.VIN, nil)
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNoContent, rec.Code)
	})
}

// TestVehicleHandler_List tests the List Vehicles HTTP handler
func TestVehicleHandler_List(t *testing.T) {

	// Prepare router with vehicles
	router := testhelpers.NewTestVehicleRouter()

	// Not empty list case
	t.Run("list all vehicles", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/vehicles", nil)
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		var got []entity.Vehicle
		err := json.Unmarshal(rec.Body.Bytes(), &got)
		assert.NoError(t, err)
		assert.Len(t, got, 2)
	})
}
