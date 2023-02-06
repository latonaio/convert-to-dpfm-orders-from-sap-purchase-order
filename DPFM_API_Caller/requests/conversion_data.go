package requests

type ConversionData struct {
	PurchaseOrder     string  `json:"PurchaseOrder"`
	OrderID           int     `json:"OrderID"`
	PurchaseOrderItem string  `json:"PurchaseOrderItem"`
	OrderItem         int     `json:"OrderItem"`
	Seller            int     `json:"Seller"`
	Supplier          *string `json:"Supplier"`
	Material          string  `json:"Material"`
	Product           *string `json:"Product"`
}
