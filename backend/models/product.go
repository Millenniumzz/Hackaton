package models

type Product struct {
	Product_id          int    `json:"product_id"`
	Product_name        string `json:"product_name"`
	Product_cost        int    `json:"product_cost"`
	Product_type        string `json:"product_type"`
	Product_description string `json:"product_description"`
	Product_picture     []byte `json:"product_picture"`
}
