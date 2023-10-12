package requests

type SpecGeneral struct {
	InspectionLot	    int	    `json:"InspectionLot"`
    HeatNumber	        string	`json:"HeatNumber"`
    CreationDate	    string	`json:"CreationDate"`
    LastChangeDate	    string	`json:"LastChangeDate"`
    IsCancelled	        *bool	`json:"IsCancelled"`
    IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
