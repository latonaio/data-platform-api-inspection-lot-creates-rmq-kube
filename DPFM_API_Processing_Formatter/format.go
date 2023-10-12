package dpfm_api_processing_formatter

import dpfm_api_input_reader "data-platform-api-inspection-lot-creates-rmq-kube/DPFM_API_Input_Reader"

func ConvertToHeaderUpdates(header dpfm_api_input_reader.Header) *HeaderUpdates {
	data := header

	return &HeaderUpdates{
			InspectionLot:                  data.InspectionLot,
			InspectionLotHeaderText:        data.InspectionLotHeaderText,
	}
}	

func ConvertToInspectionUpdates(inspectionUpdates dpfm_api_input_reader.Inspection) *InspectionUpdates {
	data := inspectionUpdates

	return &InspectionUpdates{
			InspectionLot:                            data.InspectionLot,
			InspectionLotInspectionText:              data.InspectionLotInspectionText,
	}
}
