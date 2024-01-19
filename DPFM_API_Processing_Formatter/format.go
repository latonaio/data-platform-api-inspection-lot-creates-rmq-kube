package dpfm_api_processing_formatter

import dpfm_api_input_reader "data-platform-api-inspection-lot-creates-rmq-kube/DPFM_API_Input_Reader"

func ConvertToHeaderUpdates(header dpfm_api_input_reader.Header) *HeaderUpdates {
	data := header

	return &HeaderUpdates{
		InspectionLot:          	data.InspectionLot,
		InspectionLotDate:      	data.InspectionLotDate,
		InspectionLotHeaderText:	data.InspectionLotHeaderText,
		ExternalReferenceDocument:	data.ExternalReferenceDocument,
		IsLocked:				 	data.IsLocked,
	}
}

func ConvertToSpecGeneralUpdates(specGeneralUpdates dpfm_api_input_reader.SpecGeneral) *SpecGeneralUpdates {
	data := SpecGeneralUpdates

	return &SpecGeneralUpdates{
		InspectionLot: data.InspectionLot,
		HeatNumber:    data.HeatNumber,
		IsLocked:	   data.IsLocked,
	}
}

func ConvertToInspectionUpdates(inspectionUpdates dpfm_api_input_reader.Inspection) *InspectionUpdates {
	data := inspectionUpdates

	return &InspectionUpdates{
		InspectionLot:               				data.InspectionLot,
		Inspection:					 				data.Inspection,
		InspectionDate:				 				data.InspectionDate,
		InspectionTypeCertificateValueInText:		data.InspectionTypeCertificateValueInText,
		InspectionTypeCertificateValueInQuantity:	data.InspectionTypeCertificateValueInQuantity,	
		InspectionLotInspectionText: 				data.InspectionLotInspectionText,
		IsLocked:				 	 				data.IsLocked,
	}
}

func ConvertToOperationUpdates(operationUpdates dpfm_api_input_reader.Operation) *OperationUpdates {
	data := operationUpdates

	return &OperationUpdates{
		InspectionLot:               		data.InspectionLot,
		Operations:					 		data.Operations,
		OperationsItem:				 		data.OperationsItem,
		OperationID: 				 		data.OperationID,
		PlannedOperationStandardValue:		data.PlannedOperationStandardValue,
		PlannedOperationLowerValue:			data.PlannedOperationLowerValue,
		PlannedOperationUpperValue:			data.PlannedOperationUpperValue,
		PlannedOperationValueUnit:			data.PlannedOperationValueUnit,
		OperationErlstSchedldExecStrtDte:	data.OperationErlstSchedldExecStrtDte,
		OperationErlstSchedldExecStrtTme:	data.OperationErlstSchedldExecStrtTme,
		OperationErlstSchedldExecEndDte:	data.OperationErlstSchedldExecEndDte,
		OperationErlstSchedldExecEndTme:	data.OperationErlstSchedldExecEndTme,
		OperationActualExecutionStartDate:	data.OperationActualExecutionStartDate,
		OperationActualExecutionStartTime:	data.OperationActualExecutionStartTime,
		OperationActualExecutionEndDate:	data.OperationActualExecutionEndDate,
		OperationActualExecutionEndTime:	data.OperationActualExecutionEndTime,
		IsLocked:				 	 		data.IsLocked,
	}
}
