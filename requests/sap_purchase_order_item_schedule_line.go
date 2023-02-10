package requests

type SAPPurchaseOrderItemScheduleLine struct {
	PurchasingDocument            string  `json:"PurchasingDocument"`
	PurchasingDocumentItem        string  `json:"PurchasingDocumentItem"`
	ScheduleLine                  string  `json:"ScheduleLine"`
	DelivDateCategory             *string `json:"DelivDateCategory"`
	ScheduleLineDeliveryDate      *string `json:"ScheduleLineDeliveryDate"`
	PurchaseOrderQuantityUnit     *string `json:"PurchaseOrderQuantityUnit"`
	ScheduleLineOrderQuantity     *string `json:"ScheduleLineOrderQuantity"`
	ScheduleLineDeliveryTime      *string `json:"ScheduleLineDeliveryTime"`
	PurchaseRequisition           *string `json:"PurchaseRequisition"`
	PurchaseRequisitionItem       *string `json:"PurchaseRequisitionItem"`
	ScheduleLineCommittedQuantity *string `json:"ScheduleLineCommittedQuantity"`
}
