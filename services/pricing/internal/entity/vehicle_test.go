package entity_test

import (
	"testing"

	"github.com/alechekz/online-car-auction-clean-archv2/services/pricing/internal/entity"

	"github.com/stretchr/testify/assert"
)

// test is a struct for build data tests
type test struct {
	name     string
	data     func() *entity.Vehicle
	isValid  bool
	expected uint64
}

// newTestVehicle is a test valid vehicle instance
func newTestVehicle() *entity.Vehicle {
	return &entity.Vehicle{
		VIN:      "1HGBH41JXMN109186",
		Grade:    47,
		Odometer: 30_000,
		Msrp:     99_000,
	}
}

// TestVehicle_Validate tests the Validate method of the Vehicle struct
func TestVehicle_Validate(t *testing.T) {
	tests := []test{
		{
			name: "valid VIN",
			data: func() *entity.Vehicle {
				return newTestVehicle()
			},
			isValid: true,
		},
		{
			name: "zero odometer",
			data: func() *entity.Vehicle {
				v := newTestVehicle()
				v.Odometer = 0
				return v
			},
			isValid: true,
		},
		{
			name: "missing VIN",
			data: func() *entity.Vehicle {
				i := newTestVehicle()
				i.VIN = ""
				return i
			},
			isValid: false,
		},
		{
			name: "invalid VIN",
			data: func() *entity.Vehicle {
				i := newTestVehicle()
				i.VIN = "123"
				return i
			},
			isValid: false,
		},
		{
			name: "missing grade",
			data: func() *entity.Vehicle {
				v := newTestVehicle()
				v.Grade = 0
				return v
			},
			isValid: false,
		},
		{
			name: "too high grade",
			data: func() *entity.Vehicle {
				v := newTestVehicle()
				v.Grade = 51
				return v
			},
			isValid: false,
		},
		{
			name: "too low grade",
			data: func() *entity.Vehicle {
				v := newTestVehicle()
				v.Grade = -1
				return v
			},
			isValid: false,
		},
	}

	// Run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.data().Validate()
			if test.isValid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

// TestVehicle_CalcPrice tests the CalcPrice method of the Vehicle struct
func TestVehicle_CalcPrice(t *testing.T) {
	tests := []test{
		{
			name: "only grade affects price",
			data: func() *entity.Vehicle {
				return newTestVehicle()
			},
			expected: 96_030,
		},
		{
			name: "grade and odometer affect price",
			data: func() *entity.Vehicle {
				data := newTestVehicle()
				data.Odometer = 70_000
				return data
			},
			expected: 91_229,
		},
		{
			name: "grade and exterior color affect price",
			data: func() *entity.Vehicle {
				data := newTestVehicle()
				data.ExteriorColor = "Black"
				return data
			},
			expected: 97_951,
		},
		{
			name: "grade and exterior color affect price",
			data: func() *entity.Vehicle {
				data := newTestVehicle()
				data.InteriorColor = "Red"
				return data
			},
			expected: 92_189,
		},
	}

	// Run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			v := test.data()
			v.CalcPrice()
			assert.Equal(t, test.expected, v.Price)
		})
	}
}
