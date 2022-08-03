package models

import (
	"gorm.io/gorm"
)

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

func (order *Order) CheckOrderValueChanged(orderRequest OrderRequest) (err error) {
	if orderRequest.FarmerId != "" && order.FarmerId != orderRequest.FarmerId {
		order.FarmerId = orderRequest.FarmerId
	}

	if orderRequest.CustomerId != "" && order.CustomerId != orderRequest.CustomerId {
		order.CustomerId = orderRequest.CustomerId
	}

	if orderRequest.Status != "" && order.Status != orderRequest.Status {
		order.Status = orderRequest.Status
	}

	if orderRequest.Chanel != "" && order.Chanel != orderRequest.Chanel {
		order.Chanel = orderRequest.Chanel
	}
	return nil
}

// func (orderList *OrderList) CheckOrderListValueChanged(orderListRequest OrderList) (err error) {

// 	// orderListRequest.OrderId ต้องไม่สามารถแก้ไขได้
// 	if orderListRequest.ProductId != "" && orderList.ProductId != orderListRequest.ProductId {
// 		orderList.ProductId = orderListRequest.ProductId
// 	}
// 	if orderListRequest.ProductName != "" && orderList.ProductName != orderListRequest.ProductName {
// 		orderList.ProductName = orderListRequest.ProductName
// 	}
// 	if orderListRequest.Price != "" && orderList.Price != orderListRequest.Price {
// 		orderList.Price = orderListRequest.Price
// 	}
// 	if orderListRequest.Qty != "" && orderList.Qty != orderListRequest.Qty {
// 		orderList.Qty = orderListRequest.Qty
// 	}
// 	return nil
// }
