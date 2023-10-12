package requests

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
