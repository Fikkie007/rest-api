package customer

import (
	"context"
	"database/sql"
)

type Repository interface {
	FindAll(ctx context.Context) ([]Customer, error)
	FindById(ctx context.Context, id string) (Customer, error)
	Save(ctx context.Context, c *Customer) error
	Update(ctx context.Context, c *Customer) error
	Delete(ctx context.Context, id string) error
}

type customerRepository struct {
	db *sql.DB
}

func (r *customerRepository) FindAll(
	ctx context.Context,
) ([]Customer, error) {
	// TODO: implement database query
	return nil, nil
}

func (r *customerRepository) FindById(
	ctx context.Context,
	id string,
) (Customer, error) {
	// TODO
	return Customer{}, nil
}

func (r *customerRepository) Save(
	ctx context.Context,
	c *Customer,
) error {
	// TODO
	return nil
}

func (r *customerRepository) Update(
	ctx context.Context,
	c *Customer,
) error {
	// TODO
	return nil
}

func (r *customerRepository) Delete(
	ctx context.Context,
	id string,
) error {
	// TODO
	return nil
}
