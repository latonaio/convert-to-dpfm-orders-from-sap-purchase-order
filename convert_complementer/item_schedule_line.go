package convert_complementer

import (
	dpfm_api_input_reader "convert-to-dpfm-orders-from-sap-purchase-order/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "convert-to-dpfm-orders-from-sap-purchase-order/DPFM_API_Processing_Formatter"
	"strings"
)

// Mapping Item Schedule Lineの処理
func (c *ConvertComplementer) ComplementMappingItemScheduleLine(sdc *dpfm_api_input_reader.SDC, psdc *dpfm_api_processing_formatter.SDC) (*[]dpfm_api_processing_formatter.MappingItemScheduleLine, error) {
	res, err := psdc.ConvertToMappingItemScheduleLine(sdc)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *ConvertComplementer) CodeConversionItemScheduleLine(sdc *dpfm_api_input_reader.SDC, psdc *dpfm_api_processing_formatter.SDC) (*[]dpfm_api_processing_formatter.CodeConversionItemScheduleLine, error) {
	var data []dpfm_api_processing_formatter.CodeConversionItemScheduleLine

	for _, item := range sdc.SAPPurchaseOrderHeader.SAPPurchaseOrderItem {
		for _, itemScheduleLine := range item.SAPPurchaseOrderItemScheduleLine {
			var dataKey []*dpfm_api_processing_formatter.CodeConversionKey
			var args []interface{}

			dataKey = append(dataKey, psdc.ConvertToCodeConversionKey(sdc, "ScheduleLine", "ScheduleLine", itemScheduleLine.ScheduleLine))

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

			datum, err := psdc.ConvertToCodeConversionItemScheduleLine(dataQueryGets)
			if err != nil {
				return nil, err
			}

			data = append(data, *datum)
		}
	}

	return &data, nil
}
