package convert_complementer

import (
	dpfm_api_input_reader "convert-to-dpfm-orders-from-sap-purchase-order/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "convert-to-dpfm-orders-from-sap-purchase-order/DPFM_API_Processing_Formatter"
	"strings"
)

// Mapping Itemの処理
func (c *ConvertComplementer) ComplementMappingItemPricingElement(sdc *dpfm_api_input_reader.SDC, psdc *dpfm_api_processing_formatter.SDC) (*[]dpfm_api_processing_formatter.MappingItemPricingElement, error) {
	res, err := psdc.ConvertToMappingItemPricingElement(sdc)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *ConvertComplementer) CodeConversionItemPricingElement(sdc *dpfm_api_input_reader.SDC, psdc *dpfm_api_processing_formatter.SDC) (*[]dpfm_api_processing_formatter.CodeConversionItemPricingElement, error) {
	var data []dpfm_api_processing_formatter.CodeConversionItemPricingElement

	for _, item := range sdc.SAPPurchaseOrderHeader.SAPPurchaseOrderItem {
		for _, itemPricingElement := range item.SAPPurchaseOrderItemPricingElement {
			var dataKey []*dpfm_api_processing_formatter.CodeConversionKey
			var args []interface{}

			dataKey = append(dataKey, psdc.ConvertToCodeConversionKey(sdc, "ConditionSequentialNumber", "ConditionSequentialNumber", itemPricingElement.ConditionSequentialNumber))

			repeat := strings.Repeat("(?,?,?,?,?,?),", len(dataKey)-1) + "(?,?,?,?,?,?)"
			for _, v := range dataKey {
				args = append(args, v.SystemConvertTo, v.SystemConvertFrom, v.LabelConvertTo, v.LabelConvertFrom, v.CodeConvertFrom, v.BusinessPartner)
			}

			rows, err := c.db.Query(
				`SELECT CodeConversionID, SystemConvertTo, SystemConvertFrom, LabelConvertTo, LabelConvertFrom,
				CodeConvertFrom, CodeConvertTo, BusinessPartner
				FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_code_conversion_code_conversion_data
				WHERE (SystemConvertTo, SystemConvertFrom, LabelConvertTo, LabelConvertFrom, CodeConvertFrom, BusinessPartner) IN ( `+repeat+` );`, args...,
			)
			if err != nil {
				return nil, err
			}

			dataQueryGets, err := psdc.ConvertToCodeConversionQueryGets(rows)
			if err != nil {
				return nil, err
			}

			datum, err := psdc.ConvertToCodeConversionItemPricingElement(dataQueryGets)
			if err != nil {
				return nil, err
			}

			data = append(data, *datum)
		}
	}

	return &data, nil
}

func (c *ConvertComplementer) SetConditionType(sdc *dpfm_api_input_reader.SDC, psdc *dpfm_api_processing_formatter.SDC) *[]dpfm_api_processing_formatter.SetConditionType {
	var data []dpfm_api_processing_formatter.SetConditionType

	for _, item := range sdc.SAPPurchaseOrderHeader.SAPPurchaseOrderItem {
		for _, v := range item.SAPPurchaseOrderItemPricingElement {
			var datum *dpfm_api_processing_formatter.SetConditionType
			conditionType := v.ConditionType
			if conditionType != nil {
				if *conditionType == "PR00" || *conditionType == "MWST" {
					datum = psdc.ConvertToSetConditionType(conditionType)
				} else {
					datum = psdc.ConvertToSetConditionType(nil)
				}
			}
			data = append(data, *datum)
		}
	}

	return &data
}
