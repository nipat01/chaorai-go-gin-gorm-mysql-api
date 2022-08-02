package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	FarmerId   string `json:"farmerId"`
	CustomerId string `json:"customerId"`
	Status     string `json:"status"`
	Chanel     string `json:"chanel"` //store, subscription

}

type OrderList struct {
	gorm.Model
	OrderId     uint   `json:"orderId"`
	ProductId   string `json:"productId"`
	ProductName string `json:"productName"`
	Price       string `json:"price"`
	Qty         string `json:"qty"`
}

type OrderRequest struct {
	gorm.Model
	FarmerId   string      `json:"farmerId"`
	CustomerId string      `json:"customerId"`
	Status     string      `json:"status"`
	Chanel     string      `json:"chanel"` //store, subscription
	OrderList  []OrderList `json:"orderList"`
}

type OrderResponse struct {
	gorm.Model
	FarmerId   string      `json:"farmerId"`
	CustomerId string      `json:"customerId"`
	Status     string      `json:"status"`
	Chanel     string      `json:"chanel"` //store, subscription
	OrderList  []OrderList `json:"orderList"`
}
