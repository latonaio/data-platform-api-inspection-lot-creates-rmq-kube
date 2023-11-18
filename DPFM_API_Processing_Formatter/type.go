package dpfm_api_processing_formatter

type HeaderUpdates struct {
	InspectionLot           int     `json:"InspectionLot"`
	InspectionLotHeaderText *string `json:"InspectionLotHeaderText"`
}

type SpecGeneralUpdates struct {
	InspectionLot int    `json:"InspectionLot"`
	HeatNumber    string `json:"HeatNumber"`
}

type SpecDetailUpdates struct {
	InspectionLot int    `json:"InspectionLot"`
	SpecType      string `json:"SpecType"`
}

type ComponentCompositionUpdates struct {
	InspectionLot            int    `json:"InspectionLot"`
	ComponentCompositionType string `json:"ComponentCompositionType"`
}

type InspectionUpdates struct {
	InspectionLot               int     `json:"InspectionLot"`
	Inspection                  int     `json:"Inspection"`
	InspectionLotInspectionText *string `json:"InspectionLotInspectionText"`
}

type OperationUpdates struct {
	InspectionLot  int `json:"InspectionLot"`
	Operations     int `json:"Operations"`
	OperationsItem int `json:"OperationsItem"`
	OperationID    int `json:"OperationID"`
}
