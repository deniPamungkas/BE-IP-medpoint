// Code generated by raiden-cli; DO NOT EDIT.
package bootstrap

import (
	"github.com/sev-2/raiden/pkg/resource"
	"ipmedpointsistem/internal/models"
)

func RegisterModels() {
	resource.RegisterModels(
		&models.DoctorSchedule{},
		&models.Doctors{},
		&models.PublicUsers{},
		&models.Specialization{},
	)
}
