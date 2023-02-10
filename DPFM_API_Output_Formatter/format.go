package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "convert-to-dpfm-orders-from-sap-purchase-order/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "convert-to-dpfm-orders-from-sap-purchase-order/DPFM_API_Processing_Formatter"
)

func OutputFormatter(
	sdc *dpfm_api_input_reader.SDC,
	psdc *dpfm_api_processing_formatter.ProcessingFormatterSDC,
	osdc *Output,
) error {
	header := ConvertToHeader(*sdc, *psdc)
	// item := ConvertToItem(*sdc, *psdc)
	// itemPricingElement := ConvertToItemPricingElement(*sdc, *psdc)
	// itemScheduleLine := ConvertToItemScheduleLine(*sdc, *psdc)
	// address := ConvertToAddress(*sdc, *psdc)
	// partner := ConvertToPartner(*sdc, *psdc)

	osdc.Message = Message{
		Header: header,
		// Item:               item,
		// ItemPricingElement: itemPricingElement,
		// ItemScheduleLine:   itemScheduleLine,
		// Address:            address,
		// Partner:            partner,
	}

	return nil
}

func ConvertToHeader(
	sdc dpfm_api_input_reader.SDC,
	psdc dpfm_api_processing_formatter.ProcessingFormatterSDC,
) *Header {
	dataProcessingHeader := psdc.Header
	dataConversionProcessingHeader := psdc.ConversionProcessingHeader

	header := &Header{
		OrderID:                   *dataConversionProcessingHeader.ConvertedOrderID,
		OrderDate:                 dataProcessingHeader.OrderDate,
		OrderType:                 dataConversionProcessingHeader.ConvertedOrderType,
		Buyer:                     dataProcessingHeader.Buyer,
		Seller:                    dataConversionProcessingHeader.ConvertedSeller,
		BillToParty:               dataProcessingHeader.BillToParty,
		BillToCountry:             dataProcessingHeader.BillToCountry,
		Payer:                     dataProcessingHeader.Payer,
		CreationDate:              dataProcessingHeader.CreationDate,
		LastChangeDate:            dataProcessingHeader.LastChangeDate,
		OrderValidityStartDate:    dataProcessingHeader.OrderValidityStartDate,
		OrderValidityEndDate:      dataProcessingHeader.OrderValidityEndDate,
		TransactionCurrency:       dataProcessingHeader.TransactionCurrency,
		PricingDate:               dataProcessingHeader.PricingDate,
		PriceDetnExchangeRate:     dataProcessingHeader.PriceDetnExchangeRate,
		Incoterms:                 dataProcessingHeader.Incoterms,
		PaymentTerms:              dataProcessingHeader.PaymentTerms,
		AccountingExchangeRate:    dataProcessingHeader.AccountingExchangeRate,
		HeaderBlockStatus:         dataProcessingHeader.HeaderBlockStatus,
		HeaderBillingBlockStatus:  dataProcessingHeader.HeaderBillingBlockStatus,
		HeaderDeliveryBlockStatus: dataProcessingHeader.HeaderDeliveryBlockStatus,
		IsCancelled:               dataProcessingHeader.IsCancelled,
		IsDeleted:                 dataProcessingHeader.IsDeleted,
	}

	return header
}

// func ConvertToItem(
// 	sdc dpfm_api_input_reader.SDC,
// 	psdc dpfm_api_processing_formatter.ProcessingFormatterSDC,
// ) []*Item {
// 	dataProcessingItem := psdc.Item
// 	dataConversionProcessingHeader := psdc.ConversionProcessingHeader
// 	dataConversionProcessingItem := psdc.ConversionProcessingItem

// 	items := make([]*Item, 0)
// 	for i := range dataProcessingItem {
// 		item := &Item{
// 			OrderID:                                 *dataConversionProcessingHeader.ConvertedOrderID,
// 			OrderItem:                               *dataConversionProcessingItem[i].ConvertedOrderItem,
// 			OrderItemCategory:                       dataConversionProcessingItem[i].ConvertedOrderItemCategory,
// 			OrderItemTextBySeller:                   dataProcessingItem[i].OrderItemTextBySeller,
// 			Product:                                 dataConversionProcessingItem[i].ConvertedProduct,
// 			ProductGroup:                            dataConversionProcessingItem[i].ConvertedProductGroup,
// 			BaseUnit:                                dataProcessingItem[i].BaseUnit,
// 			PricingDate:                             dataProcessingItem[i].PricingDate,
// 			PriceDetnExchangeRate:                   dataProcessingItem[i].PriceDetnExchangeRate,
// 			RequestedDeliveryDate:                   dataProcessingItem[i].RequestedDeliveryDate,
// 			DeliverToParty:                          dataConversionProcessingItem[i].ConvertedDeliverToParty,
// 			DeliverFromParty:                        dataProcessingItem[i].DeliverFromParty,
// 			CreationDate:                            dataProcessingItem[i].CreationDate,
// 			LastChangeDate:                          dataProcessingItem[i].LastChangeDate,
// 			DeliverFromPlant:                        dataProcessingItem[i].DeliverFromPlant,
// 			DeliverFromPlantStorageLocation:         dataProcessingItem[i].DeliverFromPlantStorageLocation,
// 			DeliveryUnit:                            dataProcessingItem[i].DeliveryUnit,
// 			StockConfirmationBusinessPartner:        dataProcessingItem[i].StockConfirmationBusinessPartner,
// 			StockConfirmationPlant:                  dataProcessingItem[i].StockConfirmationPlant,
// 			StockConfirmationPlantBatch:             dataProcessingItem[i].StockConfirmationPlantBatch,
// 			OrderQuantityInDeliveryUnit:             dataProcessingItem[i].OrderQuantityInDeliveryUnit,
// 			ConfirmedOrderQuantityInBaseUnit:        dataProcessingItem[i].ConfirmedOrderQuantityInBaseUnit,
// 			ItemWeightUnit:                          dataProcessingItem[i].ItemWeightUnit,
// 			ItemGrossWeight:                         dataProcessingItem[i].ItemGrossWeight,
// 			ItemNetWeight:                           dataProcessingItem[i].ItemNetWeight,
// 			NetAmount:                               dataProcessingItem[i].NetAmount,
// 			TaxAmount:                               dataProcessingItem[i].TaxAmount,
// 			GrossAmount:                             dataProcessingItem[i].GrossAmount,
// 			InvoiceDocumentDate:                     dataProcessingItem[i].InvoiceDocumentDate,
// 			ProductionPlantBusinessPartner:          dataProcessingItem[i].ProductionPlantBusinessPartner,
// 			ProductionPlant:                         dataProcessingItem[i].ProductionPlant,
// 			Incoterms:                               dataProcessingItem[i].Incoterms,
// 			TransactionTaxClassification:            dataProcessingItem[i].TransactionTaxClassification,
// 			ProductTaxClassificationBillToCountry:   dataProcessingItem[i].ProductTaxClassificationBillToCountry,
// 			ProductTaxClassificationBillFromCountry: dataProcessingItem[i].ProductTaxClassificationBillFromCountry,
// 			AccountAssignmentGroup:                  dataProcessingItem[i].AccountAssignmentGroup,
// 			ProductAccountAssignmentGroup:           dataProcessingItem[i].ProductAccountAssignmentGroup,
// 			PaymentTerms:                            dataProcessingItem[i].PaymentTerms,
// 			TaxCode:                                 dataProcessingItem[i].TaxCode,
// 			ItemBlockStatus:                         dataProcessingItem[i].ItemBlockStatus,
// 			ItemDeliveryBlockStatus:                 dataProcessingItem[i].ItemDeliveryBlockStatus,
// 			ItemBillingBlockStatus:                  dataProcessingItem[i].ItemBillingBlockStatus,
// 			IsCancelled:                             dataProcessingItem[i].IsCancelled,
// 			IsDeleted:                               dataProcessingItem[i].IsDeleted,
// 		}

// 		items = append(items, item)
// 	}

// 	return items
// }

// func ConvertToItemPricingElement(
// 	sdc dpfm_api_input_reader.SDC,
// 	psdc dpfm_api_processing_formatter.ProcessingFormatterSDC,
// ) []*ItemPricingElement {
// 	dataProcessingItemPricingElement := psdc.ItemPricingElement
// 	dataConversionProcessingHeader := psdc.ConversionProcessingHeader
// 	dataConversionProcessingItem := psdc.ConversionProcessingItem
// 	dataConversionProcessingItemPricingElement := psdc.ConversionProcessingItemPricingElement

// 	dataConversionProcessingItemMap := make(map[string]*dpfm_api_processing_formatter.ConversionProcessingItem, len(dataConversionProcessingItem))
// 	for _, v := range dataConversionProcessingItem {
// 		dataConversionProcessingItemMap[*v.ConvertingOrderItem] = v
// 	}

// 	itemPricingElements := make([]*ItemPricingElement, 0)
// 	for i, v := range dataProcessingItemPricingElement {
// 		if _, ok := dataConversionProcessingItemMap[v.ConvertingOrderItem]; !ok {
// 			continue
// 		}

// 		itemPricingElements = append(itemPricingElements, &ItemPricingElement{
// 			OrderID:                    *dataConversionProcessingHeader.ConvertedOrderID,
// 			OrderItem:                  *dataConversionProcessingItemMap[v.ConvertingOrderItem].ConvertedOrderItem,
// 			Buyer:                      *dataConversionProcessingHeader.ConvertedBuyer,
// 			Seller:                     *dataProcessingItemPricingElement[i].Seller,
// 			ConditionRecord:            dataConversionProcessingItemPricingElement[i].ConvertedConditionRecord,
// 			ConditionSequentialNumber:  dataConversionProcessingItemPricingElement[i].ConvertedConditionSequentialNumber,
// 			PricingDate:                dataProcessingItemPricingElement[i].PricingDate,
// 			ConditionRateValue:         dataProcessingItemPricingElement[i].ConditionRateValue,
// 			ConditionCurrency:          dataProcessingItemPricingElement[i].ConditionCurrency,
// 			ConditionQuantity:          dataProcessingItemPricingElement[i].ConditionQuantity,
// 			ConditionQuantityUnit:      dataProcessingItemPricingElement[i].ConditionQuantityUnit,
// 			TaxCode:                    dataProcessingItemPricingElement[i].TaxCode,
// 			ConditionAmount:            dataProcessingItemPricingElement[i].ConditionAmount,
// 			TransactionCurrency:        dataProcessingItemPricingElement[i].TransactionCurrency,
// 			ConditionIsManuallyChanged: dataProcessingItemPricingElement[i].ConditionIsManuallyChanged,
// 		})
// 	}

// 	return itemPricingElements
// }

// func ConvertToItemScheduleLine(
// 	sdc dpfm_api_input_reader.SDC,
// 	psdc dpfm_api_processing_formatter.ProcessingFormatterSDC,
// ) []*ItemScheduleLine {
// 	dataProcessingItemScheduleLine := psdc.ItemScheduleLine
// 	dataConversionProcessingHeader := psdc.ConversionProcessingHeader
// 	dataConversionProcessingItem := psdc.ConversionProcessingItem
// 	dataConversionProcessingItemScheduleLine := psdc.ConversionProcessingItemScheduleLine

// 	dataConversionProcessingItemMap := make(map[string]*dpfm_api_processing_formatter.ConversionProcessingItem, len(dataConversionProcessingItem))
// 	for _, v := range dataConversionProcessingItem {
// 		dataConversionProcessingItemMap[*v.ConvertingOrderItem] = v
// 	}

// 	itemScheduleLines := make([]*ItemScheduleLine, 0)
// 	for i, v := range dataProcessingItemScheduleLine {
// 		if _, ok := dataConversionProcessingItemMap[v.ConvertingOrderItem]; !ok {
// 			continue
// 		}

// 		itemScheduleLines = append(itemScheduleLines, &ItemScheduleLine{
// 			OrderID:                               *dataConversionProcessingHeader.ConvertedOrderID,
// 			OrderItem:                             *dataConversionProcessingItemMap[v.ConvertingOrderItem].ConvertedOrderItem,
// 			ScheduleLine:                          *dataConversionProcessingItemScheduleLine[i].ConvertedScheduleLine,
// 			Product:                               dataConversionProcessingItem[i].ConvertedProduct,
// 			StockConfirmationBussinessPartner:     dataProcessingItemScheduleLine[i].StockConfirmationBussinessPartner,
// 			StockConfirmationPlant:                dataProcessingItemScheduleLine[i].StockConfirmationPlant,
// 			StockConfirmationPlantBatch:           dataProcessingItemScheduleLine[i].StockConfirmationPlantBatch,
// 			RequestedDeliveryDate:                 dataProcessingItemScheduleLine[i].RequestedDeliveryDate,
// 			ConfirmedDeliveryDate:                 dataProcessingItemScheduleLine[i].ConfirmedDeliveryDate,
// 			OrderQuantityInBaseUnit:               dataProcessingItemScheduleLine[i].OrderQuantityInBaseUnit,
// 			ConfirmedOrderQuantityByPDTAvailCheck: dataProcessingItemScheduleLine[i].ConfirmedOrderQuantityByPDTAvailCheck,
// 			DeliveredQuantityInBaseUnit:           dataProcessingItemScheduleLine[i].DeliveredQuantityInBaseUnit,
// 			OpenConfirmedQuantityInBaseUnit:       dataProcessingItemScheduleLine[i].OpenConfirmedQuantityInBaseUnit,
// 		})
// 	}
// 	return itemScheduleLines
// }

// func ConvertToAddress(
// 	sdc dpfm_api_input_reader.SDC,
// 	psdc dpfm_api_processing_formatter.ProcessingFormatterSDC,
// ) []*Address {
// 	dataConversionProcessingHeader := psdc.ConversionProcessingHeader

// 	addresses := make([]*Address, 0)
// 	addresses = append(addresses, &Address{
// 		OrderID: *dataConversionProcessingHeader.ConvertedOrderID,
// 	})

// 	return addresses
// }

// func ConvertToPartner(
// 	sdc dpfm_api_input_reader.SDC,
// 	psdc dpfm_api_processing_formatter.ProcessingFormatterSDC,
// ) []*Partner {
// 	dataProcessingPartner := psdc.Partner
// 	dataConversionProcessingHeader := psdc.ConversionProcessingHeader
// 	dataConversionProcessingPartner := psdc.ConversionProcessingPartner

// 	partners := make([]*Partner, 0)
// 	for i := range dataProcessingPartner {
// 		partners = append(partners, &Partner{
// 			OrderID:            *dataConversionProcessingHeader.ConvertedOrderID,
// 			PartnerFunction:    *dataConversionProcessingPartner[i].ConvertedPartnerFunction,
// 			BusinessPartner:    *dataConversionProcessingPartner[i].ConvertedBusinessPartner,
// 			ExternalDocumentID: dataProcessingPartner[i].ExternalDocumentID,
// 		})
// 	}

// 	return partners
// }
