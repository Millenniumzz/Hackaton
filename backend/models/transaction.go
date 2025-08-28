package models

type Transaction struct {
	Trans_id   int    `json:"trans_id"`
	User_id    string `json:"user_id"`
	Product_id int    `json:"product_id"`
	Date       string    `json:"date"`
}
