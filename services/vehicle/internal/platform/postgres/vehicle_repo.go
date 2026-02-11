package postgres

import (
	"context"
	"time"

	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/entity"
	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// PostgresVehicleRepository is a PostgreSQL implementation of VehicleRepository interface
type PostgresVehicleRepository struct {
	db *pgxpool.Pool
}

// NewVehicleRepo creates a new instance of PostgresVehicleRepository
func NewVehicleRepo(ctx context.Context, conn string) (*PostgresVehicleRepository, error) {
	pool, err := pgxpool.New(ctx, conn)
	if err != nil {
		return nil, err
	}
	return &PostgresVehicleRepository{db: pool}, nil
}

// Save saves a vehicle to the PostgreSQL database
func (r *PostgresVehicleRepository) Save(ctx context.Context, v *entity.Vehicle) error {
	_, err := r.db.Exec(ctx,
		`INSERT INTO vehicles
		(vin, year, odometer, brand, engine, transmission, msrp, grade, price, exterior_color, interior_color, small_scratches, strong_scratches, electric_fail, suspension_fail)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)`,
		v.VIN, v.Year, v.Odometer, v.Brand, v.Engine, v.Transmission, v.MSRP, v.Grade,
		v.Price,
		v.ExteriorColor,
		v.InteriorColor,
		v.SmallScratches,
		v.StrongScratches,
		v.ElectricFail,
		v.SuspensionFail,
	)
	return err
}

// FindByVIN retrieves a vehicle by its VIN
func (r *PostgresVehicleRepository) FindByVIN(ctx context.Context, vin string) (*entity.Vehicle, error) {
	row := r.db.QueryRow(ctx,
		`SELECT vin, year, msrp, odometer, brand, engine, transmission, grade, price, exterior_color, interior_color, small_scratches, strong_scratches, electric_fail, suspension_fail
		FROM vehicles WHERE vin=$1`, vin,
	)
	var v entity.Vehicle
	if err := row.Scan(
		&v.VIN, &v.Year, &v.MSRP, &v.Odometer, &v.Brand, &v.Engine, &v.Transmission, &v.Grade, &v.Price, &v.ExteriorColor, &v.InteriorColor, &v.SmallScratches, &v.StrongScratches, &v.ElectricFail, &v.SuspensionFail,
	); err != nil {
		return nil, err
	}
	return &v, nil
}

// Update updates an existing vehicle in the PostgreSQL database
func (r *PostgresVehicleRepository) Update(ctx context.Context, v *entity.Vehicle) error {
	_, err := r.db.Exec(ctx,
		`UPDATE vehicles
		SET year=$1, odometer=$2, brand=$3, engine=$4, transmission=$5, msrp=$6, grade=$7, price=$8, exterior_color=$9, interior_color=$10, small_scratches=$11, strong_scratches=$12, electric_fail=$13, suspension_fail=$14
		WHERE vin=$15`,
		v.Year, v.Odometer, v.Brand, v.Engine, v.Transmission, v.MSRP, v.Grade, v.Price, v.ExteriorColor, v.InteriorColor, v.SmallScratches, v.StrongScratches, v.ElectricFail, v.SuspensionFail, v.VIN,
	)
	return err
}

// Delete removes a vehicle from the PostgreSQL database by its VIN
func (r *PostgresVehicleRepository) Delete(ctx context.Context, vin string) error {
	_, err := r.db.Exec(ctx,
		`DELETE FROM vehicles WHERE vin=$1`, vin,
	)
	return err
}

// List lists all vehicles in the PostgreSQL database
func (r *PostgresVehicleRepository) List(ctx context.Context) ([]*entity.Vehicle, error) {
	rows, err := r.db.Query(ctx,
		`SELECT
		vin, year, msrp, odometer, brand, engine, transmission, grade, price, exterior_color, interior_color, small_scratches, strong_scratches, electric_fail, suspension_fail
		FROM vehicles`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vehicles []*entity.Vehicle
	for rows.Next() {
		var v entity.Vehicle
		if err := rows.Scan(
			&v.VIN, &v.Year, &v.MSRP, &v.Odometer, &v.Brand, &v.Engine, &v.Transmission, &v.Grade, &v.Price, &v.ExteriorColor, &v.InteriorColor, &v.SmallScratches, &v.StrongScratches, &v.ElectricFail, &v.SuspensionFail,
		); err != nil {
			return nil, err
		}
		vehicles = append(vehicles, &v)
	}
	return vehicles, nil
}

// copy copies a batch of vehicles to the PostgreSQL database
func (r *PostgresVehicleRepository) copy(ctx context.Context, tx pgx.Tx, table string, vehicles []*entity.Vehicle) error {

	// Prepare data for bulk insert
	values := make([][]any, len(vehicles))
	for i, v := range vehicles {
		values[i] = []any{v.VIN, v.Year, v.Odometer, v.Brand, v.Engine, v.Transmission, v.MSRP, v.Grade, v.Price, v.ExteriorColor, v.InteriorColor, v.SmallScratches, v.StrongScratches, v.ElectricFail, v.SuspensionFail}
	}

	// Perform bulk insert using CopyFrom
	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{table},
		[]string{"vin", "year", "odometer", "brand", "engine", "transmission", "msrp", "grade", "price", "exterior_color", "interior_color", "small_scratches", "strong_scratches", "electric_fail", "suspension_fail"},
		pgx.CopyFromRows(values),
	)
	return err

}

// SaveBulk saves a vehicles bulk to the PostgreSQL database
func (r *PostgresVehicleRepository) SaveBulk(vb *model.VehiclesBulk) error {

	// Begin a transaction
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx) // nolint:errcheck

	// Insert vehicles in batches
	size := 200
	for i := 0; i < len(vb.Vehicles); i += size {
		end := min(i+size, len(vb.Vehicles))
		if err := r.copy(ctx, tx, "vehicles", vb.Vehicles[i:end]); err != nil {
			return err
		}
	}

	// Commit the transaction
	return tx.Commit(ctx)
}

// UpdateBulk updates a vehicles bulk in the PostgreSQL database
func (r *PostgresVehicleRepository) UpdateBulk(vb *model.VehiclesBulk) error {

	// Begin a transaction
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx) // nolint:errcheck

	// Create a temporary table to hold the updated data
	_, err = tx.Exec(ctx, `
        CREATE TEMP TABLE tmp_vehicles (
			vin VARCHAR(17) PRIMARY KEY,
  			year INT NOT NULL,
  			odometer INT NOT NULL,
  			msrp NUMERIC,
  			price NUMERIC,
  			grade NUMERIC,
			brand VARCHAR(50),
			engine VARCHAR(50),
			transmission VARCHAR(50),
			exterior_color VARCHAR(50),
			interior_color VARCHAR(50),
			small_scratches BOOLEAN,
			strong_scratches BOOLEAN,
			electric_fail BOOLEAN,
			suspension_fail BOOLEAN
        ) ON COMMIT DROP;
    `)
	if err != nil {
		return err
	}

	// Copy the updated data into the temporary table
	size := 200
	for i := 0; i < len(vb.Vehicles); i += size {
		end := min(i+size, len(vb.Vehicles))
		if err := r.copy(ctx, tx, "tmp_vehicles", vb.Vehicles[i:end]); err != nil {
			return err
		}
	}

	// Update the main vehicles table using the data from the temporary table
	_, err = tx.Exec(ctx, `
        UPDATE vehicles v
        SET price = t.price,
            year  = t.year,
			odometer = t.odometer,
			msrp = t.msrp,
			grade = t.grade,
			brand = t.brand,
			engine = t.engine,
			transmission = t.transmission,
			exterior_color = t.exterior_color,
			interior_color = t.interior_color,
			small_scratches = t.small_scratches,
			strong_scratches = t.strong_scratches,
			electric_fail = t.electric_fail,
			suspension_fail = t.suspension_fail
        FROM tmp_vehicles t
        WHERE v.vin = t.vin;
    `)
	if err != nil {
		return err
	}

	// Commit the transaction
	return tx.Commit(ctx)
}
