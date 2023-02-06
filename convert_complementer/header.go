package convert_complementer

import (
	dpfm_api_input_reader "convert-to-dpfm-orders-from-sap-purchase-order/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "convert-to-dpfm-orders-from-sap-purchase-order/DPFM_API_Processing_Formatter"
	"strings"
)

// Mapping Headerの処理
func (c *ConvertComplementer) ComplementMappingHeader(sdc *dpfm_api_input_reader.SDC, psdc *dpfm_api_processing_formatter.SDC) (*dpfm_api_processing_formatter.MappingHeader, error) {
	res, err := psdc.ConvertToMappingHeader(sdc)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *ConvertComplementer) CodeConversionHeader(sdc *dpfm_api_input_reader.SDC, psdc *dpfm_api_processing_formatter.SDC) (*dpfm_api_processing_formatter.CodeConversionHeader, error) {
	var dataKey []*dpfm_api_processing_formatter.CodeConversionKey
	var args []interface{}

	dataKey = append(dataKey, psdc.ConvertToCodeConversionKey(sdc, "PurchaseOrder", "OrderID", sdc.SAPPurchaseOrderHeader.PurchaseOrder))
	dataKey = append(dataKey, psdc.ConvertToCodeConversionKey(sdc, "PurchaseOrderType", "OrderType", sdc.SAPPurchaseOrderHeader.PurchaseOrderType))
	dataKey = append(dataKey, psdc.ConvertToCodeConversionKey(sdc, "Supplier", "Seller", sdc.SAPPurchaseOrderHeader.Supplier))

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

	data, err := psdc.ConvertToCodeConversionHeader(dataQueryGets)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *ConvertComplementer) TotalNetAmount(sdc *dpfm_api_input_reader.SDC, psdc *dpfm_api_processing_formatter.SDC) (*dpfm_api_processing_formatter.TotalNetAmount, error) {
	var totalNetAmount float32

	for _, v := range sdc.SAPPurchaseOrderHeader.SAPPurchaseOrderItem {
		netAmount, err := dpfm_api_processing_formatter.ParseFloat32Ptr(v.NetPriceAmount)
		if err != nil {
			return nil, err
		}

		if netAmount != nil {
			totalNetAmount += *netAmount
		}

	}

	data := psdc.ConvertToTotalNetAmount(&totalNetAmount)

	return data, nil
}
