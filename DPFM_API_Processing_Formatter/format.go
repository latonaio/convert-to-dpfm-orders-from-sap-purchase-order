package dpfm_api_processing_formatter

import (
	"convert-to-dpfm-orders-from-sap-purchase-order/DPFM_API_Caller/requests"
	dpfm_api_input_reader "convert-to-dpfm-orders-from-sap-purchase-order/DPFM_API_Input_Reader"
	"database/sql"
	"fmt"
	"strconv"
)

// データ連携基盤とSAP Purchase Orderとの項目マッピング
// Header
func (psdc *SDC) ConvertToMappingHeader(sdc *dpfm_api_input_reader.SDC) (*MappingHeader, error) {
	var res MappingHeader
	dataHeader := sdc.SAPPurchaseOrderHeader
	dataHeaderItem := sdc.SAPPurchaseOrderHeader.SAPPurchaseOrderItem

	systemDate := GetSystemDatePtr()

	for _, dataHeaderItem := range dataHeaderItem {
		dataHeaderItemPricingElement := dataHeaderItem.SAPPurchaseOrderItemPricingElement
		for _, dataHeaderItemPricingElement := range dataHeaderItemPricingElement {

			priceDetnExchangeRate, err := ParseFloat32Ptr(dataHeaderItemPricingElement.PriceDetnExchangeRate)
			if err != nil {
				return nil, err
			}
			exchangeRate, err := ParseFloat32Ptr(dataHeader.ExchangeRate)
			if err != nil {
				return nil, err
			}

			res = MappingHeader{
				OrderDate:              dataHeader.PurchaseOrderDate,
				Buyer:                  sdc.BusinessPartnerID,
				BillToParty:            sdc.BusinessPartnerID,
				Payer:                  sdc.BusinessPartnerID,
				BillToCountry:          dataHeader.AddressCountry,
				CreationDate:           systemDate,
				LastChangeDate:         systemDate,
				OrderValidityStartDate: dataHeader.ValidityStartDate,
				OrderValidityEndDate:   dataHeader.ValidityEndDate,
				TransactionCurrency:    dataHeader.DocumentCurrency,
				PricingDate:            dataHeaderItemPricingElement.PricingDateTime,
				PriceDetnExchangeRate:  priceDetnExchangeRate,
				Incoterms:              dataHeader.IncotermsClassification,
				PaymentTerms:           dataHeader.PaymentTerms,
				AccountingExchangeRate: exchangeRate,
			}
		}
	}

	return &res, nil
}

// Item
func (psdc *SDC) ConvertToMappingItem(sdc *dpfm_api_input_reader.SDC) (*[]MappingItem, error) {
	var res []MappingItem
	dataHeader := sdc.SAPPurchaseOrderHeader
	dataItem := sdc.SAPPurchaseOrderHeader.SAPPurchaseOrderItem

	systemDate := GetSystemDatePtr()

	for _, dataItem := range dataItem {
		dataItemPricingElement := dataItem.SAPPurchaseOrderItemPricingElement
		for _, dataItemPricingElement := range dataItemPricingElement {

			priceDetnExchangeRate, err := ParseFloat32Ptr(dataItemPricingElement.PriceDetnExchangeRate)
			if err != nil {
				return nil, err
			}
			orderQuantity, err := ParseFloat32Ptr(dataItem.OrderQuantity)
			if err != nil {
				return nil, err
			}
			itemNetWeight, err := ParseFloat32Ptr(dataItem.ItemNetWeight)
			if err != nil {
				return nil, err
			}
			netPriceAmount, err := ParseFloat32Ptr(dataItem.NetPriceAmount)
			if err != nil {
				return nil, err
			}
			exchangeRate, err := ParseFloat32Ptr(dataHeader.ExchangeRate)
			if err != nil {
				return nil, err
			}

			res = append(res, MappingItem{
				PurchaseOrder:                    dataHeader.PurchaseOrder,
				OrderItemTextByBuyer:             dataItem.PurchaseOrderItemText,
				BaseUnit:                         dataItem.PurchaseOrderQuantityUnit,
				PriceDetnExchangeRate:            priceDetnExchangeRate,
				DeliverToParty:                   sdc.BusinessPartnerID,
				CreationDate:                     systemDate,
				LastChangeDate:                   systemDate,
				DeliverToPlant:                   dataItem.Plant,
				DeliverToPlantStorageLocation:    dataItem.StorageLocation,
				DeliverFromPlant:                 dataHeader.SupplyingPlant,
				DeliveryUnit:                     dataItem.PurchaseOrderQuantityUnit,
				StockConfirmationBusinessPartner: sdc.BusinessPartnerID,
				StockConfirmationPlant:           dataItem.Plant,
				OrderQuantityInBaseUnit:          orderQuantity,
				ItemWeightUnit:                   dataItem.ItemWeightUnit,
				ItemNetWeight:                    itemNetWeight,
				NetAmount:                        netPriceAmount,
				ProductionPlantBusinessPartner:   sdc.BusinessPartnerID,
				ProductionPlant:                  dataHeader.SupplyingPlant,
				ProductionPlantStorageLocation:   dataItem.StorageLocation,
				Incoterms:                        dataHeader.IncotermsClassification,
				PaymentTerms:                     dataHeader.PaymentTerms,
				AccountingExchangeRate:           exchangeRate,
				TaxCode:                          dataItem.TaxCode,
			})
		}
	}

	return &res, nil
}

// ItemPricingElement
func (psdc *SDC) ConvertToMappingItemPricingElement(sdc *dpfm_api_input_reader.SDC) (*[]MappingItemPricingElement, error) {
	var res []MappingItemPricingElement
	dataHeader := sdc.SAPPurchaseOrderHeader
	dataItem := sdc.SAPPurchaseOrderHeader.SAPPurchaseOrderItem

	for _, dataItem := range dataItem {
		dataItemPricingElement := dataItem.SAPPurchaseOrderItemPricingElement
		for _, dataItemPricingElement := range dataItemPricingElement {
			conditionRateValue, err := ParseFloat32Ptr(dataItemPricingElement.ConditionRateValue)
			if err != nil {
				return nil, err
			}
			conditionQuantity, err := ParseFloat32Ptr(dataItemPricingElement.ConditionQuantity)
			if err != nil {
				return nil, err
			}
			conditionAmount, err := ParseFloat32Ptr(dataItemPricingElement.ConditionAmount)
			if err != nil {
				return nil, err
			}

			res = append(res, MappingItemPricingElement{
				PurchaseOrder:         dataHeader.PurchaseOrder,
				PurchaseOrderItem:     dataItem.PurchaseOrderItem,
				ConditionRateValue:    conditionRateValue,
				ConditionCurrency:     dataItemPricingElement.ConditionCurrency,
				ConditionQuantity:     conditionQuantity,
				ConditionQuantityUnit: dataItemPricingElement.ConditionQuantityUnit,
				TaxCode:               dataItem.TaxCode,
				ConditionAmount:       conditionAmount,
				TransactionCurrency:   dataItemPricingElement.TransactionCurrency,
			})
		}
	}

	return &res, nil
}

func (psdc *SDC) ConvertToMappingItemScheduleLine(sdc *dpfm_api_input_reader.SDC) (*[]MappingItemScheduleLine, error) {
	var res []MappingItemScheduleLine
	dataHeader := sdc.SAPPurchaseOrderHeader
	dataItem := sdc.SAPPurchaseOrderHeader.SAPPurchaseOrderItem

	for _, dataItem := range dataItem {
		dataItemScheduleLine := dataItem.SAPPurchaseOrderItemScheduleLine
		for _, dataItemScheduleLine := range dataItemScheduleLine {
			scheduleLineOrderQuantity, err := ParseFloat32Ptr(dataItemScheduleLine.ScheduleLineOrderQuantity)
			if err != nil {
				return nil, err
			}
			scheduleLineCommittedQuantity, err := ParseFloat32Ptr(dataItemScheduleLine.ScheduleLineCommittedQuantity)
			if err != nil {
				return nil, err
			}

			res = append(res, MappingItemScheduleLine{
				PurchaseOrder:                     dataHeader.PurchaseOrder,
				PurchaseOrderItem:                 dataItem.PurchaseOrderItem,
				Material:                          *dataItem.Material,
				StockConfirmationBussinessPartner: sdc.BusinessPartnerID,
				StockConfirmationPlant:            dataItem.Plant,
				RequestedDeliveryDate:             dataItemScheduleLine.ScheduleLineDeliveryDate,
				OrderQuantityInBaseUnit:           scheduleLineOrderQuantity,
				OpenConfirmedQuantityInBaseUnit:   scheduleLineCommittedQuantity,
			})
		}
	}

	return &res, nil
}

// Address
func (psdc *SDC) ConvertToMappingAddress(sdc *dpfm_api_input_reader.SDC) (*MappingAddress, error) {
	dataHeader := sdc.SAPPurchaseOrderHeader

	res := MappingAddress{
		PurchaseOrder: dataHeader.PurchaseOrder,
		PostalCode:    dataHeader.AddressPostalCode,
		LocalRegion:   dataHeader.AddressRegion,
		Country:       dataHeader.AddressCountry,
		StreetName:    dataHeader.AddressStreetName,
		CityName:      dataHeader.AddressCityName,
	}

	return &res, nil
}

// Partner
// 「ビジネスパートナのデフォルト値がセットされる」の対応が必要 ↓
func (psdc *SDC) ConvertToMappingPartner(sdc *dpfm_api_input_reader.SDC) (*[]MappingPartner, error) {
	var res []MappingPartner
	dataHeader := sdc.SAPPurchaseOrderHeader
	dataItem := sdc.SAPPurchaseOrderHeader.SAPPurchaseOrderItem

	for _, dataItem := range dataItem {
		dataItemAccount := dataItem.SAPPurchaseOrderItemAccount
		for _, dataItemAccount := range dataItemAccount {
			res = append(res, MappingPartner{
				PurchaseOrder:      dataHeader.PurchaseOrder,
				ExternalDocumentID: dataItemAccount.SalesOrder,
			})
		}
	}

	return &res, nil
}

// 番号処理
func (psdc *SDC) ConvertToCodeConversionKey(sdc *dpfm_api_input_reader.SDC, labelConvertFrom, labelConvertTo string, codeConvertFrom any) *CodeConversionKey {
	pm := &requests.CodeConversionKey{
		SystemConvertTo:   "DPFM",
		SystemConvertFrom: "SAP",
		BusinessPartner:   *sdc.BusinessPartnerID,
	}

	pm.LabelConvertFrom = labelConvertFrom
	pm.LabelConvertTo = labelConvertTo

	switch codeConvertFrom := codeConvertFrom.(type) {
	case string:
		pm.CodeConvertFrom = codeConvertFrom
	case int:
		pm.CodeConvertFrom = strconv.FormatInt(int64(codeConvertFrom), 10)
	case float32:
		pm.CodeConvertFrom = strconv.FormatFloat(float64(codeConvertFrom), 'f', -1, 32)
	case bool:
		pm.CodeConvertFrom = strconv.FormatBool(codeConvertFrom)
	case *string:
		if codeConvertFrom != nil {
			pm.CodeConvertFrom = *codeConvertFrom
		}
	case *int:
		if codeConvertFrom != nil {
			pm.CodeConvertFrom = strconv.FormatInt(int64(*codeConvertFrom), 10)
		}
	case *float32:
		if codeConvertFrom != nil {
			pm.CodeConvertFrom = strconv.FormatFloat(float64(*codeConvertFrom), 'f', -1, 32)
		}
	case *bool:
		if codeConvertFrom != nil {
			pm.CodeConvertFrom = strconv.FormatBool(*codeConvertFrom)
		}
	}

	data := pm
	res := CodeConversionKey{
		SystemConvertTo:   data.SystemConvertTo,
		SystemConvertFrom: data.SystemConvertFrom,
		LabelConvertTo:    data.LabelConvertTo,
		LabelConvertFrom:  data.LabelConvertFrom,
		CodeConvertFrom:   data.CodeConvertFrom,
		BusinessPartner:   data.BusinessPartner,
	}

	return &res
}

func (psdc *SDC) ConvertToCodeConversionQueryGets(rows *sql.Rows) (*[]CodeConversionQueryGets, error) {
	var res []CodeConversionQueryGets

	for i := 0; true; i++ {
		pm := &requests.CodeConversionQueryGets{}
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_code_conversion_code_conversion_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.CodeConversionID,
			&pm.SystemConvertTo,
			&pm.SystemConvertFrom,
			&pm.LabelConvertTo,
			&pm.LabelConvertFrom,
			&pm.CodeConvertFrom,
			&pm.CodeConvertTo,
			&pm.BusinessPartner,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, CodeConversionQueryGets{
			CodeConversionID:  data.CodeConversionID,
			SystemConvertTo:   data.SystemConvertTo,
			SystemConvertFrom: data.SystemConvertFrom,
			LabelConvertTo:    data.LabelConvertTo,
			LabelConvertFrom:  data.LabelConvertFrom,
			CodeConvertFrom:   data.CodeConvertFrom,
			CodeConvertTo:     data.CodeConvertTo,
			BusinessPartner:   data.BusinessPartner,
		})
	}

	return &res, nil
}

func (psdc *SDC) ConvertToCodeConversionHeader(dataQueryGets *[]CodeConversionQueryGets) (*CodeConversionHeader, error) {
	var err error

	dataMap := make(map[string]CodeConversionQueryGets, len(*dataQueryGets))
	for _, v := range *dataQueryGets {
		dataMap[v.LabelConvertTo] = v
	}

	pm := &requests.CodeConversionHeader{}

	pm.PurchaseOrder = dataMap["OrderID"].CodeConvertFrom
	pm.OrderID, err = ParseInt(dataMap["OrderID"].CodeConvertTo)
	if err != nil {
		return nil, err
	}
	pm.OrderType = GetStringPtr(dataMap["OrderType"].CodeConvertTo)
	pm.Supplier = GetStringPtr(dataMap["Seller"].CodeConvertFrom)
	pm.Seller, err = ParseIntPtr(GetStringPtr(dataMap["Seller"].CodeConvertTo))
	if err != nil {
		return nil, err
	}

	data := pm
	res := CodeConversionHeader{
		PurchaseOrder: data.PurchaseOrder,
		OrderID:       data.OrderID,
		OrderType:     data.OrderType,
		Supplier:      data.Supplier,
		Seller:        data.Seller,
	}

	return &res, nil
}

func (psdc *SDC) ConvertToCodeConversionItem(dataQueryGets *[]CodeConversionQueryGets) (*CodeConversionItem, error) {
	var err error

	dataMap := make(map[string]CodeConversionQueryGets, len(*dataQueryGets))
	for _, v := range *dataQueryGets {
		dataMap[v.LabelConvertTo] = v
	}

	pm := &requests.CodeConversionItem{}

	pm.PurchaseOrderItem = dataMap["OrderItem"].CodeConvertFrom
	pm.OrderItem, err = ParseInt(dataMap["OrderItem"].CodeConvertTo)
	if err != nil {
		return nil, err
	}
	pm.OrderItemCategory = GetStringPtr(dataMap["OrderItemCategory"].CodeConvertTo)
	pm.Material = dataMap["Product"].CodeConvertFrom
	pm.Product = GetStringPtr(dataMap["Product"].CodeConvertTo)
	pm.ProductGroup = GetStringPtr(dataMap["ProductGroup"].CodeConvertTo)
	if err != nil {
		return nil, err
	}

	data := pm
	res := CodeConversionItem{
		PurchaseOrderItem: data.PurchaseOrderItem,
		OrderItem:         data.OrderItem,
		OrderItemCategory: data.OrderItemCategory,
		Material:          data.Material,
		Product:           data.Product,
		ProductGroup:      data.ProductGroup,
	}

	return &res, nil
}

func (psdc *SDC) ConvertToConversionData() *[]ConversionData {
	var res []ConversionData

	for _, v := range *psdc.CodeConversionItem {
		pm := &requests.ConversionData{}

		pm.PurchaseOrder = psdc.CodeConversionHeader.PurchaseOrder
		pm.OrderID = psdc.CodeConversionHeader.OrderID
		pm.PurchaseOrderItem = v.PurchaseOrderItem
		pm.OrderItem = v.OrderItem
		pm.Supplier = psdc.CodeConversionHeader.Supplier
		pm.Seller = *psdc.CodeConversionHeader.Seller
		pm.Material = v.Material
		pm.Product = v.Product

		data := pm
		res = append(res, ConversionData{
			PurchaseOrder:     data.PurchaseOrder,
			OrderID:           data.OrderID,
			PurchaseOrderItem: data.PurchaseOrderItem,
			OrderItem:         data.OrderItem,
			Supplier:          data.Supplier,
			Seller:            data.Seller,
			Material:          data.Material,
			Product:           data.Product,
		})
	}

	return &res
}

func (psdc *SDC) ConvertToCodeConversionItemPricingElement(dataQueryGets *[]CodeConversionQueryGets) (*CodeConversionItemPricingElement, error) {
	var err error

	dataMap := make(map[string]CodeConversionQueryGets, len(*dataQueryGets))
	for _, v := range *dataQueryGets {
		dataMap[v.LabelConvertTo] = v
	}

	pm := &requests.CodeConversionItemPricingElement{}

	pm.PricingProcedureCounter, err = ParseIntPtr(GetStringPtr(dataMap["PricingProcedureCounter"].CodeConvertTo))
	if err != nil {
		return nil, err
	}

	data := pm
	res := CodeConversionItemPricingElement{
		PricingProcedureCounter: data.PricingProcedureCounter,
	}

	return &res, nil
}

func (psdc *SDC) ConvertToCodeConversionItemScheduleLine(dataQueryGets *[]CodeConversionQueryGets) (*CodeConversionItemScheduleLine, error) {
	var err error

	dataMap := make(map[string]CodeConversionQueryGets, len(*dataQueryGets))
	for _, v := range *dataQueryGets {
		dataMap[v.LabelConvertTo] = v
	}

	pm := &requests.CodeConversionItemScheduleLine{}

	pm.ScheduleLine, err = ParseInt(dataMap["ScheduleLine"].CodeConvertTo)
	if err != nil {
		return nil, err
	}

	data := pm
	res := CodeConversionItemScheduleLine{
		ScheduleLine: data.ScheduleLine,
	}

	return &res, nil
}

// 個別処理
func (psdc *SDC) ConvertToTotalNetAmount(totalNetAmount *float32) *TotalNetAmount {
	pm := &requests.TotalNetAmount{}

	pm.TotalNetAmount = totalNetAmount

	data := pm
	res := TotalNetAmount{
		TotalNetAmount: data.TotalNetAmount,
	}

	return &res
}

func (psdc *SDC) ConvertToSetConditionType(conditionType *string) *SetConditionType {
	pm := &requests.SetConditionType{}

	pm.ConditionType = conditionType

	data := pm
	res := SetConditionType{
		ConditionType: data.ConditionType,
	}

	return &res
}
