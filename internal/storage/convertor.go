package storage

import "smart_school_for_mirea/internal/model"

func convertProductFromDB(product Product) model.Product {
	return model.Product{
		ID:          product.ID,
		Title:       product.Title,
		Description: product.Description,
		Tags:        product.Tags,
		ImageURLs:   product.ImageURL,
	}
}
