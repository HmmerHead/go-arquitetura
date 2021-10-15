package app

import (
	"errors"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetId() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

type ProductServiceInterface interface {
	Get(id string) (ProductInterface, error)
	Create(name string, price float64) (ProductInterface, error)
	Enable(product ProductInterface) (ProductInterface, error)
	Disable(product ProductInterface) (ProductInterface, error)
}

type ProductReader interface {
	Get(id string) (ProductInterface, error)
}

type ProductWriter interface {
	Save(product ProductInterface) (ProductInterface, error)
}

type ProductPersistenceInterface interface {
	ProductReader
	ProductWriter
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

type Product struct {
	ID     string  `valid:"uuidv4"`
	Name   string  `valid:"required"`
	Status string  `valid:"required"`
	Price  float64 `valid:"float,optional"`
}

func NewProduct() *Product {
	product := Product{
		ID:     uuid.NewV4().String(),
		Status: DISABLED,
	}
	return &product
}

func (prod *Product) IsValid() (bool, error) {
	if prod.Status == "" {
		prod.Status = DISABLED
	}

	if prod.Status != ENABLED && prod.Status != DISABLED {
		return false, errors.New("precisa de um status")
	}

	if prod.Price < 0 {
		return false, errors.New("precisa de um preco")
	}

	_, err := govalidator.ValidateStruct(prod)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (prod *Product) Enable() error {
	if prod.Price > 0 {
		prod.Status = ENABLED
		return nil
	}

	return errors.New("o preco tem que ser maior que zero para ativar o prod")
}

func (prod *Product) Disable() error {
	if prod.Price == 0 {
		prod.Status = DISABLED
		return nil
	}
	return errors.New("o preco tem que ser igual a zero para desativar o prod")
}

func (prod *Product) GetId() string {
	return prod.ID
}

func (prod *Product) GetName() string {
	return prod.Name
}

func (prod *Product) GetStatus() string {
	return prod.Status
}

func (prod *Product) GetPrice() float64 {
	return prod.Price
}
