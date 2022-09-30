package application_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/ruancaetano/hexagonal-go/application"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10.0

	err := product.Enable()

	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()

	require.Equal(t, err.Error(), "the price must be greater than zero to enable the product")
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 0

	err := product.Disable()

	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()

	require.Equal(t, err.Error(), "the price must be zero to disable the product")
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewString()
	product.Name = "Name"
	product.Status = ""
	product.Price = 10

	_, err := product.IsValid()
	fmt.Println(err)
	require.Nil(t, err)
	require.Equal(t, product.Status, application.DISABLED)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, err.Error(), "status must be DISABLED or ENABLED")

	product.Status = application.DISABLED
	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, err.Error(), "price must be greater or equal 0")

	product.Status = application.DISABLED
	product.Price = 10
	product.ID = "INVALID"
	_, err = product.IsValid()
	require.NotNil(t, err)
}
