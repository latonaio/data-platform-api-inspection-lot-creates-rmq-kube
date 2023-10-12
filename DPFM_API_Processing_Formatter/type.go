package dpfm_api_processing_formatter

type HeaderUpdates struct {
	InspectionLot                  int     `json:"InspectionLot"`
	InspectionLotHeaderText        *string `json:"InspectionLotHeaderText"`
}

type InspectionUpdates struct {
	InspectionLot	               int		`json:"InspectionLot"`
    Inspection	                   int		`json:"Inspection"`
    InspectionLotInspectionText	   *string	`json:"InspectionLotInspectionText"`
}
