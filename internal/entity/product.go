package entity

import (
	"api_mysql/pkg/entity"
	"errors"
	"time"
)

var (
	ErrIDIsRequired    = errors.New("Id is required")
	ErrInvalidID       = errors.New("Id is invalid")
	ErrNameIsRequired  = errors.New("Id is required")
	ErrPriceIsRequired = errors.New("Id is required")
	ErrInvalidPrice    = errors.New("Price is invalid")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		ID:        entity.NewId(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}
	err := product.Validate()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIDIsRequired
	}
	if _, err := entity.Parse(p.ID.String()); err != nil {
		return ErrInvalidID
	}
	if p.Name == "" {
		return ErrNameIsRequired
	}
	if p.Price < 0 {
		return ErrInvalidPrice
	}
	return nil
}
