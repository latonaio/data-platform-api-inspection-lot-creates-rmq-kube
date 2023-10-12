package requests

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
