package cli

import (
	"fmt"

	"github.com/ruancaetano/hexagonal-go/application"
)

func Run(
	service application.ProductServiceInterface,
	action string,
	productId string,
	productName string,
	productPrice float64,
) (string, error) {
	result := ""

	switch action {
	case "create":
		product, err := service.Create(productName, productPrice)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf(
			"Product ID %s with name %s has been created with price %f and status %s\n",
			product.GetID(),
			product.GetName(),
			product.GetPrice(),
			product.GetStatus(),
		)

	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		updatedProduct, err := service.Enable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf(
			"Product ID %s has been enabled\n",
			updatedProduct.GetID(),
		)

	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		updatedProduct, err := service.Disable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf(
			"Product ID %s has been disabled\n",
			updatedProduct.GetID(),
		)

	default:
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf(
			"Produc ID: %s\n Name: %s\n Price: %f\n Status: %s\n",
			product.GetID(),
			product.GetName(),
			product.GetPrice(),
			product.GetStatus(),
		)
	}

	return result, nil
}
