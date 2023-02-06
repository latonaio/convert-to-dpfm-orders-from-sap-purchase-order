package convert_complementer

import (
	"context"
	dpfm_api_input_reader "convert-to-dpfm-orders-from-sap-purchase-order/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "convert-to-dpfm-orders-from-sap-purchase-order/DPFM_API_Output_Formatter"
	dpfm_api_processing_data_formatter "convert-to-dpfm-orders-from-sap-purchase-order/DPFM_API_Processing_Formatter"
	"sync"

	database "github.com/latonaio/golang-mysql-network-connector"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type ConvertComplementer struct {
	ctx context.Context
	db  *database.Mysql
	l   *logger.Logger
}

func NewConvertComplementer(ctx context.Context, db *database.Mysql, l *logger.Logger) *ConvertComplementer {
	return &ConvertComplementer{
		ctx: ctx,
		db:  db,
		l:   l,
	}
}

func (c *ConvertComplementer) CreateSdc(
	sdc *dpfm_api_input_reader.SDC,
	psdc *dpfm_api_processing_data_formatter.SDC,
	osdc *dpfm_api_output_formatter.SDC,
) error {
	var err error
	var e error

	wg := sync.WaitGroup{}
	wg.Add(10)

	// Header
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 1-0. データ連携基盤Orders HeaderとSAP Purchase Orderとの項目マッピング変換
		psdc.MappingHeader, e = c.ComplementMappingHeader(sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	// Item
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 5-1. データ連携基盤Orders ItemとSAP Purchase Orderとの項目マッピング変換
		psdc.MappingItem, e = c.ComplementMappingItem(sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// Header

		psdc.TotalNetAmount, e = c.TotalNetAmount(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// <1-1. 番号変換＞
		psdc.CodeConversionHeader, e = c.CodeConversionHeader(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// Item

		// <5-1. 番号変換＞
		psdc.CodeConversionItem, e = c.CodeConversionItem(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 1-1-3．OrderIDの保持 (Orders Header), 2-4．OrderItemの保持 (Orders Item), 2-4．Productの保持 (Orders Item)
		psdc.ConversionData = c.ConversionData(sdc, psdc)
	}(&wg)

	// ItemPricingElement
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 8-1. データ連携基盤Orders Item Pricing ElementとSAP Purchase Orderとの項目マッピング変換
		psdc.MappingItemPricingElement, e = c.ComplementMappingItemPricingElement(sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// <8-2. 変換元のConditionTypeのセット>
		psdc.SetConditionType = c.SetConditionType(sdc, psdc)

		// <8-1. 番号変換＞
		psdc.CodeConversionItemPricingElement, e = c.CodeConversionItemPricingElement(sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	// ItemScheduleLine
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 4-1. データ連携基盤Orders Item Schedule LineとSAP Purchase Orderとの項目マッピング変換
		psdc.MappingItemScheduleLine, e = c.ComplementMappingItemScheduleLine(sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// <4-2. コード変換＞
		psdc.CodeConversionItemScheduleLine, e = c.CodeConversionItemScheduleLine(sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	// Address
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 6-1. データ連携基盤Orders AddressとSAP Purchase Orderとの項目マッピング変換
		psdc.MappingAddress, e = c.ComplementMappingAddress(sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	// Partner
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 2-1. データ連携基盤Orders PartnerとSAP Purchase Order との項目マッピング変換
		psdc.MappingPartner, e = c.ComplementMappingPartner(sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		if e != nil {
			err = e
			return
		}
	}(&wg)

	wg.Wait()
	if err != nil {
		return err
	}

	c.l.Info(psdc)
	osdc, err = c.SetValue(sdc, psdc, osdc)
	if err != nil {
		return err
	}

	return nil
}
