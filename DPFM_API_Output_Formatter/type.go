package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey       string      `json:"connection_key"`
	Result              bool        `json:"result"`
	RedisKey            string      `json:"redis_key"`
	Filepath            string      `json:"filepath"`
	APIStatusCode       int         `json:"api_status_code"`
	RuntimeSessionID    string      `json:"runtime_session_id"`
	BusinessPartnerID   *int        `json:"business_partner"`
	ServiceLabel        string      `json:"service_label"`
	APIType             string      `json:"api_type"`
	Message             interface{} `json:"message"`
	APISchema           string      `json:"api_schema"`
	Accepter            []string    `json:"accepter"`
	Deleted             bool        `json:"deleted"`
	SQLUpdateResult     *bool       `json:"sql_update_result"`
	SQLUpdateError      string      `json:"sql_update_error"`
	SubfuncResult       *bool       `json:"subfunc_result"`
	SubfuncError        string      `json:"subfunc_error"`
	ExconfResult        *bool       `json:"exconf_result"`
	ExconfError         string      `json:"exconf_error"`
	APIProcessingResult *bool       `json:"api_processing_result"`
	APIProcessingError  string      `json:"api_processing_error"`
}

type Message struct {
	Header 						*[]Header 				`json:"Header"`
	SpecGeneral					*[]SpecGeneral			`json:"SpecGeneral"`
	SpecDetail					*[]SpecDetail			`json:"SpecDetail"`
	ComponentComposition		*[]ComponentComposition	`json:"ComponentComposition"`
	Inspection					*[]Inspection			`json:"Inspection"`
	Operation					*[]Operation			`json:"Operation"`
}

type Header struct {
	InspectionLot                  int     `json:"InspectionLot"`
	InspectionPlan                 int     `json:"InspectionPlan"`
	InspectionPlantBusinessPartner int     `json:"InspectionPlantBusinessPartner"`
	InspectionPlant                string  `json:"InspectionPlant"`
	Product                        string  `json:"Product"`
	ProductSpecification           *string `json:"ProductSpecification"`
	InspectionSpecification        *string `json:"InspectionSpecification"`
	ProductionOrder                *int    `json:"ProductionOrder"`
	ProductionOrderItem            *int    `json:"ProductionOrderItem"`
	InspectionLotHeaderText        *string `json:"InspectionLotHeaderText"`
	CreationDate                   string  `json:"CreationDate"`
	LastChangeDate                 string  `json:"LastChangeDate"`
	IsCancelled                    *bool   `json:"IsCancelled"`
	IsMarkedForDeletion            *bool   `json:"IsMarkedForDeletion"`
}

type SpecGeneral struct {
	InspectionLot	    int	    `json:"InspectionLot"`
    HeatNumber	        string	`json:"HeatNumber"`
    CreationDate	    string	`json:"CreationDate"`
    LastChangeDate	    string	`json:"LastChangeDate"`
    IsCancelled	        *bool	`json:"IsCancelled"`
    IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}

type SpecDetail struct {
	InspectionLot	        int	    `json:"InspectionLot"`
    SpecType	            string	`json:"SpecType"`
    UpperLimitValue	        float32	`json:"UpperLimitValue"`
    LowerLimitValue	        float32	`json:"LowerLimitValue"`
    StandardValue	        float32	`json:"StandardValue"`
    SpecTypeUnit	        string	`json:"SpecTypeUnit"`
    CreationDate	        string	`json:"CreationDate"`
    LastChangeDate	        string	`json:"LastChangeDate"`
    IsCancelled	            *bool	`json:"IsCancelled"`
    IsMarkedForDeletion	    *bool	`json:"IsMarkedForDeletion"`
}

type ComponentComposition struct {
	InspectionLot								int		`json:"InspectionLot"`
	ComponentCompositionType					string	`json:"ComponentCompositionType"`
	ComponentCompositionUperLimitInPercent		float32	`json:"ComponentCompositionUperLimitInPercent"`
	ComponentCompositionLowerLimitInPercent		float32	`json:"ComponentCompositionLowerLimitInPercent"`
	ComponentCompositionStandardValueInPercent	float32	`json:"ComponentCompositionStandardValueInPercent"`
	CreationDate								string	`json:"CreationDate"`
	LastChangeDate								string	`json:"LastChangeDate"`
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
    CreationDate	                            string	    `json:"CreationDate"`
    LastChangeDate	                            string	    `json:"LastChangeDate"`
	IsCancelled                                 *bool       `json:"IsCancelled"`
    IsMarkedForDeletion	                        *bool	    `json:"IsMarkedForDeletion"`
}

type Operation struct {
	InspectionLot                            int      `json:"InspectionLot"`
	Operations                               int      `json:"Operations"`
	OperationsItem                           int      `json:"OperationsItem"`
	OperationID                              int      `json:"OperationID"`
	Inspection                               int      `json:"Inspection"`
	OperationType                            string   `json:"OperationType"`
	SupplyChainRelationshipID                int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID        int      `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID   int      `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipProductionPlantID int      `json:"SupplyChainRelationshipProductionPlantID"`
	Product                                  string   `json:"Product"`
	Buyer                                    int      `json:"Buyer"`
	Seller                                   int      `json:"Seller"`
	DeliverToParty                           int      `json:"DeliverToParty"`
	DeliverToPlant                           string   `json:"DeliverToPlant"`
	DeliverFromParty                         int      `json:"DeliverFromParty"`
	DeliverFromPlant                         string   `json:"DeliverFromPlant"`
	InspectionPlantBusinessPartner           int      `json:"InspectionPlantBusinessPartner"`
	InspectionPlant                          string   `json:"InspectionPlant"`
	Sequence                                 int      `json:"Sequence"`
	SequenceText                             *string  `json:"SequenceText"`
	OperationText                            string   `json:"OperationText"`
	OperationStatus                          *string  `json:"OperationStatus"`
	ResponsiblePlannerGroup                  *string  `json:"ResponsiblePlannerGroup"`
	OperationUnit                            string   `json:"OperationUnit"`
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
	CreationDate                             string   `json:"Creation"`
	LastChangeDate                           string   `json:"LastChange"`
	IsCancelled                              *bool    `json:"IsCancelled"`
	IsMarkedForDeletion                      *bool    `json:"IsMarkedForDeletion"`
}
