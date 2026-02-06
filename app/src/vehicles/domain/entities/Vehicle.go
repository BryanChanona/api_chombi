package entities

import "github.com/go-playground/validator/v10"

type Vehicle struct {
	Id           int
	LicencePlate string
	Driver       string
	UnitNumber   int
	Shift        string
	IsWorking    bool
}

type VehicleResponse struct {
	Id           int    `json:"id"`
	LicencePlate string `json:"licencePlate"`
	Driver       string `json:"driver"`
	UnitNumber   int    `json:"unitNumber"`
	Shift        string `json:"shift"`
	IsWorking    bool   `json:"isWorking"`
}

func ValidateVehicle(vehicle Vehicle) error {
	validate := validator.New()
	return validate.Struct(vehicle)
}