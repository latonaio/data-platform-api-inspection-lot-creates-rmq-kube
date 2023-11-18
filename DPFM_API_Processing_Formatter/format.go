package dpfm_api_processing_formatter

import dpfm_api_input_reader "data-platform-api-inspection-lot-creates-rmq-kube/DPFM_API_Input_Reader"

func ConvertToHeaderUpdates(header dpfm_api_input_reader.Header) *HeaderUpdates {
	data := header

	return &HeaderUpdates{
		InspectionLot:           data.InspectionLot,
		InspectionLotHeaderText: data.InspectionLotHeaderText,
	}
}

func ConvertToSpecGeneralUpdates(SpecGeneralUpdates dpfm_api_input_reader.SpecGeneral) *SpecGeneralUpdates {
	data := SpecGeneralUpdates

	return &SpecGeneralUpdates{
		InspectionLot: data.InspectionLot,
		HeatNumber:    data.HeatNumber,
	}
}

func ConvertToSpecDetailUpdates(specDetailUpdates dpfm_api_input_reader.SpecDetail) *SpecDetailUpdates {
	data := specDetailUpdates

	return &SpecDetailUpdates{
		InspectionLot: data.InspectionLot,
		SpecType:      data.SpecType,
	}
}

func ConvertToComponentCompositionUpdates(componentComposition dpfm_api_input_reader.ComponentComposition) *ComponentCompositionUpdates {
	data := componentComposition

	return &ComponentCompositionUpdates{
		InspectionLot:            data.InspectionLot,
		ComponentCompositionType: data.ComponentCompositionType,
	}
}

func ConvertToInspectionUpdates(inspectionUpdates dpfm_api_input_reader.Inspection) *InspectionUpdates {
	data := inspectionUpdates

	return &InspectionUpdates{
		InspectionLot:               data.InspectionLot,
		InspectionLotInspectionText: data.InspectionLotInspectionText,
	}
}

func ConvertToOperationUpdates(OperationUpdates dpfm_api_input_reader.Operation) *OperationUpdates {
	data := operationUpdates

	return &OperationUpdates{
		InspectionLot:  data.InspectionLot,
		Operations:     data.Operations,
		OperationsItem: data.OperationsItem,
		OperationID:    data.OperationID,
	}
}
