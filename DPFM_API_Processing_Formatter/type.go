package dpfm_api_processing_formatter

type ProcessingFormatterSDC struct {
	Header                     *Header                     `json:"Header"`
	ConversionProcessingHeader *ConversionProcessingHeader `json:"ConversionProcessingHeader"`
}

type ConversionProcessingKey struct {
	SystemConvertTo       string   `json:"SystemConvertTo"`
	SystemConvertFrom     string   `json:"SystemConvertFrom"`
	LabelConvertTo        string   `json:"LabelConvertTo"`
	LabelConvertFrom      string   `json:"LabelConvertFrom"`
	CodeConvertFromInt    *int     `json:"CodeConvertFromInt"`
	CodeConvertFromFloat  *float32 `json:"CodeConvertFromFloat"`
	CodeConvertFromString *string  `json:"CodeConvertFromString"`
	BusinessPartner       int      `json:"BusinessPartner"`
}

type ConversionProcessingCommonQueryGets struct {
	CodeConversionID      int      `json:"CodeConversionID"`
	SystemConvertTo       string   `json:"SystemConvertTo"`
	SystemConvertFrom     string   `json:"SystemConvertFrom"`
	LabelConvertTo        string   `json:"LabelConvertTo"`
	LabelConvertFrom      string   `json:"LabelConvertFrom"`
	CodeConvertFromInt    *int     `json:"CodeConvertFromInt"`
	CodeConvertFromFloat  *float32 `json:"CodeConvertFromFloat"`
	CodeConvertFromString *string  `json:"CodeConvertFromString"`
	CodeConvertToInt      *int     `json:"CodeConvertToInt"`
	CodeConvertToFloat    *float32 `json:"CodeConvertToFloat"`
	CodeConvertToString   *string  `json:"CodeConvertToString"`
	BusinessPartner       int      `json:"BusinessPartner"`
}

type Header struct {
	ConvertingOrderID         string   `json:"ConvertingOrderID"`
	OrderDate                 *string  `json:"OrderDate"`
	ConvertingOrderType       *string  `json:"ConvertingOrderType"`
	Buyer                     *int     `json:"Buyer"`
	ConvertingSeller          *string  `json:"ConvertingSeller"`
	BillToParty               *int     `json:"BillToParty"`
	BillToCountry             *string  `json:"BillToCountry"`
	Payer                     *int     `json:"Payer"`
	CreationDate              *string  `json:"CreationDate"`
	LastChangeDate            *string  `json:"LastChangeDate"`
	OrderValidityStartDate    *string  `json:"OrderValidityStartDate"`
	OrderValidityEndDate      *string  `json:"OrderValidityEndDate"`
	TransactionCurrency       *string  `json:"TransactionCurrency"`
	PricingDate               *string  `json:"PricingDate"`
	PriceDetnExchangeRate     *float32 `json:"PriceDetnExchangeRate"`
	Incoterms                 *string  `json:"Incoterms"`
	PaymentTerms              *string  `json:"PaymentTerms"`
	AccountingExchangeRate    *float32 `json:"AccountingExchangeRate"`
	HeaderBlockStatus         *bool    `json:"HeaderBlockStatus"`
	HeaderBillingBlockStatus  *bool    `json:"HeaderBillingBlockStatus"`
	HeaderDeliveryBlockStatus *bool    `json:"HeaderDeliveryBlockStatus"`
	IsCancelled               *bool    `json:"IsCancelled"`
	IsDeleted                 *bool    `json:"IsDeleted"`
}

type ConversionProcessingHeader struct {
	ConvertingOrderID   *string `json:"ConvertingOrderID"`
	ConvertedOrderID    *int    `json:"ConvertedOrderID"`
	ConvertingOrderType *string `json:"ConvertingOrderType"`
	ConvertedOrderType  *string `json:"ConvertedOrderType"`
	ConvertingSeller    *string `json:"ConvertingSeller"`
	ConvertedSeller     *int    `json:"ConvertedSeller"`
}

// type Item struct {
// ConvertingOrderID                       string   `json:"ConvertingOrderID"`
// ConvertingOrderItem                     string   `json:"ConvertingOrderItem"`
// ConvertingOrderItemCategory             *string  `json:"ConvertingOrderItemCategory"`
// OrderItemTextBySeller                   *string  `json:"OrderItemTextBySeller"`
// ConvertingProduct                       *string  `json:"ConvertingProduct"`
// ConvertingProductGroup                  *string  `json:"ConvertingProductGroup"`
// BaseUnit                                *string  `json:"BaseUnit"`
// PricingDate                             *string  `json:"PricingDate"`
// PriceDetnExchangeRate                   *float32 `json:"PriceDetnExchangeRate"`
// RequestedDeliveryDate                   *string  `json:"RequestedDeliveryDate"`
// ConvertingDeliverToParty                *string  `json:"ConvertingDeliverToParty"`
// DeliverFromParty                        *int     `json:"DeliverFromParty"`
// CreationDate                            *string  `json:"CreationDate"`
// LastChangeDate                          *string  `json:"LastChangeDate"`
// DeliverFromPlant                        *string  `json:"DeliverFromPlant"`
// DeliverFromPlantStorageLocation         *string  `json:"DeliverFromPlantStorageLocation"`
// DeliveryUnit                            *string  `json:"DeliveryUnit"`
// StockConfirmationBusinessPartner        *int     `json:"StockConfirmationBusinessPartner"`
// StockConfirmationPlant                  *string  `json:"StockConfirmationPlant"`
// StockConfirmationPlantBatch             *string  `json:"StockConfirmationPlantBatch"`
// OrderQuantityInDeliveryUnit             *float32 `json:"OrderQuantityInDeliveryUnit"`
// ConfirmedOrderQuantityInBaseUnit        *float32 `json:"ConfirmedOrderQuantityInBaseUnit"`
// ItemWeightUnit                          *string  `json:"ItemWeightUnit"`
// ItemGrossWeight                         *float32 `json:"ItemGrossWeight"`
// ItemNetWeight                           *float32 `json:"ItemNetWeight"`
// NetAmount                               *float32 `json:"NetAmount"`
// TaxAmount                               *float32 `json:"TaxAmount"`
// GrossAmount                             *float32 `json:"GrossAmount"`
// InvoiceDocumentDate                     *string  `json:"InvoiceDocumentDate"`
// ProductionPlantBusinessPartner          *int     `json:"ProductionPlantBusinessPartner"`
// ProductionPlant                         *string  `json:"ProductionPlant"`
// Incoterms                               *string  `json:"Incoterms"`
// TransactionTaxClassification            *string  `json:"TransactionTaxClassification"`
// ProductTaxClassificationBillToCountry   *string  `json:"ProductTaxClassificationBillToCountyr"`
// ProductTaxClassificationBillFromCountry *string  `json:"ProductTaxClassificationBillFromCountry"`
// AccountAssignmentGroup                  *string  `json:"AccountAssignmentGroup"`
// ProductAccountAssignmentGroup           *string  `json:"ProductAccountAssignmentGroup"`
// PaymentTerms                            *string  `json:"PaymentTerms"`
// TaxCode                                 *string  `json:"TaxCode"`
// ItemBlockStatus                         *bool    `json:"ItemBlockStatus"`
// ItemBillingBlockStatus                  *bool    `json:"ItemBillingBlockStatus"`
// ItemDeliveryBlockStatus                 *bool    `json:"ItemDeliveryBlockStatus"`
// IsCancelled                             *bool    `json:"IsCancelled"`
// IsDeleted                               *bool    `json:"IsDeleted"`
// }
//
// type ConversionProcessingItem struct {
// ConvertingOrderItem         *string `json:"ConvertingOrderItem"`
// ConvertedOrderItem          *int    `json:"ConvertedOrderItem"`
// ConvertingOrderItemCategory *string `json:"ConvertingOrderItemCategory"`
// ConvertedOrderItemCategory  *string `json:"ConvertedOrderItemCategory"`
// ConvertingProduct           *string `json:"ConvertingProduct"`
// ConvertedProduct            *string `json:"ConvertedProduct"`
// ConvertingProductGroup      *string `json:"ConvertingProductGroup"`
// ConvertedProductGroup       *string `json:"ConvertedProductGroup"`
// ConvertingDeliverToParty    *string `json:"ConvertingDeliverToParty"`
// ConvertedDeliverToParty     *int    `json:"ConvertedDeliverToParty"`
// }
//
// type ItemPricingElement struct {
// ConvertingOrderID                   string   `json:"ConvertingOrderID"`
// ConvertingOrderItem                 string   `json:"ConvertingOrderItem"`
// ConvertingBuyer                     *string  `json:"ConvertingBuyer"`
// Seller                              *int     `json:"Seller"`
// ConvertingConditionRecord           *string  `json:"ConvertingConditionRecord"`
// ConvertingConditionSequentialNumber *string  `json:"ConvertingConditionSequentialNumber"`
// PricingDate                         *string  `json:"PricingDate"`
// ConditionRateValue                  *float32 `json:"ConditionRateValue"`
// ConditionCurrency                   *string  `json:"ConditionCurrency"`
// ConditionQuantity                   *float32 `json:"ConditionQuantity"`
// ConditionQuantityUnit               *string  `json:"ConditionQuantityUnit"`
// TaxCode                             *string  `json:"TaxCode"`
// ConditionAmount                     *float32 `json:"ConditionAmount"`
// TransactionCurrency                 *string  `json:"TransactionCurrency"`
// ConditionIsManuallyChanged          *bool    `json:"ConditionIsManuallyChanged"`
// }
//
// type ConversionProcessingItemPricingElement struct {
// ConvertingConditionRecord           *string `json:"ConvertingConditionRecord"`
// ConvertedConditionRecord            *int    `json:"ConvertedConditionRecord"`
// ConvertingConditionSequentialNumber *string `json:"ConvertingConditionSequentialNumber"`
// ConvertedConditionSequentialNumber  *int    `json:"ConvertedConditionSequentialNumber"`
// }
//
// type ItemScheduleLine struct {
// ConvertingOrderID                     string   `json:"ConvertingOrderID"`
// ConvertingOrderItem                   string   `json:"ConvertingOrderItem"`
// ConvertingScheduleLine                string   `json:"ConvertingScheduleLine"`
// Product                               *string  `json:"Product"`
// StockConfirmationBussinessPartner     *int     `json:"StockConfirmationBussinessPartner"`
// StockConfirmationPlant                *string  `json:"StockConfirmationPlant"`
// StockConfirmationPlantBatch           *string  `json:"StockConfirmationPlantBatch"`
// RequestedDeliveryDate                 *string  `json:"RequestedDeliveryDate"`
// ConfirmedDeliveryDate                 *string  `json:"ConfirmedDeliveryDate"`
// OrderQuantityInBaseUnit               *float32 `json:"OrderQuantityInBaseUnit"`
// ConfirmedOrderQuantityByPDTAvailCheck *float32 `json:"ConfirmedOrderQuantityByPDTAvailCheck"`
// DeliveredQuantityInBaseUnit           *float32 `json:"DeliveredQuantityInBaseUnit"`
// OpenConfirmedQuantityInBaseUnit       *float32 `json:"OpenConfirmedQuantityInBaseUnit"`
// }
//
// type ConversionProcessingItemScheduleLine struct {
// ConvertingScheduleLine *string `json:"ConvertingScheduleLine"`
// ConvertedScheduleLine  *int    `json:"ConvertedScheduleLine"`
// }
//
// type Address struct {
// ConvertingOrderID string `json:"ConvertingOrderID"`
// }
//
// type Partner struct {
// ConvertingOrderID         string  `json:"ConvertingOrderID"`
// ConvertingPartnerFunction string  `json:"ConvertingPartnerFunction"`
// ConvertingBusinessPartner *string `json:"ConvertingBusinessPartner"`
// ExternalDocumentID        *string `json:"ExternalDocumentID"`
// }
//
// type ConversionProcessingPartner struct {
// ConvertingPartnerFunction *string `json:"ConvertingPartnerFunction"`
// ConvertedPartnerFunction  *string `json:"ConvertedPartnerFunction"`
// ConvertingBusinessPartner *string `json:"ConvertingBusinessPartner"`
// ConvertedBusinessPartner  *int    `json:"ConvertedBusinessPartner"`
// }
