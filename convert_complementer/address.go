package convert_complementer

import (
	dpfm_api_input_reader "convert-to-dpfm-orders-from-sap-purchase-order/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "convert-to-dpfm-orders-from-sap-purchase-order/DPFM_API_Processing_Formatter"
)

// Mapping Addressの処理
func (c *ConvertComplementer) ComplementMappingAddress(sdc *dpfm_api_input_reader.SDC, psdc *dpfm_api_processing_formatter.SDC) (*dpfm_api_processing_formatter.MappingAddress, error) {
	res, err := psdc.ConvertToMappingAddress(sdc)
	if err != nil {
		return nil, err
	}

	return res, nil
}
