package gasto

import (
	"context"
	"errors"

	"github.com/rodrigocarsin/Gastos/internal/domain"
)

// Errors
var (
	ErrNotFound             = errors.New("product not found")
	ErrDuplicateProductCode = errors.New("duplicate product code")
	ErrDoesNotExist         = errors.New("product does not exist")
	ErrForeignKey           = errors.New("entity corresponding to foreign key does not exist")
)

// Service encapsulates usecase logic for products.
type Service interface {
	Get(ctx context.Context, id int64) (domain.Gasto, error)
	Count(ctx context.Context) (int64, error)
	Query(ctx context.Context, offset, limit int64) ([]domain.Gasto, error)
	Create(ctx context.Context, gasto domain.Gasto) (domain.Gasto, error)
	Update(ctx context.Context, descripcion *string, monto *float64, fecha *string, categoria *string, comercio *string, tipoPago *string, id int64) (domain.Gasto, error)
	Delete(ctx context.Context, id int64) error
	GetAll(ctx context.Context) ([]domain.Gasto, error)
}

// Repository is the interface that provides product methods.
type service struct {
	repo Repository
}

// NewService creates a new product service.
func NewService(repo Repository) Service {
	return service{
		repo: repo,
	}
}

// Get returns the product with the specified the product ID.
func (s service) Get(ctx context.Context, id int64) (domain.Gasto, error) {

	gasto, err := s.repo.Get(ctx, id)
	if err != nil {
		return domain.Gasto{}, ErrNotFound
	}
	return gasto, nil
}

// Count returns the number of products.
func (s service) Count(ctx context.Context) (int64, error) {

	count, err := s.repo.Count(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// Query returns the products with the specified offset and limit.
func (s service) Query(ctx context.Context, offset, limit int64) ([]domain.Gasto, error) {

	gastos, err := s.repo.Query(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	return gastos, nil
}

// GetAll returns all products.
func (s service) GetAll(ctx context.Context) ([]domain.Gasto, error) {

	gastos, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return gastos, nil
}

// Create creates a new product.
func (s service) Create(ctx context.Context, gasto domain.Gasto) (domain.Gasto, error) {

	id, err := s.repo.Create(ctx, gasto)
	if err != nil {
		return domain.Gasto{}, err
	}
	createdGasto, _ := s.repo.Get(ctx, id)
	return createdGasto, nil
}

// Update updates the product with the specified ID.
func (s service) Update(ctx context.Context, descripcion *string, monto *float64, fecha *string, categoria *string, comercio *string, tipoPago *string, id int64) (domain.Gasto, error) {

	g, err := s.repo.Get(ctx, id)
	if err != nil {
		return domain.Gasto{}, ErrNotFound
	}

	gasto := domain.Gasto{ID: id}

	if descripcion != nil {
		gasto.Descripcion = *descripcion
	} else {
		gasto.Descripcion = g.Descripcion
	}

	if monto != nil {
		gasto.Monto = *monto
	} else {
		gasto.Monto = g.Monto
	}

	if fecha != nil {
		gasto.Fecha = *fecha
	} else {
		gasto.Fecha = g.Fecha
	}

	if categoria != nil {
		gasto.Categoria = *categoria
	} else {
		gasto.Categoria = g.Categoria
	}

	if comercio != nil {
		gasto.Comercio = *comercio
	} else {
		gasto.Comercio = g.Comercio
	}

	if tipoPago != nil {
		gasto.TipoPago = *tipoPago
	} else {
		gasto.TipoPago = g.TipoPago
	}

	err = s.repo.Update(ctx, gasto)
	if err != nil {
		return domain.Gasto{}, err
	}

	updatedGasto, _ := s.repo.Get(ctx, id)

	return updatedGasto, nil

}

// Delete deletes the product with the specified ID.
func (s service) Delete(ctx context.Context, id int64) error {

	_, err := s.repo.Get(ctx, id)
	if err != nil {
		return ErrNotFound
	}

	err = s.repo.Delete(ctx, id)

	return err
}
