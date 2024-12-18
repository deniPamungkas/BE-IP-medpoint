// Code generated by raiden-cli; DO NOT EDIT.
package bootstrap

import (
	"github.com/sev-2/raiden/pkg/resource"
	"ipmedpointsistem/internal/models"
)

func RegisterModels() {
	resource.RegisterModels(
		&models.Doctor{},
		&models.DoctorSchedule{},
		&models.Patients{},
		&models.PaymentMethods{},
		&models.Payments{},
		&models.Reservations{},
		&models.Specialist{},
		&models.Specialization{},
		&models.Test{},
		&models.User{},
		&models.Users{},
	)
}
