package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-inspection-lot-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-inspection-lot-creates-rmq-kube/DPFM_API_Output_Formatter"
	dpfm_api_processing_formatter "data-platform-api-inspection-lot-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"data-platform-api-inspection-lot-creates-rmq-kube/sub_func_complementer"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *DPFMAPICaller) createSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var header *dpfm_api_output_formatter.Header
	var specGeneral *[]dpfm_api_output_formatter.SpecGeneral
	var specDetail *[]dpfm_api_output_formatter.SpecDetail
	var componentComposition *[]dpfm_api_output_formatter.ComponentComposition
	var inspection *[]dpfm_api_output_formatter.Inspection
	var operation *[]dpfm_api_output_formatter.Operation
	for _, fn := range accepter {
		switch fn {
		case "Header":
			header = c.headerCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		case "SpecGeneral":
			specGeneral = c.specGeneralCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		case "SpecDetail":
			specDetail = c.specDetailCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		case "ComponentComposition":
			componentComposition = c.componentCompositionCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		case "Inspection":
			inspection = c.inspectionCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		case "Operation":
			operation = c.operationCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header:               header,
		SpecGeneral:          specGeneral,
		SpecDetail:           specDetail,
		ComponentComposition: componentComposition,
		Inspection:           inspection,
		Operation:            operation,
	}

	return data
}

func (c *DPFMAPICaller) updateSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var header *dpfm_api_output_formatter.Header
	var specGeneral *[]dpfm_api_output_formatter.SpecGeneral
	var specDetail *dpfm_api_output_formatter.SpecDetail
	var componentComposition *[]dpfm_api_output_formatter.ComponentComposition
	var inspection *dpfm_api_output_formatter.Inspection
	var operation *[]dpfm_api_output_formatter.Operation
	for _, fn := range accepter {
		switch fn {
		case "Header":
			header = c.headerUpdateSql(mtx, input, output, errs, log)
		case "Specgeneral":
			specGeneral = c.specGeneralUpdateSql(mtx, input, output, errs, log)
		case "Specdetail":
			specDetail = c.specDetailUpdateSql(mtx, input, output, errs, log)
		case "ComponentComposition":
			componentComposition = c.componentCompositionUpdateSql(mtx, input, output, errs, log)
		case "Inspection":
			inspection = c.inspectionUpdateSql(mtx, input, output, errs, log)
		case "Operation":
			operation = c.operationUpdateSql(mtx, input, output, errs, log)
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header:               header,
		SpecGeneral:          specGeneral,
		SpecDetail:           specDetail,
		ComponentComposition: componentComposition,
		Inspection:           inspection,
		Operation:            operation,
	}

	return data
}

func (c *DPFMAPICaller) headerCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.Header {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID

	headerData := subfuncSDC.Message.Header
	res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": headerData, "function": "InspectionLotHeader", "runtime_session_id": sessionID})
	if err != nil {
		err = xerrors.Errorf("rmq error: %w", err)
		return nil
	}
	res.Success()
	if !checkResult(res) {
		output.SQLUpdateResult = getBoolPtr(false)
		output.SQLUpdateError = "Header Data cannot insert"
		return nil
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToHeaderCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) specGeneralCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.SpecGeneral {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for _, specGeneralData := range *subfuncSDC.Message.SpecGeneral {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": specGeneralData, "function": "InspectionLotSpecGeneral", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "SpecGeneral Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToSpecGeneralCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) specDetailCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.SpecDetail {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for _, specDetailData := range *subfuncSDC.Message.SpecDetail {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": specDetailData, "function": "InspectionLotSpecDetail", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "SpecDetail Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToSpecDetailCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) componentCompositionCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ComponentComposition {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for _, componentCompositionData := range *subfuncSDC.Message.ComponentComposition {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": componentCompositionData, "function": "InspectionLotComponentComposition", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "ComponentComposition Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToComponentCompositionCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) inspectionCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Inspection {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for _, inspectionData := range *subfuncSDC.Message.Inspection {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": inspectionData, "function": "InspectionLotInspection", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Inspection Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToInspectionCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) operationCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Operation {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for _, operationData := range *subfuncSDC.Message.Operation {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": operationData, "function": "InspectionLotOperation", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Operation Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToOperationCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) headerUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.Header {
	header := input.Header
	headerData := dpfm_api_processing_formatter.ConvertToHeaderUpdates(header)

	sessionID := input.RuntimeSessionID
	if headerIsUpdate(headerData) {
		res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": headerData, "function": "InspectionLotHeader", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			*errs = append(*errs, err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Header Data cannot update"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToHeaderUpdates(header)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) SpecGeneralUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.SpecGeneral {
	req := make([]dpfm_api_processing_formatter.SpecGeneralUpdates, 0)
	sessionID := input.RuntimeSessionID

	header := input.Header
	for _, specGeneral := range header.SpecGeneral {
		specGeneralData := *dpfm_api_processing_formatter.ConvertToSpecGeneralUpdates(specGeneral)

		if specGeneralIsUpdate(&specGeneralData) {
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": specGeneralData, "function": "InspectionLotSpecGeneral", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				*errs = append(*errs, err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "SpecGeneral Data cannot update"
				return nil
			}
		}
		req = append(req, specGeneralData)
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToSpecGeneralUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) specDetailUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.SpecDetail {
	req := make([]dpfm_api_processing_formatter.SpecDetailUpdates, 0)
	sessionID := input.RuntimeSessionID

	header := input.Header
	for _, specDetail := range header.SpecDetail {
		specDetailData := *dpfm_api_processing_formatter.ConvertToSpecDetailUpdates(specDetail)

		if specDetailIsUpdate(&specDetailData) {
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": specDetailData, "function": "InspectionLotSpecDetail", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				*errs = append(*errs, err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "SpecDetail Data cannot update"
				return nil
			}
		}
		req = append(req, specDetailData)
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToSpecDetailUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) componentCompositionUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ComponentComposition {
	req := make([]dpfm_api_processing_formatter.ComponentCompositionUpdates, 0)
	sessionID := input.RuntimeSessionID

	header := input.Header
	for _, componentComposition := range header.ComponentComposition {
		componentCompositionData := *dpfm_api_processing_formatter.ConvertToComponentCompositionUpdates(componentComposition)

		if componentCompositionIsUpdate(&componentCompositionData) {
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": componentCompositionData, "function": "InspectionLotComponentComposition", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				*errs = append(*errs, err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "ComponentComposition Data cannot update"
				return nil
			}
		}
		req = append(req, componentCompositionData)
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToComponentCompositionUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) inspectionUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Inspection {
	req := make([]dpfm_api_processing_formatter.InspectionUpdates, 0)
	sessionID := input.RuntimeSessionID

	header := input.Header
	for _, inspection := range header.Inspection {
		inspectionData := *dpfm_api_processing_formatter.ConvertToInspectionUpdates(inspection)

		if inspectionIsUpdate(&inspectionData) {
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": inspectionData, "function": "InspectionLotInspection", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				*errs = append(*errs, err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "Inspection Data cannot update"
				return nil
			}
		}
		req = append(req, inspectionData)
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToInspectionUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) operationUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Operation {
	req := make([]dpfm_api_processing_formatter.OperationUpdates, 0)
	sessionID := input.RuntimeSessionID

	header := input.Header
	for _, operation := range header.Operation {
		operationData := *dpfm_api_processing_formatter.ConvertToOperationUpdates(operation)

		if operationIsUpdate(&operationData) {
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": operationData, "function": "InspectionLotOperation", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				*errs = append(*errs, err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "Operation Data cannot update"
				return nil
			}
		}
		req = append(req, operationData)
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToOperationUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func headerIsUpdate(header *dpfm_api_processing_formatter.HeaderUpdates) bool {
	inspectionLot := header.InspectionLot

	return !(inspectionLot == 0)
}

func specGeneralIsUpdate(specGeneral *dpfm_api_processing_formatter.SpecGeneralUpdates) bool {
	inspectionLot := specGeneral.InspectionLot
	heatNumber := specGeneral.HeatNumber

	return !(inspectionLot == 0 || heatNumber == "")
}

func specDetailIsUpdate(specDetail *dpfm_api_processing_formatter.SpecDetailUpdates) bool {
	inspectionLot := specDetail.InspectionLot
	specType := specDetail.SpecType

	return !(inspectionLot == 0 || specType == "")
}

func componentCompositionIsUpdate(componentComposition *dpfm_api_processing_formatter.ComponentCompositionUpdates) bool {
	inspectionLot := componentComposition.InspectionLot
	componentCompositionType := componentComposition.ComponentCompositionType

	return !(inspectionLot == 0 || componentCompositionType == "")
}

func inspectionIsUpdate(inspection *dpfm_api_processing_formatter.InspectionUpdates) bool {
	inspectionLot := inspection.InspectionLot
	//	inspection := inspection.Inspection
	inspectionLotInspectionText := inspection.InspectionLotInspectionText

	return !(inspectionLot == 0 || inspectionLotInspectionText == 0)
}

func OperationIsUpdate(operation *dpfm_api_processing_formatter.OperationUpdates) bool {
	inspectionLot := operation.InspectionLot
	operations := operation.Operations
	operationsItem := operation.OperationsItem
	OperationID := operation.OperationID

	return !(inspectionLot == 0 || operations == 0 || operationsItem == 0 || OperationID == 0)
}
