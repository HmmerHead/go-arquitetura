package cli

import (
	"fmt"

	app "github.com/HmmerHead/go-arquit/application"
)

func Run(service app.ProductServiceInterface,
	action string,
	prodId string,
	productName string,
	price float64) (string, error) {

	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, price)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Create Prod ID %s with the name %s created. Price = %f and status = %s",
			product.GetId(), product.GetName(), product.GetPrice(), product.GetStatus())

	case "enabled":
		product, err := service.Get(prodId)
		if err != nil {
			return result, err
		}
		res, err := service.Enable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Prod %s enable", res.GetName())

	case "disabled":
		product, err := service.Get(prodId)
		if err != nil {
			return result, err
		}
		res, err := service.Disable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Prod %s disable", res.GetName())
	default:
		product, err := service.Get(prodId)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Prod ID %s with the name %s created. Price = %f and status = %s",
			product.GetId(), product.GetName(), product.GetPrice(), product.GetStatus())
	}

	return result, nil
}
