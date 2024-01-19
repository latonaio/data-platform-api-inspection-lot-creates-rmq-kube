package requests

type Inspection struct {
	InspectionLot	                            int	     `json:"InspectionLot"`
    Inspection	                                int	     `json:"Inspection"`
	InspectionDate				                string   `json:"InspectionDate"`
    InspectionType                            	string	 `json:"InspectionType"`
    InspectionTypeValueUnit	                    *string	 `json:"InspectionTypeValueUnit"`
    InspectionTypePlannedValue	                *float32 `json:"InspectionTypePlannedValue"`
    InspectionTypeCertificateType	            *string	 `json:"InspectionTypeCertificateType"`
    InspectionTypeCertificateValueInText	    *string	 `json:"InspectionTypeCertificateValueInText"`
    InspectionTypeCertificateValueInQuantity	*float32 `json:"InspectionTypeCertificateValueInQuantity"`
    InspectionLotInspectionText	                *string	 `json:"InspectionLotInspectionText"`
	CreationDate            					string   `json:"CreationDate"`
	CreationTime            					string   `json:"CreationTime"`
	LastChangeDate          					string   `json:"LastChangeDate"`
	LastChangeTime          					string   `json:"LastChangeTime"`
	IsReleased              					*bool    `json:"IsReleased"`
	IsLocked                					*bool    `json:"IsLocked"`
	IsCancelled             					*bool    `json:"IsCancelled"`
	IsMarkedForDeletion     					*bool    `json:"IsMarkedForDeletion"`
}
