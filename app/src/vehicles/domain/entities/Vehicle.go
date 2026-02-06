package entities


type Vehicle struct {
	Id           int `json:"id"`
	LicencePlate string `json:"licencePlate"`
	Driver       string `json:"driver"`
	UnitNumber   int `json:"unitNumber"`
	Shift        string `json:"shift"`
	IsWorking    bool `json:"isWorking"`
	UserId       int `json:"userId"`
}

type VehicleResponse struct {
	Id           int    `json:"id"`
	LicencePlate string `json:"licencePlate"`
	Driver       string `json:"driver"`
	UnitNumber   int    `json:"unitNumber"`
	Shift        string `json:"shift"`
	IsWorking    bool   `json:"isWorking"`
}

