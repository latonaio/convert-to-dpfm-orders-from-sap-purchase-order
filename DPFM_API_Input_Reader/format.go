package dpfm_api_input_reader

import (
	"convert-to-dpfm-orders-from-sap-purchase-order/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToHeader() *requests.Header {
	data := sdc.Header
	return &requests.Header{
		PurchaseOrder:               data.PurchaseOrder,
		CompanyCode:                 data.CompanyCode,
		PurchaseOrderType:           data.PurchaseOrderType,
		PurchasingProcessingStatus:  data.PurchasingProcessingStatus,
		CreationDate:                data.CreationDate,
		LastChangeDateTime:          data.LastChangeDateTime,
		Supplier:                    data.Supplier,
		Language:                    data.Language,
		PaymentTerms:                data.PaymentTerms,
		PurchasingOrganization:      data.PurchasingOrganization,
		PurchasingGroup:             data.PurchasingGroup,
		PurchaseOrderDate:           data.PurchaseOrderDate,
		DocumentCurrency:            data.DocumentCurrency,
		ExchangeRate:                data.ExchangeRate,
		ValidityStartDate:           data.ValidityStartDate,
		ValidityEndDate:             data.ValidityEndDate,
		SupplierRespSalesPersonName: data.SupplierRespSalesPersonName,
		SupplierPhoneNumber:         data.SupplierPhoneNumber,
		SupplyingPlant:              data.SupplyingPlant,
		IncotermsClassification:     data.IncotermsClassification,
		ManualSupplierAddressID:     data.ManualSupplierAddressID,
		AddressName:                 data.AddressName,
		AddressCityName:             data.AddressCityName,
		AddressFaxNumber:            data.AddressFaxNumber,
		AddressPostalCode:           data.AddressPostalCode,
		AddressStreetName:           data.AddressStreetName,
		AddressPhoneNumber:          data.AddressPhoneNumber,
		AddressRegion:               data.AddressRegion,
		AddressCountry:              data.AddressCountry,
	}
}

func (sdc *SDC) ConvertToItem(num int) *requests.Item {
	dataPurchaseOrder := sdc.Header
	data := sdc.Header.Item[num]

	return &requests.Item{
		PurchaseOrder:                  dataPurchaseOrder.PurchaseOrder,
		PurchaseOrderItem:              data.PurchaseOrderItem,
		PurchaseOrderItemText:          data.PurchaseOrderItemText,
		Plant:                          data.Plant,
		StorageLocation:                data.StorageLocation,
		MaterialGroup:                  data.MaterialGroup,
		PurchasingInfoRecord:           data.PurchasingInfoRecord,
		SupplierMaterialNumber:         data.SupplierMaterialNumber,
		OrderQuantity:                  data.OrderQuantity,
		PurchaseOrderQuantityUnit:      data.PurchaseOrderQuantityUnit,
		OrderPriceUnit:                 data.OrderPriceUnit,
		DocumentCurrency:               data.DocumentCurrency,
		NetPriceAmount:                 data.NetPriceAmount,
		NetPriceQuantity:               data.NetPriceQuantity,
		TaxCode:                        data.TaxCode,
		OverdelivTolrtdLmtRatioInPct:   data.OverdelivTolrtdLmtRatioInPct,
		UnlimitedOverdeliveryIsAllowed: data.UnlimitedOverdeliveryIsAllowed,
		UnderdelivTolrtdLmtRatioInPct:  data.UnderdelivTolrtdLmtRatioInPct,
		IsCompletelyDelivered:          data.IsCompletelyDelivered,
		IsFinallyInvoiced:              data.IsFinallyInvoiced,
		PurchaseOrderItemCategory:      data.PurchaseOrderItemCategory,
		AccountAssignmentCategory:      data.AccountAssignmentCategory,
		GoodsReceiptIsExpected:         data.GoodsReceiptIsExpected,
		GoodsReceiptIsNonValuated:      data.GoodsReceiptIsNonValuated,
		InvoiceIsExpected:              data.InvoiceIsExpected,
		InvoiceIsGoodsReceiptBased:     data.InvoiceIsGoodsReceiptBased,
		Customer:                       data.Customer,
		SupplierIsSubcontractor:        data.SupplierIsSubcontractor,
		ItemNetWeight:                  data.ItemNetWeight,
		ItemWeightUnit:                 data.ItemWeightUnit,
		IncotermsClassification:        data.IncotermsClassification,
		PurchaseRequisition:            data.PurchaseRequisition,
		PurchaseRequisitionItem:        data.PurchaseRequisitionItem,
		RequisitionerName:              data.RequisitionerName,
		Material:                       data.Material,
		InternationalArticleNumber:     data.InternationalArticleNumber,
		DeliveryAddressID:              data.DeliveryAddressID,
		DeliveryAddressName:            data.DeliveryAddressName,
		DeliveryAddressPostalCode:      data.DeliveryAddressPostalCode,
		DeliveryAddressStreetName:      data.DeliveryAddressStreetName,
		DeliveryAddressCityName:        data.DeliveryAddressCityName,
		DeliveryAddressRegion:          data.DeliveryAddressRegion,
		DeliveryAddressCountry:         data.DeliveryAddressCountry,
		PurchasingDocumentDeletionCode: data.PurchasingDocumentDeletionCode,
	}
}

func (sdc *SDC) ConvertToItemAccount(num int) *requests.ItemAccount {
	dataPurchaseOrder := sdc.Header
	dataItem := dataPurchaseOrder.Item[num]
	data := dataItem.ItemAccount[num]
	return &requests.ItemAccount{
		PurchaseOrder:           data.PurchaseOrder,
		PurchaseOrderItem:       data.PurchaseOrderItem,
		AccountAssignmentNumber: data.AccountAssignmentNumber,
		GLAccount:               data.GLAccount,
		BusinessArea:            data.BusinessArea,
		CostCenter:              data.CostCenter,
		SalesOrder:              data.SalesOrder,
		SalesOrderItem:          data.SalesOrderItem,
		SalesOrderScheduleLine:  data.SalesOrderScheduleLine,
		MasterFixedAsset:        data.MasterFixedAsset,
		FixedAsset:              data.FixedAsset,
		GoodsRecipientName:      data.GoodsRecipientName,
		UnloadingPointName:      data.UnloadingPointName,
		ControllingArea:         data.ControllingArea,
		CostObject:              data.CostObject,
		OrderID:                 data.OrderID,
		ProfitCenter:            data.ProfitCenter,
		WBSElement:              data.WBSElement,
		ProjectNetwork:          data.ProjectNetwork,
		FunctionalArea:          data.FunctionalArea,
		TaxCode:                 data.TaxCode,
		CostCtrActivityType:     data.CostCtrActivityType,
		IsDeleted:               data.IsDeleted,
	}
}

func (sdc *SDC) ConvertToItemPricingElement(num int) *requests.ItemPricingElement {
	dataPurchaseOrder := sdc.Header
	dataItem := dataPurchaseOrder.Item[num]
	data := dataItem.ItemPricingElement[num]
	return &requests.ItemPricingElement{
		PurchaseOrder:               dataPurchaseOrder.PurchaseOrder,
		PurchaseOrderItem:           dataItem.PurchaseOrderItem,
		PricingProcedureStep:        data.PricingProcedureStep,
		PricingProcedureCounter:     data.PricingProcedureCounter,
		PricingDocument:             data.PricingDocument,
		PricingDocumentItem:         data.PricingDocumentItem,
		ConditionType:               data.ConditionType,
		ConditionRateValue:          data.ConditionRateValue,
		ConditionCurrency:           data.ConditionCurrency,
		PriceDetnExchangeRate:       data.PriceDetnExchangeRate,
		TransactionCurrency:         data.TransactionCurrency,
		ConditionAmount:             data.ConditionAmount,
		ConditionQuantityUnit:       data.ConditionQuantityUnit,
		ConditionQuantity:           data.ConditionQuantity,
		ConditionApplication:        data.ConditionApplication,
		PricingDateTime:             data.PricingDateTime,
		ConditionCalculationType:    data.ConditionCalculationType,
		ConditionBaseValue:          data.ConditionBaseValue,
		ConditionToBaseQtyNmrtr:     data.ConditionToBaseQtyNmrtr,
		ConditionToBaseQtyDnmntr:    data.ConditionToBaseQtyDnmntr,
		ConditionCategory:           data.ConditionCategory,
		PricingScaleType:            data.PricingScaleType,
		ConditionOrigin:             data.ConditionOrigin,
		IsGroupCondition:            data.IsGroupCondition,
		ConditionSequentialNumber:   data.ConditionSequentialNumber,
		ConditionInactiveReason:     data.ConditionInactiveReason,
		PricingScaleBasis:           data.PricingScaleBasis,
		ConditionScaleBasisValue:    data.ConditionScaleBasisValue,
		ConditionScaleBasisCurrency: data.ConditionScaleBasisCurrency,
		ConditionScaleBasisUnit:     data.ConditionScaleBasisUnit,
		ConditionIsManuallyChanged:  data.ConditionIsManuallyChanged,
		ConditionRecord:             data.ConditionRecord,
	}
}

func (sdc *SDC) ConvertToItemScheduleLine(num int) *requests.ItemScheduleLine {
	dataPurchaseOrder := sdc.Header
	dataItem := dataPurchaseOrder.Item[num]
	data := dataItem.ItemScheduleLine[num]
	return &requests.ItemScheduleLine{
		PurchasingDocument:            dataPurchaseOrder.PurchaseOrder,
		PurchasingDocumentItem:        dataItem.PurchaseOrderItem,
		ScheduleLine:                  data.ScheduleLine,
		DelivDateCategory:             data.DelivDateCategory,
		ScheduleLineDeliveryDate:      data.ScheduleLineDeliveryDate,
		PurchaseOrderQuantityUnit:     data.PurchaseOrderQuantityUnit,
		ScheduleLineOrderQuantity:     data.ScheduleLineOrderQuantity,
		ScheduleLineDeliveryTime:      data.ScheduleLineDeliveryTime,
		PurchaseRequisition:           data.PurchaseRequisition,
		PurchaseRequisitionItem:       data.PurchaseRequisitionItem,
		ScheduleLineCommittedQuantity: data.ScheduleLineCommittedQuantity,
	}
}
