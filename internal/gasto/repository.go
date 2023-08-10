package gasto

import (
	"context"
	"database/sql"
	"errors"

	"github.com/rodrigocarsin/Gastos/internal/domain"
)

// Repository encapsulates the logic to access gasto data.
type Repository interface {
	// Get returns the gasto with the specified gasto ID.
	Get(ctx context.Context, id int64) (domain.Gasto, error)
	// Count returns the number of gastos.
	Count(ctx context.Context) (int64, error)
	// GetAll returns all gastos.
	GetAll(ctx context.Context) ([]domain.Gasto, error)
	// Create saves a new gasto in the storage.
	Create(ctx context.Context, gasto domain.Gasto) (int64, error)
	// Update updates the gasto with given ID in the storage.
	Update(ctx context.Context, gasto domain.Gasto) error
	// Delete removes the gasto with given ID from the storage.
	Delete(ctx context.Context, id int64) error
	// Query returns the list of gastos with the given offset and limit.
	Query(ctx context.Context, offset, limit int64) ([]domain.Gasto, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

// Count implements Repository
func (r *repository) Count(ctx context.Context) (int64, error) {
	query := "SELECT COUNT(*) FROM gasto;"
	row := r.db.QueryRow(query)
	var count int64
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// Create implements Repository
func (r *repository) Create(ctx context.Context, gasto domain.Gasto) (int64, error) {
	query := "INSERT INTO gasto (descripcion, monto, fecha, categoria, tipoPago, comercio) VALUES (?, ?, ?, ?, ?, ?);"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(gasto.Descripcion, gasto.Monto, gasto.Fecha, gasto.Categoria, gasto.TipoPago, gasto.Comercio)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, errors.New("No se pudo obtener el ID")
	}

	gasto.ID = id

	return int64(id), nil
}

// Delete implements Repository
func (r *repository) Delete(ctx context.Context, id int64) error {
	query := "DELETE FROM gasto WHERE id = ?;"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affect < 1 {
		return ErrNotFound
	}

	return nil
}

// Get implements Repository
func (r *repository) Get(ctx context.Context, id int64) (domain.Gasto, error) {
	query := "SELECT id, descripcion, monto, fecha, categoria, tipoPago, comercio FROM gasto WHERE id = ?;"
	row := r.db.QueryRow(query, id)
	gasto := domain.Gasto{}
	err := row.Scan(&gasto.ID, &gasto.Descripcion, &gasto.Monto, &gasto.Fecha, &gasto.Categoria, &gasto.TipoPago, &gasto.Comercio)
	if err != nil {
		return domain.Gasto{}, err
	}

	return gasto, nil
}

// GetAll implements Repository
func (r *repository) GetAll(ctx context.Context) ([]domain.Gasto, error) {
	query := "SELECT id, descripcion, monto, fecha, categoria, tipoPago, comercio FROM gasto"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var gastos []domain.Gasto

	for rows.Next() {
		var gasto domain.Gasto
		if err := rows.Scan(&gasto.ID, &gasto.Descripcion, &gasto.Monto, &gasto.Fecha, &gasto.Categoria, &gasto.TipoPago, &gasto.Comercio); err != nil {
			return nil, err
		}
		gastos = append(gastos, gasto)
	}

	return gastos, nil
}

// Update implements Repository
func (r *repository) Update(ctx context.Context, gasto domain.Gasto) error {
	query := "UPDATE gasto SET descripcion = ?, monto = ?, fecha = ?, categoria = ?, tipoPago = ?, comercio = ? WHERE id = ?;"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(gasto.Descripcion, gasto.Monto, gasto.Fecha, gasto.Categoria, gasto.TipoPago, gasto.Comercio, gasto.ID)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

// Query implements Repository
func (r *repository) Query(ctx context.Context, offset, limit int64) ([]domain.Gasto, error) {
	panic("unimplemented")
}
