package dpfm_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type SDC struct {
	ConnectionKey     string          `json:"connection_key"`
	Result            bool            `json:"result"`
	RedisKey          string          `json:"redis_key"`
	Filepath          string          `json:"filepath"`
	APIStatusCode     int             `json:"api_status_code"`
	RuntimeSessionID  string          `json:"runtime_session_id"`
	BusinessPartnerID *int            `json:"business_partner"`
	ServiceLabel      string          `json:"service_label"`
	APIType           string          `json:"api_type"`
	Header            Header          `json:"InspectionLot"`
	APISchema         string          `json:"api_schema"`
	Accepter          []string        `json:"accepter"`
	Deleted           bool            `json:"deleted"`
}

type Header struct {
	InspectionLot                  int     `json:"InspectionLot"`
	InspectionPlan                 *int    `json:"InspectionPlan"`
	InspectionPlantBusinessPartner *int    `json:"InspectionPlantBusinessPartner"`
	InspectionPlant                *string `json:"InspectionPlant"`
	Product                        *string `json:"Product"`
	ProductSpecification           *string `json:"ProductSpecification"`
	InspectionSpecification        *string `json:"InspectionSpecification"`
	ProductionOrder                *int    `json:"ProductionOrder"`
	ProductionOrderItem            *int    `json:"ProductionOrderItem"`
	InspectionLotHeaderText        *string `json:"InspectionLotHeaderText"`
	CreationDate                   *string `json:"CreationDate"`
	LastChangeDate                 *string `json:"LastChangeDate"`
	IsCancelled                    *bool   `json:"IsCancelled"`
	IsMarkedForDeletion            *bool   `json:"IsMarkedForDeletion"`
	SpecGeneral                    []SpecGeneral			`json:"SpecGeneral"`
	SpecDetail                     []SpecDetail				`json:"SpecDetail"`
	ComponentComposition		   []ComponentComposition	`json:"ComponentComposition"`
	Inspection                     []Inspection				`json:"Inspection"`
	Operation                      []Operation				`json:"Operation"`	
}

type SpecGeneral struct {
	InspectionLot	    int	    `json:"InspectionLot"`
    HeatNumber	        string	`json:"HeatNumber"`
    CreationDate	    *string	`json:"CreationDate"`
    LastChangeDate	    *string	`json:"LastChangeDate"`
    IsCancelled	        *bool	`json:"IsCancelled"`
    IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
	SpecDetail          []SpecDetail	`json:"SpecDetail"`
}

type SpecDetail struct {
	InspectionLot	        int	    `json:"InspectionLot"`
    SpecType	            string	`json:"SpecType"`
    UpperLimitValue	        *float32 `json:"UpperLimitValue"`
    LowerLimitValue	        *float32 `json:"LowerLimitValue"`
    StandardValue	        *float32 `json:"StandardValue"`
    SpecTypeUnit	        *string	`json:"SpecTypeUnit"`
    CreationDate	        *string	`json:"CreationDate"`
    LastChangeDate	        *string	`json:"LastChangeDate"`
    IsCancelled	            *bool	`json:"IsCancelled"`
    IsMarkedForDeletion	    *bool	`json:"IsMarkedForDeletion"`
}

type ComponentComposition struct {
	InspectionLot								int		`json:"InspectionLot"`
	ComponentCompositionType					string	`json:"ComponentCompositionType"`
	ComponentCompositionUperLimitInPercent		*float32 `json:"ComponentCompositionUperLimitInPercent"`
	ComponentCompositionLowerLimitInPercent		*float32 `json:"ComponentCompositionLowerLimitInPercent"`
	ComponentCompositionStandardValueInPercent	*float32 `json:"ComponentCompositionStandardValueInPercent"`
	CreationDate								*string	`json:"CreationDate"`
	LastChangeDate								*string	`json:"LastChangeDate"`
	IsCancelled 								*bool	`json:"IsCancelled" `
	IsMarkedForDeletion							*bool	`json:"IsMarkedForDeletion"`
}

type Inspection struct {
	InspectionLot	                            int	        `json:"InspectionLot"`
    Inspection	                                int	        `json:"Inspection"`
    InspectionType                            	string	    `json:"InspectionType"`
    InspectionTypeValueUnit	                    *string	    `json:"InspectionTypeValueUnit"`
    InspectionTypePlannedValue	                *float32	`json:"InspectionTypePlannedValue"`
    InspectionTypeCertificateType	            *string	    `json:"InspectionTypeCertificateType"`
    InspectionTypeCertificateValueInText	    *string	    `json:"InspectionTypeCertificateValueInText"`
    InspectionTypeCertificateValueInQuantity	*float32	`json:"InspectionTypeCertificateValueInQuantity"`
    InspectionLotInspectionText	                *string	    `json:"InspectionLotInspectionText"`
    CreationDate	                            *string	    `json:"CreationDate"`
    LastChangeDate	                            *string	    `json:"LastChangeDate"`
	IsCancelled                                 *bool       `json:"IsCancelled"`
    IsMarkedForDeletion	                        *bool	    `json:"IsMarkedForDeletion"`
}

type Operation struct {
	InspectionLot                            int      `json:"InspectionLot"`
	Operations                               int      `json:"Operations"`
	OperationsItem                           int      `json:"OperationsItem"`
	OperationID                              int      `json:"OperationID"`
	Inspection                               *int     `json:"Inspection"`
	OperationType                            *string  `json:"OperationType"`
	SupplyChainRelationshipID                *int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID        *int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID   *int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipProductionPlantID *int     `json:"SupplyChainRelationshipProductionPlantID"`
	Product                                  *string  `json:"Product"`
	Buyer                                    *int     `json:"Buyer"`
	Seller                                   *int     `json:"Seller"`
	DeliverToParty                           *int     `json:"DeliverToParty"`
	DeliverToPlant                           *string  `json:"DeliverToPlant"`
	DeliverFromParty                         *int     `json:"DeliverFromParty"`
	DeliverFromPlant                         *string  `json:"DeliverFromPlant"`
	InspectionPlantBusinessPartner           *int     `json:"InspectionPlantBusinessPartner"`
	InspectionPlant                          *string  `json:"InspectionPlant"`
	Sequence                                 *int     `json:"Sequence"`
	SequenceText                             *string  `json:"SequenceText"`
	OperationText                            *string  `json:"OperationText"`
	OperationStatus                          *string  `json:"OperationStatus"`
	ResponsiblePlannerGroup                  *string  `json:"ResponsiblePlannerGroup"`
	OperationUnit                            *string  `json:"OperationUnit"`
	StandardLotSizeQuantity                  *float32 `json:"StandardLotSizeQuantity"`
	MinimumLotSizeQuantity                   *float32 `json:"MinimumLotSizeQuantity"`
	MaximumLotSizeQuantity                   *float32 `json:"MaximumLotSizeQuantity"`
	PlainLongText                            *string  `json:"PlainLongText"`
	WorkCenter                               *int     `json:"WorkCenter"`
	CapacityCategoryCode                     *string  `json:"CapacityCategoryCode"`
	OperationCostingRelevancyType            *string  `json:"OperationCostingRelevancyType"`
	OperationSetupType                       *string  `json:"OperationSetupType"`
	OperationSetupGroupCategory              *string  `json:"OperationSetupGroupCategory"`
	OperationSetupGroup                      *string  `json:"OperationSetupGroup"`
	OperationReferenceQuantity               *float32 `json:"OperationReferenceQuantity"`
	MaximumWaitDuration                      *float32 `json:"MaximumWaitDuration"`
	StandardWaitDuration                     *float32 `json:"StandardWaitDuration"`
	MinimumWaitDuration                      *float32 `json:"MinimumWaitDuration"`
	WaitDurationUnit                         *string  `json:"WaitDurationUnit"`
	MaximumQueueDuration                     *float32 `json:"MaximumQueueDuration"`
	StandardQueueDuration                    *float32 `json:"StandardQueueDuration"`
	MinimumQueueDuration                     *float32 `json:"MinimumQueueDuration"`
	QueueDurationUnit                        *string  `json:"QueueDurationUnit"`
	MaximumMoveDuration                      *float32 `json:"MaximumMoveDuration"`
	StandardMoveDuration                     *float32 `json:"StandardMoveDuration"`
	MinimumMoveDuration                      *float32 `json:"MinimumMoveDuration"`
	MoveDurationUnit                         *string  `json:"MoveDurationUnit"`
	StandardDeliveryDuration                 *float32 `json:"StandardDeliveryDuration"`
	StandardDeliveryDurationUnit             *string  `json:"StandardDeliveryDurationUnit"`
	StandardOperationScrapPercent            *float32 `json:"StandardOperationScrapPercent"`
	CostElement                              *string  `json:"CostElement"`
	ValidityStartDate                        *string  `json:"ValidityStartDate"`
	ValidityEndDate                          *string  `json:"ValidityEndDate"`
	CreationDate                             *string  `json:"CreationDate"`
	LastChangeDate                           *string  `json:"LastChangeDate"`
	IsCancelled                              *bool    `json:"IsCancelled"`
	IsMarkedForDeletion                      *bool    `json:"IsMarkedForDeletion"`
}
