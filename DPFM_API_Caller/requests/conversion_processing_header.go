package requests

type ConversionProcessingHeader struct {
	ConvertingOrderID   *string `json:"ConvertingOrderID"`
	ConvertedOrderID    *int    `json:"ConvertedOrderID"`
	ConvertingOrderType *string `json:"ConvertingOrderType"`
	ConvertedOrderType  *string `json:"ConvertedOrderType"`
	ConvertingSeller    *string `json:"ConvertingSeller"`
	ConvertedSeller     *int    `json:"ConvertedSeller"`
}
