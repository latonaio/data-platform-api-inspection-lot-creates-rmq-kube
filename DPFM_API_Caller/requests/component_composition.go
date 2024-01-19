package requests

type ComponentComposition struct {
	InspectionLot								int		`json:"InspectionLot"`
	ComponentCompositionType					string	`json:"ComponentCompositionType"`
	ComponentCompositionUpperLimitInPercent		float32	`json:"ComponentCompositionUpperLimitInPercent"`
	ComponentCompositionLowerLimitInPercent		float32	`json:"ComponentCompositionLowerLimitInPercent"`
	ComponentCompositionStandardValueInPercent	float32	`json:"ComponentCompositionStandardValueInPercent"`
	CreationDate            					string   `json:"CreationDate"`
	CreationTime            					string   `json:"CreationTime"`
	LastChangeDate          					string   `json:"LastChangeDate"`
	LastChangeTime          					string   `json:"LastChangeTime"`
	IsReleased              					*bool    `json:"IsReleased"`
	IsLocked                					*bool    `json:"IsLocked"`
	IsCancelled             					*bool    `json:"IsCancelled"`
	IsMarkedForDeletion     					*bool    `json:"IsMarkedForDeletion"`
}
