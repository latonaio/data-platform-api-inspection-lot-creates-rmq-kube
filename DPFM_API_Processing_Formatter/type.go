package dpfm_api_processing_formatter

type HeaderUpdates struct {
	InspectionLot           	int     `json:"InspectionLot"`
	InspectionLotDate       	string  `json:"InspectionLotDate"`
	InspectionLotHeaderText 	*string `json:"InspectionLotHeaderText"`
	ExternalReferenceDocument	*string `json:"ExternalReferenceDocument"`
	IsLocked                	*bool   `json:"IsLocked"`
}

type SpecGeneralUpdates struct {
	InspectionLot int    `json:"InspectionLot"`
	HeatNumber    string `json:"HeatNumber"`
	IsLocked      *bool  `json:"IsLocked"`
}

type InspectionUpdates struct {
	InspectionLot               				int      `json:"InspectionLot"`
	Inspection                  				int      `json:"Inspection"`
	InspectionDate								string 	 `json:"InspectionDate"`
    InspectionTypeCertificateValueInText	    *string	 `json:"InspectionTypeCertificateValueInText"`
    InspectionTypeCertificateValueInQuantity	*float32 `json:"InspectionTypeCertificateValueInQuantity"`
	InspectionLotInspectionText 				*string  `json:"InspectionLotInspectionText"`
	IsLocked                					*bool    `json:"IsLocked"`
}

type OperationUpdates struct {
	InspectionLot               					int      `json:"InspectionLot"`
	Operations                  					int      `json:"Operations"`
	OperationsItem              					int      `json:"OperationsItem"`
	OperationID                 					int      `json:"OperationID"`
	PlannedOperationStandardValue   				*float32 `json:"PlannedOperationStandardValue"`
	PlannedOperationLowerValue      				*float32 `json:"PlannedOperationLowerValue"`
	PlannedOperationUpperValue      				*float32 `json:"PlannedOperationUpperValue"`
	PlannedOperationValueUnit       				*string  `json:"PlannedOperationValueUnit"`
	OperationErlstSchedldExecStrtDte                *string  `json:"OperationErlstSchedldExecStrtDte"`
	OperationErlstSchedldExecStrtTme                *string  `json:"OperationErlstSchedldExecStrtTme"`
	OperationErlstSchedldExecEndDte                 *string  `json:"OperationErlstSchedldExecEndDte"`
	OperationErlstSchedldExecEndTme                 *string  `json:"OperationErlstSchedldExecEndTme"`
	OperationActualExecutionStartDate               *string  `json:"OperationActualExecutionStartDate"`
	OperationActualExecutionStartTime               *string  `json:"OperationActualExecutionStartTime"`
	OperationActualExecutionEndDate                 *string  `json:"OperationActualExecutionEndDate"`
	OperationActualExecutionEndTime                 *string  `json:"OperationActualExecutionEndTime"`
	IsLocked                						*bool    `json:"IsLocked"`
}
