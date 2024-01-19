package requests

type SpecDetail struct {
	InspectionLot	        int	    `json:"InspectionLot"`
    SpecType	            string	`json:"SpecType"`
    UpperLimitValue	        float32	`json:"UpperLimitValue"`
    LowerLimitValue	        float32	`json:"LowerLimitValue"`
    StandardValue	        float32	`json:"StandardValue"`
    SpecTypeUnit	        string	`json:"SpecTypeUnit"`
    Formula			        *string	`json:"Formula"`
	CreationDate            string   `json:"CreationDate"`
	CreationTime            string   `json:"CreationTime"`
	LastChangeDate          string   `json:"LastChangeDate"`
	LastChangeTime          string   `json:"LastChangeTime"`
	IsReleased              *bool    `json:"IsReleased"`
	IsLocked                *bool    `json:"IsLocked"`
	IsCancelled             *bool    `json:"IsCancelled"`
	IsMarkedForDeletion     *bool    `json:"IsMarkedForDeletion"`
}
