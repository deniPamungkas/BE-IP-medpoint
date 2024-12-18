package controllers

import (
	"ipmedpointsistem/internal/models"

	"github.com/google/uuid"
	"github.com/sev-2/raiden"
	"github.com/sev-2/raiden/pkg/db"
)

type DoctorRequest struct {
	User_Id       uuid.UUID `json:"user_id"`
	Name          string    `json:"name"`
	Gender        string    `json:"gender"`
	Specialist_Id uuid.UUID `json:"specialist_id"`
}

type DeleteDoctorRequest struct {
	Id string `path:"id"`
}

type AddDoctorController struct {
}

type DoctorResponse struct {
	Success bool `json:"success"`
	Data    any  `json:"data"`
	Message string
}

// type DeleteDoctorController struct {
// 	raiden.ControllerBase
// 	Http    string `path:"/doctor/{id}" type:"custom"`
// 	Model   models.Doctor
// 	Payload *DeleteDoctorRequest
// }

type DoctorController struct {
	raiden.ControllerBase
	Http    string `path:"/doctor" type:"custom"`
	Model   models.Doctor
	Payload *DoctorRequest
}

func (c *DoctorController) Get(ctx raiden.Context) error {

	data, err := db.NewQuery(ctx).From(models.Doctor{}).Get()

	if err != nil {
		return ctx.SendError("error")
	}

	response := DoctorResponse{
		Success: true,
		Data:    data,
	}

	return ctx.SendJson(response)
}
