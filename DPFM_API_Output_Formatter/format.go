package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "convert-to-dpfm-orders-from-sap-purchase-order/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "convert-to-dpfm-orders-from-sap-purchase-order/DPFM_API_Processing_Formatter"
	"encoding/json"
)

func ConvertToHeader(
	sdc dpfm_api_input_reader.SDC,
	psdc dpfm_api_processing_formatter.SDC,
) (*Header, error) {
	mappingHeader := psdc.MappingHeader
	codeConversionHeader := psdc.CodeConversionHeader
	totalNetAmount := psdc.TotalNetAmount

	header := Header{}

	data, err := json.Marshal(mappingHeader)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &header)
	if err != nil {
		return nil, err
	}

	header.OrderID = codeConversionHeader.OrderID
	header.OrderType = codeConversionHeader.OrderType
	header.Seller = codeConversionHeader.Seller
	header.TotalNetAmount = totalNetAmount.TotalNetAmount

	return &header, nil
}

func ConvertToItem(
	sdc dpfm_api_input_reader.SDC,
	psdc dpfm_api_processing_formatter.SDC,
) (*[]Item, error) {
	var items []Item
	mappingItem := psdc.MappingItem
	codeConversionItem := psdc.CodeConversionItem
	conversionData := psdc.ConversionData

	for i := range *mappingItem {
		item := Item{}

		data, err := json.Marshal((*mappingItem)[i])
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(data, &item)
		if err != nil {
			return nil, err
		}

		for _, v := range *conversionData {
			if v.PurchaseOrder == (*mappingItem)[i].PurchaseOrder {
				item.OrderID = v.OrderID
				break
			}
		}
		item.OrderItem = (*codeConversionItem)[i].OrderItem
		item.OrderItemCategory = (*codeConversionItem)[i].OrderItemCategory
		item.Product = (*codeConversionItem)[i].Product
		item.ProductGroup = (*codeConversionItem)[i].ProductGroup

		items = append(items, item)
	}

	return &items, nil
}

func ConvertToItemPricingElement(
	sdc dpfm_api_input_reader.SDC,
	psdc dpfm_api_processing_formatter.SDC,
) (*[]ItemPricingElement, error) {
	var itemPricingElements []ItemPricingElement
	mappingItemPricingElement := psdc.MappingItemPricingElement
	codeConversionItemPricingElement := psdc.CodeConversionItemPricingElement
	conditionType := psdc.SetConditionType
	conversionData := psdc.ConversionData

	for i := range *mappingItemPricingElement {
		itemPricingElement := ItemPricingElement{}

		data, err := json.Marshal((*mappingItemPricingElement)[i])
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(data, &itemPricingElement)
		if err != nil {
			return nil, err
		}

		for _, v := range *conversionData {
			if v.PurchaseOrder == (*mappingItemPricingElement)[i].PurchaseOrder && v.PurchaseOrderItem == (*mappingItemPricingElement)[i].PurchaseOrderItem {
				itemPricingElement.OrderID = v.OrderID
				itemPricingElement.OrderItem = v.OrderItem
				// itemPricingElement.SupplyChainRelationshipID = v.SupplyChainRelationshipID
				break
			}
		}
		itemPricingElement.PricingProcedureCounter = *(*codeConversionItemPricingElement)[i].PricingProcedureCounter
		itemPricingElement.ConditionType = (*conditionType)[i].ConditionType

		itemPricingElements = append(itemPricingElements, itemPricingElement)
	}

	return &itemPricingElements, nil
}

func ConvertToItemScheduleLine(
	sdc dpfm_api_input_reader.SDC,
	psdc dpfm_api_processing_formatter.SDC,
) (*[]ItemScheduleLine, error) {
	var itemScheduleLines []ItemScheduleLine
	mappingItemScheduleLine := psdc.MappingItemScheduleLine
	codeConversionItemScheduleLine := psdc.CodeConversionItemScheduleLine
	conversionData := psdc.ConversionData

	for i := range *mappingItemScheduleLine {
		itemScheduleLine := ItemScheduleLine{}

		data, err := json.Marshal((*mappingItemScheduleLine)[i])
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(data, &itemScheduleLine)
		if err != nil {
			return nil, err
		}

		for _, v := range *conversionData {
			if v.PurchaseOrder == (*mappingItemScheduleLine)[i].PurchaseOrder && v.PurchaseOrderItem == (*mappingItemScheduleLine)[i].PurchaseOrderItem {
				itemScheduleLine.OrderID = v.OrderID
				itemScheduleLine.OrderItem = v.OrderItem
				itemScheduleLine.Product = v.Product
				break
			}
		}
		itemScheduleLine.ScheduleLine = (*codeConversionItemScheduleLine)[i].ScheduleLine

		itemScheduleLines = append(itemScheduleLines, itemScheduleLine)
	}

	return &itemScheduleLines, nil
}

func ConvertToAddress(
	sdc dpfm_api_input_reader.SDC,
	psdc dpfm_api_processing_formatter.SDC,
) (*Address, error) {
	mappingAddress := psdc.MappingAddress
	conversionData := psdc.ConversionData

	address := Address{}

	data, err := json.Marshal(mappingAddress)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &address)
	if err != nil {
		return nil, err
	}

	for _, v := range *conversionData {
		if v.PurchaseOrder == mappingAddress.PurchaseOrder {
			address.OrderID = v.OrderID
			break
		}
	}

	return &address, nil
}

func ConvertToPartner(
	sdc dpfm_api_input_reader.SDC,
	psdc dpfm_api_processing_formatter.SDC,
) (*[]Partner, error) {
	var partners []Partner
	mappingPartner := psdc.MappingPartner
	conversionData := psdc.ConversionData

	for i := range *mappingPartner {
		partner := Partner{}

		data, err := json.Marshal((*mappingPartner)[i])
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(data, &partner)
		if err != nil {
			return nil, err
		}

		for _, v := range *conversionData {
			if v.PurchaseOrder == (*mappingPartner)[i].PurchaseOrder {
				partner.OrderID = v.OrderID
				break
			}
		}

		partners = append(partners, partner)
	}

	return &partners, nil
}
