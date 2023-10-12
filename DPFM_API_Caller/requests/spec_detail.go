package requests

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
