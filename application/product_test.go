package app_test

import (
	"testing"

	app "github.com/HmmerHead/go-arquit/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := app.Product{}
	product.Name = "Hello"
	product.Status = app.DISABLED
	product.Price = 10

	t.Run("Habilitar Produto", func(t *testing.T) {
		err := product.Enable()
		require.Nil(t, err)
	})

	t.Run("Habilitar produto com valor menor/igual que zero", func(t *testing.T) {
		product.Price = 0
		err := product.Enable()
		require.Equal(t, "o preco tem que ser maior que zero para ativar o prod", err.Error())
	})
}

func TestProduct_Disabled(t *testing.T) {
	product := app.Product{}
	product.Name = "Hello"
	product.Status = app.ENABLED
	product.Price = 0

	t.Run("Desabilitar Produto", func(t *testing.T) {
		err := product.Disable()
		require.Nil(t, err)
	})

	t.Run("Desabilitar produto com valor menor/igual que zero", func(t *testing.T) {
		product.Price = 10
		err := product.Disable()
		require.Equal(t, "o preco tem que ser igual a zero para desativar o prod", err.Error())
	})
}

func TestProduct_IsValid(t *testing.T) {
	product := app.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Hello"
	product.Status = app.DISABLED
	product.Price = 10

	t.Run("Produto Valido", func(t *testing.T) {
		_, err := product.IsValid()
		require.Nil(t, err)
	})

	t.Run("Produto sem status", func(t *testing.T) {
		product.Status = "INVALID"
		_, err := product.IsValid()
		require.Equal(t, "precisa de um status", err.Error())
	})

	t.Run("Produto habilitado", func(t *testing.T) {
		product.Status = app.ENABLED
		_, err := product.IsValid()
		require.Nil(t, err)
	})

	t.Run("Produto sem preco", func(t *testing.T) {
		product.Price = -10
		_, err := product.IsValid()
		require.Equal(t, "precisa de um preco", err.Error())
	})
}
