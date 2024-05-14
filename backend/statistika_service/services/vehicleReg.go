package services

import "statistika_service/data"

type VehicleService struct {
	registrations data.VehicleRegistrationRepo
}

func NewVehicleService(registrations data.VehicleRegistrationRepo) (VehicleService, error) {
	return VehicleService{
		registrations: registrations,
	}, nil
}
