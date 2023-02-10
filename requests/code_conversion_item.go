package requests

type CodeConversionItem struct {
	PurchaseOrderItem string  `json:"PurchaseOrderItem"`
	OrderItem         int     `json:"OrderItem"`
	OrderItemCategory *string `json:"OrderItemCategory"`
	Material          string  `json:"Material"`
	Product           *string `json:"Product"`
	ProductGroup      *string `json:"ProductGroup"`
}
