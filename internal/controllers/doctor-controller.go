package controllers

import (
	"ipmedpointsistem/internal/models"

	"github.com/sev-2/raiden"
	"github.com/sev-2/raiden/pkg/db"
)

type DoctorRequest struct {
}

type DoctorResponse struct {
}

type DoctorController struct {
	raiden.ControllerBase
	Http    string `path:"/doctor" type:"custom"`
	Model   models.Doctor
	Payload *DoctorRequest
}

func (c *DoctorController) Get(ctx raiden.Context) error {
	var doctors []models.Doctor

	_, err := db.NewQuery(ctx).From(models.Doctor{}).Get()

	if err != nil {
		return ctx.SendError("error")
	}

	return ctx.SendJson(doctors)
}
