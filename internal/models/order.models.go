package models

import "time"

type OrderItem struct {
	ProductID string `json:"product_id"`
	Qty       int    `json:"qty"`
	SizeID    int    `json:"size_id"`
	IsIced    bool   `json:"is_iced"`
}

type OrderHistory struct {
	TransactionCode string    `json:"transaction_code"`
	Date            time.Time `json:"date"`
	GrandTotal      int       `json:"grand_total"`
	OrderId         int       `json:"order_id"`
	Status          string    `json:"status"`
	Path            *string   `json:"path"`
}

type CreateOrderRequest struct {
	Email            string      `json:"email"`
	Fullname         string      `json:"fullname"`
	Address          string      `json:"address"`
	DeliveryMethodID int         `json:"delivery_method_id"`
	PaymentMethodID  int         `json:"payment_method_id"`
	Items            []OrderItem `json:"items"`
}

type OrderItemResponse struct {
	ProductID   string `json:"product_id"`
	ProductName string `json:"product_name"`
	Qty         int    `json:"qty"`
	Size        string `json:"size"`
	IsIced      bool   `json:"is_iced"`
}

type CreateOrderResponse struct {
	Email          string              `json:"email"`
	Fullname       string              `json:"fullname"`
	Address        string              `json:"address"`
	DeliveryMethod int                 `json:"delivery_method"`
	PaymentMethod  int                 `json:"payment_method"`
	Items          []OrderItemResponse `json:"items"`
	DeliveryFee    int                 `json:"delivery_fee"`
	Total          int                 `json:"total"`
	Tax            int                 `json:"tax"`
	TotalAmount    int                 `json:"total_amount"`
}

type Item struct {
	ProductID   string `json:"product_id"`
	ProductName string `json:"product_name"`
	Qty         int    `json:"qty"`
	Size        string `json:"size"`
	IceOrHot    string `json:"ice_or_hot"`
	Price       string `json:"price"`
}

type OrderDetailsResponse struct {
	TransactionCode string    `json:"transaction_code"`
	Fullname        string    `json:"fullname"`
	Phone           string    `json:"phone"`
	OrderDate       time.Time `json:"order_date"`
	PaymentMethod   int       `json:"payment_method"`
	DeliveryMethod  int       `json:"delivery_method"`
	Status          string    `json:"status"`
	TotalAmount     int       `json:"total_amount"`
	Items           []Item    `json:"items"`
}
