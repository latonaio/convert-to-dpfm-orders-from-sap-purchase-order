package dpfm_api_processing_formatter

import (
	"context"
	"convert-to-dpfm-orders-from-sap-purchase-order/DPFM_API_Caller/requests"
	dpfm_api_input_reader "convert-to-dpfm-orders-from-sap-purchase-order/DPFM_API_Input_Reader"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	database "github.com/latonaio/golang-mysql-network-connector"
	"golang.org/x/xerrors"
)

type ProcessingFormatter struct {
	ctx context.Context
	db  *database.Mysql
	l   *logger.Logger
}

func NewProcessingFormatter(ctx context.Context, db *database.Mysql, l *logger.Logger) *ProcessingFormatter {
	return &ProcessingFormatter{
		ctx: ctx,
		db:  db,
		l:   l,
	}
}

func (p *ProcessingFormatter) ProcessingFormatter(
	sdc *dpfm_api_input_reader.SDC,
	psdc *ProcessingFormatterSDC,
) error {
	var err error
	var e error

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		psdc.Header, e = p.Header(sdc, psdc)
		if e != nil {
			err = e
			return
		}
		psdc.ConversionProcessingHeader, e = p.ConversionProcessingHeader(sdc, psdc)
		if e != nil {
			err = e
			return
		}

	}(&wg)

	wg.Wait()
	if err != nil {
		return err
	}

	p.l.Info(psdc)

	return nil
}

func (p *ProcessingFormatter) Header(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) (*Header, error) {
	data := sdc.Header
	dataHeaderItem := sdc.Header.Item[0]
	dataHeaderItemPricingElement := dataHeaderItem.ItemPricingElement[0]

	systemDate := getSystemDatePtr()

	priceDetnExchangeRate, err := parseFloat32Ptr(dataHeaderItemPricingElement.PriceDetnExchangeRate)
	if err != nil {
		return nil, xerrors.Errorf("Parse Error: %w", err)
	}
	accountingExchangeRate, err := parseFloat32Ptr(data.ExchangeRate)
	if err != nil {
		return nil, xerrors.Errorf("Parse Error: %w", err)
	}

	res := Header{
		ConvertingOrderID:         data.PurchaseOrder,
		OrderDate:                 data.PurchaseOrderDate,
		ConvertingOrderType:       data.PurchaseOrderType,
		Buyer:                     sdc.BusinessPartnerID,
		ConvertingSeller:          data.Supplier,
		BillToParty:               sdc.BusinessPartnerID,
		BillToCountry:             data.AddressCountry,
		Payer:                     sdc.BusinessPartnerID,
		CreationDate:              systemDate,
		LastChangeDate:            systemDate,
		OrderValidityStartDate:    data.ValidityStartDate,
		OrderValidityEndDate:      data.ValidityEndDate,
		TransactionCurrency:       data.DocumentCurrency,
		PricingDate:               dataHeaderItemPricingElement.PricingDateTime,
		PriceDetnExchangeRate:     priceDetnExchangeRate,
		Incoterms:                 data.IncotermsClassification,
		PaymentTerms:              data.PaymentTerms,
		AccountingExchangeRate:    accountingExchangeRate,
		HeaderBlockStatus:         getBoolPtr(false),
		HeaderBillingBlockStatus:  getBoolPtr(false),
		HeaderDeliveryBlockStatus: getBoolPtr(false),
		IsCancelled:               getBoolPtr(false),
		IsDeleted:                 getBoolPtr(false),
	}

	return &res, nil
}

func (p *ProcessingFormatter) ConversionProcessingHeader(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) (*ConversionProcessingHeader, error) {
	dataKey := make([]*ConversionProcessingKey, 0)

	dataKey = append(dataKey, p.ConversionProcessingKey(sdc, "PurchaseOrder", "OrderID", psdc.Header.ConvertingOrderID))
	dataKey = append(dataKey, p.ConversionProcessingKey(sdc, "PurchaseOrderType", "OrderType", psdc.Header.ConvertingOrderType))
	dataKey = append(dataKey, p.ConversionProcessingKey(sdc, "Supplier", "Seller", psdc.Header.ConvertingSeller))

	dataQueryGets, err := p.ConversionProcessingCommonQueryGets(dataKey)
	if err != nil {
		return nil, xerrors.Errorf("ConversionProcessing Error: %w", err)
	}

	data, err := p.ConvertToConversionProcessingHeader(dataKey, dataQueryGets)
	if err != nil {
		return nil, xerrors.Errorf("ConvertToConversionProcessing Error: %w", err)
	}

	return data, nil
}

func (psdc *ProcessingFormatter) ConvertToConversionProcessingHeader(conversionProcessingKey []*ConversionProcessingKey, conversionProcessingCommonQueryGets []*ConversionProcessingCommonQueryGets) (*ConversionProcessingHeader, error) {
	data := make(map[string]*ConversionProcessingCommonQueryGets, len(conversionProcessingCommonQueryGets))
	for _, v := range conversionProcessingCommonQueryGets {
		data[v.LabelConvertTo] = v
	}

	for _, v := range conversionProcessingKey {
		if _, ok := data[v.LabelConvertTo]; !ok {
			return nil, xerrors.Errorf("%s is not in the database", v.LabelConvertTo)
		}
	}

	pm := &requests.ConversionProcessingHeader{}

	pm.ConvertingOrderID = data["OrderID"].CodeConvertFromString
	pm.ConvertedOrderID = data["OrderID"].CodeConvertToInt
	pm.ConvertingOrderType = data["OrderType"].CodeConvertFromString
	pm.ConvertedOrderType = data["OrderType"].CodeConvertFromString
	pm.ConvertingSeller = data["Seller"].CodeConvertFromString
	pm.ConvertedSeller = data["Seller"].CodeConvertToInt

	res := &ConversionProcessingHeader{
		ConvertingOrderID:   pm.ConvertingOrderID,
		ConvertedOrderID:    pm.ConvertedOrderID,
		ConvertingOrderType: pm.ConvertingOrderType,
		ConvertedOrderType:  pm.ConvertedOrderType,
		ConvertingSeller:    pm.ConvertingSeller,
		ConvertedSeller:     pm.ConvertedSeller,
	}

	return res, nil
}

// func (p *ProcessingFormatter) Item(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) ([]*Item, error) {
// 	res := make([]*Item, 0)
// 	dataHeader := psdc.Header
// 	dataPreparingHeaderPartner := psdc.PreparingHeaderPartner
// 	data := sdc.Header.Item

// 	systemDate := getSystemDatePtr()

// 	for _, data := range data {
// 		requestedQuantity, err := parseFloat32Ptr(data.RequestedQuantity)
// 		if err != nil {
// 			return nil, err
// 		}
// 		confdDelivQtyInOrderQtyUnit, err := parseFloat32Ptr(data.ConfdDelivQtyInOrderQtyUnit)
// 		if err != nil {
// 			return nil, err
// 		}
// 		itemGrossWeight, err := parseFloat32Ptr(data.ItemGrossWeight)
// 		if err != nil {
// 			return nil, err
// 		}
// 		itemNetWeight, err := parseFloat32Ptr(data.ItemNetWeight)
// 		if err != nil {
// 			return nil, err
// 		}
// 		netAmount, err := parseFloat32Ptr(data.NetAmount)
// 		if err != nil {
// 			return nil, err
// 		}
// 		taxAmount, err := parseFloat32Ptr(data.TaxAmount)
// 		if err != nil {
// 			return nil, err
// 		}
// 		grossAmount, err := parseFloat32Ptr(data.GrossAmount)
// 		if err != nil {
// 			return nil, err
// 		}

// 		res = append(res, &Item{
// 			ConvertingOrderID:                       dataHeader.ConvertingOrderID,
// 			ConvertingOrderItem:                     data.SalesOrderItem,
// 			ConvertingOrderItemCategory:             data.SalesOrderItemCategory,
// 			OrderItemTextBySeller:                   data.SalesOrderItemText,
// 			ConvertingProduct:                       data.Material,
// 			ConvertingProductGroup:                  data.MaterialGroup,
// 			BaseUnit:                                data.RequestedQuantityUnit,
// 			PricingDate:                             data.PricingDate,
// 			PriceDetnExchangeRate:                   dataHeader.PriceDetnExchangeRate,
// 			RequestedDeliveryDate:                   dataHeader.RequestedDeliveryDate,
// 			ConvertingDeliverToParty:                dataPreparingHeaderPartner.ConvertingDeliverToParty,
// 			DeliverFromParty:                        sdc.BusinessPartnerID,
// 			CreationDate:                            systemDate,
// 			LastChangeDate:                          systemDate,
// 			DeliverFromPlant:                        data.ShippingPoint,
// 			DeliverFromPlantStorageLocation:         data.StorageLocation,
// 			DeliveryUnit:                            data.OrderQuantityUnit,
// 			StockConfirmationBusinessPartner:        sdc.BusinessPartnerID,
// 			StockConfirmationPlant:                  data.ProductionPlant,
// 			StockConfirmationPlantBatch:             data.Batch,
// 			OrderQuantityInDeliveryUnit:             requestedQuantity,
// 			ConfirmedOrderQuantityInBaseUnit:        confdDelivQtyInOrderQtyUnit,
// 			ItemWeightUnit:                          data.ItemWeightUnit,
// 			ItemGrossWeight:                         itemGrossWeight,
// 			ItemNetWeight:                           itemNetWeight,
// 			NetAmount:                               netAmount,
// 			TaxAmount:                               taxAmount,
// 			GrossAmount:                             grossAmount,
// 			InvoiceDocumentDate:                     data.BillingDocumentDate,
// 			ProductionPlantBusinessPartner:          sdc.BusinessPartnerID,
// 			ProductionPlant:                         data.ProductionPlant,
// 			Incoterms:                               data.IncotermsClassification,
// 			TransactionTaxClassification:            dataHeader.ConvertingCustomerTaxClassification1,
// 			ProductTaxClassificationBillToCountry:   data.ProductTaxClassification1,
// 			ProductTaxClassificationBillFromCountry: data.ProductTaxClassification1,
// 			AccountAssignmentGroup:                  dataHeader.ConvertingAccountAssignmentGroup,
// 			ProductAccountAssignmentGroup:           data.MatlAccountAssignmentGroup,
// 			PaymentTerms:                            data.CustomerPaymentTerms,
// 			TaxCode:                                 data.TaxCode,
// 			ItemBlockStatus:                         getBoolPtr(false),
// 			ItemDeliveryBlockStatus:                 getBoolPtr(false),
// 			ItemBillingBlockStatus:                  getBoolPtr(false),
// 			IsCancelled:                             getBoolPtr(false),
// 			IsDeleted:                               getBoolPtr(false),
// 		})
// 	}

// 	return res, nil
// }

// func (p *ProcessingFormatter) ConversionProcessingItem(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) ([]*ConversionProcessingItem, error) {
// 	data := make([]*ConversionProcessingItem, 0)

// 	for _, item := range psdc.Item {
// 		dataKey := make([]*ConversionProcessingKey, 0)

// 		dataKey = append(dataKey, p.ConversionProcessingKey(sdc, "SalesOrderItem", "OrderItem", item.ConvertingOrderItem))
// 		dataKey = append(dataKey, p.ConversionProcessingKey(sdc, "SalesOrderItemCategory", "OrderItemCategory", item.ConvertingOrderItemCategory))
// 		dataKey = append(dataKey, p.ConversionProcessingKey(sdc, "Material", "Product", item.ConvertingProduct))
// 		dataKey = append(dataKey, p.ConversionProcessingKey(sdc, "MaterialGroup", "ProductGroup", item.ConvertingProductGroup))
// 		dataKey = append(dataKey, p.ConversionProcessingKey(sdc, "Customer", "DeliverToParty", item.ConvertingDeliverToParty))

// 		dataQueryGets, err := p.ConversionProcessingCommonQueryGets(dataKey)
// 		if err != nil {
// 			return nil, xerrors.Errorf("ConversionProcessing Error: %w", err)
// 		}

// 		datum, err := p.ConvertToConversionProcessingItem(dataKey, dataQueryGets)
// 		if err != nil {
// 			return nil, xerrors.Errorf("ConvertToConversionProcessing Error: %w", err)
// 		}

// 		data = append(data, datum)
// 	}

// 	return data, nil
// }

// func (p *ProcessingFormatter) ConvertToConversionProcessingItem(conversionProcessingKey []*ConversionProcessingKey, conversionProcessingCommonQueryGets []*ConversionProcessingCommonQueryGets) (*ConversionProcessingItem, error) {
// 	data := make(map[string]*ConversionProcessingCommonQueryGets, len(conversionProcessingCommonQueryGets))
// 	for _, v := range conversionProcessingCommonQueryGets {
// 		data[v.LabelConvertTo] = v
// 	}

// 	for _, v := range conversionProcessingKey {
// 		if _, ok := data[v.LabelConvertTo]; !ok {
// 			return nil, xerrors.Errorf("%s is not in the database", v.LabelConvertTo)
// 		}
// 	}

// 	pm := &requests.ConversionProcessingItem{}

// 	pm.ConvertingOrderItem = data["OrderItem"].CodeConvertFromString
// 	pm.ConvertedOrderItem = data["OrderItem"].CodeConvertToInt
// 	pm.ConvertingOrderItemCategory = data["OrderItemCategory"].CodeConvertFromString
// 	pm.ConvertedOrderItemCategory = data["OrderItemCategory"].CodeConvertFromString
// 	pm.ConvertingProduct = data["Product"].CodeConvertFromString
// 	pm.ConvertedProduct = data["Product"].CodeConvertFromString
// 	pm.ConvertingProductGroup = data["ProductGroup"].CodeConvertFromString
// 	pm.ConvertedProductGroup = data["ProductGroup"].CodeConvertFromString
// 	pm.ConvertingDeliverToParty = data["DeliverToParty"].CodeConvertFromString
// 	pm.ConvertedDeliverToParty = data["DeliverToParty"].CodeConvertToInt

// 	res := &ConversionProcessingItem{
// 		ConvertingOrderItem:         pm.ConvertingOrderItem,
// 		ConvertedOrderItem:          pm.ConvertedOrderItem,
// 		ConvertingOrderItemCategory: pm.ConvertingOrderItemCategory,
// 		ConvertedOrderItemCategory:  pm.ConvertedOrderItemCategory,
// 		ConvertingProduct:           pm.ConvertingProduct,
// 		ConvertedProduct:            pm.ConvertedProduct,
// 		ConvertingProductGroup:      pm.ConvertingProductGroup,
// 		ConvertedProductGroup:       pm.ConvertedProductGroup,
// 		ConvertingDeliverToParty:    pm.ConvertingDeliverToParty,
// 		ConvertedDeliverToParty:     pm.ConvertedDeliverToParty,
// 	}

// 	return res, nil
// }

// func (p *ProcessingFormatter) ItemPricingElement(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) ([]*ItemPricingElement, error) {
// 	res := make([]*ItemPricingElement, 0)
// 	dataHeader := psdc.Header
// 	dataItem := psdc.Item

// 	for i, dataItem := range dataItem {
// 		data := sdc.Header.Item[i].ItemPricingElement
// 		for _, data := range data {
// 			conditionRateValue, err := parseFloat32Ptr(data.ConditionRateValue)
// 			if err != nil {
// 				return nil, err
// 			}
// 			conditionQuantity, err := parseFloat32Ptr(data.ConditionQuantity)
// 			if err != nil {
// 				return nil, err
// 			}
// 			conditionAmount, err := parseFloat32Ptr(data.ConditionAmount)
// 			if err != nil {
// 				return nil, err
// 			}

// 			res = append(res, &ItemPricingElement{
// 				ConvertingOrderID:                   dataHeader.ConvertingOrderID,
// 				ConvertingOrderItem:                 dataItem.ConvertingOrderItem,
// 				ConvertingBuyer:                     dataHeader.ConvertingBuyer,
// 				Seller:                              sdc.BusinessPartnerID,
// 				ConvertingConditionRecord:           data.ConditionRecord,
// 				ConvertingConditionSequentialNumber: data.ConditionSequentialNumber,
// 				PricingDate:                         dataItem.PricingDate,
// 				ConditionRateValue:                  conditionRateValue,
// 				ConditionCurrency:                   data.ConditionCurrency,
// 				ConditionQuantity:                   conditionQuantity,
// 				ConditionQuantityUnit:               data.ConditionQuantityUnit,
// 				TaxCode:                             data.TaxCode,
// 				ConditionAmount:                     conditionAmount,
// 				TransactionCurrency:                 data.TransactionCurrency,
// 				ConditionIsManuallyChanged:          getBoolPtr(true),
// 			})
// 		}
// 	}

// 	return res, nil
// }

// func (p *ProcessingFormatter) ConversionProcessingItemPricingElement(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) ([]*ConversionProcessingItemPricingElement, error) {
// 	data := make([]*ConversionProcessingItemPricingElement, 0)

// 	for _, itemPricingElement := range psdc.ItemPricingElement {
// 		dataKey := make([]*ConversionProcessingKey, 0)

// 		dataKey = append(dataKey, p.ConversionProcessingKey(sdc, "ConditionRecord", "ConditionRecord", itemPricingElement.ConvertingConditionRecord))
// 		dataKey = append(dataKey, p.ConversionProcessingKey(sdc, "ConditionSequentialNumber", "ConditionSequentialNumber", itemPricingElement.ConvertingConditionSequentialNumber))

// 		dataQueryGets, err := p.ConversionProcessingCommonQueryGets(dataKey)
// 		if err != nil {
// 			return nil, xerrors.Errorf("ConversionProcessing Error: %w", err)
// 		}

// 		datum, err := p.ConvertToConversionProcessingItemPricingElement(dataKey, dataQueryGets)
// 		if err != nil {
// 			return nil, xerrors.Errorf("ConvertToConversionProcessing Error: %w", err)
// 		}

// 		data = append(data, datum)
// 	}

// 	return data, nil
// }

// func (p *ProcessingFormatter) ConvertToConversionProcessingItemPricingElement(conversionProcessingKey []*ConversionProcessingKey, conversionProcessingCommonQueryGets []*ConversionProcessingCommonQueryGets) (*ConversionProcessingItemPricingElement, error) {
// 	data := make(map[string]*ConversionProcessingCommonQueryGets, len(conversionProcessingCommonQueryGets))
// 	for _, v := range conversionProcessingCommonQueryGets {
// 		data[v.LabelConvertTo] = v
// 	}

// 	for _, v := range conversionProcessingKey {
// 		if _, ok := data[v.LabelConvertTo]; !ok {
// 			return nil, xerrors.Errorf("%s is not in the database", v.LabelConvertTo)
// 		}
// 	}

// 	pm := &requests.ConversionProcessingItemPricingElement{}

// 	pm.ConvertingConditionRecord = data["ConditionRecord"].CodeConvertFromString
// 	pm.ConvertedConditionRecord = data["ConditionRecord"].CodeConvertToInt
// 	pm.ConvertingConditionSequentialNumber = data["ConditionSequentialNumber"].CodeConvertFromString
// 	pm.ConvertedConditionSequentialNumber = data["ConditionSequentialNumber"].CodeConvertToInt

// 	res := &ConversionProcessingItemPricingElement{
// 		ConvertingConditionRecord:           pm.ConvertingConditionRecord,
// 		ConvertedConditionRecord:            pm.ConvertedConditionRecord,
// 		ConvertingConditionSequentialNumber: pm.ConvertingConditionSequentialNumber,
// 		ConvertedConditionSequentialNumber:  pm.ConvertedConditionSequentialNumber,
// 	}

// 	return res, nil
// }

// func (p *ProcessingFormatter) ItemScheduleLine(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) ([]*ItemScheduleLine, error) {
// 	res := make([]*ItemScheduleLine, 0)
// 	dataHeader := psdc.Header
// 	dataItem := psdc.Item

// 	for i, dataItem := range dataItem {
// 		data := sdc.Header.Item[i].ItemScheduleLine
// 		for _, data := range data {
// 			scheduleLineOrderQuantity, err := parseFloat32Ptr(data.ScheduleLineOrderQuantity)
// 			if err != nil {
// 				return nil, err
// 			}
// 			confdOrderQtyByMatlAvailCheck, err := parseFloat32Ptr(data.ConfdOrderQtyByMatlAvailCheck)
// 			if err != nil {
// 				return nil, err
// 			}
// 			deliveredQtyInOrderQtyUnit, err := parseFloat32Ptr(data.DeliveredQtyInOrderQtyUnit)
// 			if err != nil {
// 				return nil, err
// 			}
// 			openConfdDelivQtyInOrdQtyUnit, err := parseFloat32Ptr(data.OpenConfdDelivQtyInOrdQtyUnit)
// 			if err != nil {
// 				return nil, err
// 			}

// 			res = append(res, &ItemScheduleLine{
// 				ConvertingOrderID:                     dataHeader.ConvertingOrderID,
// 				ConvertingOrderItem:                   dataItem.ConvertingOrderItem,
// 				ConvertingScheduleLine:                data.ScheduleLine,
// 				Product:                               dataItem.ConvertingProduct,
// 				StockConfirmationBussinessPartner:     sdc.BusinessPartnerID,
// 				StockConfirmationPlant:                dataItem.ProductionPlant,
// 				StockConfirmationPlantBatch:           dataItem.StockConfirmationPlantBatch,
// 				RequestedDeliveryDate:                 data.RequestedDeliveryDate,
// 				ConfirmedDeliveryDate:                 data.ConfirmedDeliveryDate,
// 				OrderQuantityInBaseUnit:               scheduleLineOrderQuantity,
// 				ConfirmedOrderQuantityByPDTAvailCheck: confdOrderQtyByMatlAvailCheck,
// 				DeliveredQuantityInBaseUnit:           deliveredQtyInOrderQtyUnit,
// 				OpenConfirmedQuantityInBaseUnit:       openConfdDelivQtyInOrdQtyUnit,
// 			})
// 		}
// 	}

// 	return res, nil
// }

// func (p *ProcessingFormatter) ConversionProcessingItemScheduleLine(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) ([]*ConversionProcessingItemScheduleLine, error) {
// 	data := make([]*ConversionProcessingItemScheduleLine, 0)

// 	for _, itemScheduleLine := range psdc.ItemScheduleLine {
// 		dataKey := make([]*ConversionProcessingKey, 0)

// 		dataKey = append(dataKey, p.ConversionProcessingKey(sdc, "ScheduleLine", "ScheduleLine", itemScheduleLine.ConvertingScheduleLine))

// 		dataQueryGets, err := p.ConversionProcessingCommonQueryGets(dataKey)
// 		if err != nil {
// 			return nil, xerrors.Errorf("ConversionProcessing Error: %w", err)
// 		}

// 		datum, err := p.ConvertToConversionProcessingItemScheduleLine(dataKey, dataQueryGets)
// 		if err != nil {
// 			return nil, xerrors.Errorf("ConvertToConversionProcessing Error: %w", err)
// 		}

// 		data = append(data, datum)
// 	}

// 	return data, nil
// }

// func (psdc *ProcessingFormatter) ConvertToConversionProcessingItemScheduleLine(conversionProcessingKey []*ConversionProcessingKey, conversionProcessingCommonQueryGets []*ConversionProcessingCommonQueryGets) (*ConversionProcessingItemScheduleLine, error) {
// 	data := make(map[string]*ConversionProcessingCommonQueryGets, len(conversionProcessingCommonQueryGets))
// 	for _, v := range conversionProcessingCommonQueryGets {
// 		data[v.LabelConvertTo] = v
// 	}

// 	for _, v := range conversionProcessingKey {
// 		if _, ok := data[v.LabelConvertTo]; !ok {
// 			return nil, xerrors.Errorf("%s is not in the database", v.LabelConvertTo)
// 		}
// 	}

// 	pm := &requests.ConversionProcessingItemScheduleLine{}

// 	pm.ConvertingScheduleLine = data["ScheduleLine"].CodeConvertFromString
// 	pm.ConvertedScheduleLine = data["ScheduleLine"].CodeConvertToInt

// 	res := &ConversionProcessingItemScheduleLine{
// 		ConvertingScheduleLine: pm.ConvertingScheduleLine,
// 		ConvertedScheduleLine:  pm.ConvertedScheduleLine,
// 	}

// 	return res, nil
// }

// func (p *ProcessingFormatter) Address(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) ([]*Address, error) {
// 	res := make([]*Address, 0)
// 	dataHeader := psdc.Header

// 	res = append(res, &Address{
// 		ConvertingOrderID: dataHeader.ConvertingOrderID,
// 	})

// 	return res, nil
// }

// func (p *ProcessingFormatter) Partner(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) ([]*Partner, error) {
// 	res := make([]*Partner, 0)
// 	dataHeader := psdc.Header
// 	dataPreparingHeaderPartner := psdc.PreparingHeaderPartner

// 	for _, convertingPartnerFunction := range dataPreparingHeaderPartner.ConvertingPartnerFunction {
// 		res = append(res, &Partner{
// 			ConvertingOrderID:         dataHeader.ConvertingOrderID,
// 			ConvertingPartnerFunction: convertingPartnerFunction,
// 			ConvertingBusinessPartner: dataPreparingHeaderPartner.ConvertingCustomer,
// 			ExternalDocumentID:        dataHeader.ConvertingExternalDocumentID,
// 		})
// 	}
// 	return res, nil
// }

// func (p *ProcessingFormatter) ConversionProcessingPartner(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) ([]*ConversionProcessingPartner, error) {
// 	data := make([]*ConversionProcessingPartner, 0)

// 	for _, partner := range psdc.Partner {
// 		dataKey := make([]*ConversionProcessingKey, 0)

// 		dataKey = append(dataKey, p.ConversionProcessingKey(sdc, "PartnerFunction", "PartnerFunction", partner.ConvertingPartnerFunction))
// 		dataKey = append(dataKey, p.ConversionProcessingKey(sdc, "Customer", "BusinessPartner", partner.ConvertingBusinessPartner))
// 		dataKey = append(dataKey, p.ConversionProcessingKey(sdc, "Supplier", "BusinessPartner", partner.ConvertingBusinessPartner))

// 		dataQueryGets, err := p.ConversionProcessingCommonQueryGets(dataKey)
// 		if err != nil {
// 			return nil, xerrors.Errorf("ConversionProcessing Error: %w", err)
// 		}

// 		datum, err := p.ConvertToConversionProcessingPartner(dataKey, dataQueryGets)
// 		if err != nil {
// 			return nil, xerrors.Errorf("ConvertToConversionProcessing Error: %w", err)
// 		}

// 		data = append(data, datum)
// 	}

// 	return data, nil
// }

// func (p *ProcessingFormatter) ConvertToConversionProcessingPartner(conversionProcessingKey []*ConversionProcessingKey, conversionProcessingCommonQueryGets []*ConversionProcessingCommonQueryGets) (*ConversionProcessingPartner, error) {
// 	data := make(map[string]*ConversionProcessingCommonQueryGets, len(conversionProcessingCommonQueryGets))
// 	for _, v := range conversionProcessingCommonQueryGets {
// 		data[v.LabelConvertTo] = v
// 	}

// 	for _, v := range conversionProcessingKey {
// 		if _, ok := data[v.LabelConvertTo]; !ok {
// 			return nil, xerrors.Errorf("%s is not in the database", v.LabelConvertTo)
// 		}
// 	}

// 	pm := &requests.ConversionProcessingPartner{}

// 	pm.ConvertingPartnerFunction = data["PartnerFunction"].CodeConvertFromString
// 	pm.ConvertedPartnerFunction = data["PartnerFunction"].CodeConvertToString
// 	pm.ConvertingBusinessPartner = data["BusinessPartner"].CodeConvertFromString
// 	pm.ConvertedBusinessPartner = data["BusinessPartner"].CodeConvertToInt

// 	res := &ConversionProcessingPartner{
// 		ConvertingPartnerFunction: pm.ConvertingPartnerFunction,
// 		ConvertedPartnerFunction:  pm.ConvertedPartnerFunction,
// 		ConvertingBusinessPartner: pm.ConvertingBusinessPartner,
// 		ConvertedBusinessPartner:  pm.ConvertedBusinessPartner,
// 	}

// 	return res, nil
// }
