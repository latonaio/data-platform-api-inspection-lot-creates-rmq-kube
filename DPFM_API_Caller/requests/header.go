package requests

type Header struct {
	InspectionLot                  int     `json:"InspectionLot"`
	InspectionLotDate              string  `json:"InspectionLotDate"`
	InspectionPlan                 int     `json:"InspectionPlan"`
	InspectionPlantBusinessPartner int     `json:"InspectionPlantBusinessPartner"`
	InspectionPlant                string  `json:"InspectionPlant"`
	Product                        string  `json:"Product"`
	ProductSpecification           *string `json:"ProductSpecification"`
	InspectionSpecification        *string `json:"InspectionSpecification"`
	ProductionOrder                *int    `json:"ProductionOrder"`
	ProductionOrderItem            *int    `json:"ProductionOrderItem"`
	InspectionLotHeaderText        *string `json:"InspectionLotHeaderText"`
	ExternalReferenceDocument      *string  `json:"ExternalReferenceDocument"`
	CertificateAuthorityChain      *string  `json:"CertificateAuthorityChain"`
	UsageControlChain        	   *string  `json:"UsageControlChain"`
	CreationDate                   string   `json:"CreationDate"`
	CreationTime                   string   `json:"CreationTime"`
	LastChangeDate                 string   `json:"LastChangeDate"`
	LastChangeTime                 string   `json:"LastChangeTime"`
	IsReleased                     *bool    `json:"IsReleased"`
	IsPartiallyConfirmed           *bool    `json:"IsPartiallyConfirmed"`
	IsConfirmed                    *bool    `json:"IsConfirmed"`
	IsLocked                       *bool    `json:"IsLocked"`
	IsCancelled                    *bool    `json:"IsCancelled"`
	IsMarkedForDeletion            *bool    `json:"IsMarkedForDeletion"`
}
