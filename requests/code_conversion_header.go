package requests

type CodeConversionHeader struct {
	PurchaseOrder string  `json:"PurchaseOrder"`
	OrderID       int     `json:"OrderID"`
	OrderType     *string `json:"OrderType"`
	Supplier      *string `json:"Supplier"`
	Seller        *int    `json:"Seller"`
}
