package controllers

import (
	"ipmedpointsistem/internal/models"

	"github.com/sev-2/raiden"
	"github.com/sev-2/raiden/pkg/db"
)

type ScheduleRequest struct {
	// Id            uuid.UUID `json:"id"`
	// User_Id       uuid.UUID `json:"user_id"`
	// Name          string    `json:"name"`
	// Gender        string    `json:"gender"`
	// Specialist_Id uuid.UUID `json:"specialist_id"`
}

// type DeleteDoctorRequest struct {
// 	Id string `path:"id"`
// }

// type AddDoctorController struct {
// }

// type DoctorResponse struct {
// 	Success bool `json:"success"`
// 	Data    any  `json:"data"`
// 	Message string
// }

// type DeleteDoctorController struct {
// 	raiden.ControllerBase
// 	Http    string `path:"/doctor/{id}" type:"custom"`
// 	Model   models.Doctor
// 	Payload *DeleteDoctorRequest
// }

// type GetSingleDoctorController struct {
// 	raiden.ControllerBase
// 	Http    string `path:"/doctor/{id}" type:"custom"`
// 	Model   models.Doctor
// 	Payload *DoctorRequest
// }

type ScheduleController struct {
	raiden.ControllerBase
	Http    string `path:"/schedule" type:"custom"`
	Model   models.DoctorSchedule
	Payload *DoctorRequest
}

func (c *ScheduleController) Get(ctx raiden.Context) error {

	data, err := db.NewQuery(ctx).From(models.DoctorSchedule{}).Get()

	if err != nil {
		return ctx.SendError("error")
	}

	response := DoctorResponse{
		Success: true,
		Data:    string(data),
	}

	return ctx.SendJson(response)
}

// func (c *ScheduleController) Post(ctx raiden.Context) error {

// 	payload := models.Doctor{UserId: c.Payload.User_Id, Name: c.Payload.Name, Gender: c.Payload.Gender, SpecialistId: c.Payload.Specialist_Id}

// 	data, err := db.NewQuery(ctx).From(models.Doctor{}).Insert(payload)

// 	if err != nil {
// 		return ctx.SendError("error")
// 	}

// 	response := DoctorResponse{
// 		Success: true,
// 		Data:    string(data),
// 	}

// 	return ctx.SendJson(response)
// }

// func (c *ScheduleController) Patch(ctx raiden.Context) error {

// 	payload := models.Doctor{Name: "deni"}

// 	data, err := db.NewQuery(ctx).From(models.Doctor{}).Eq("id", c.Payload.Id).Update(payload)

// 	if err != nil {
// 		return ctx.SendError("error")
// 	}

// 	response := DoctorResponse{
// 		Success: true,
// 		Data:    string(data),
// 	}

// 	return ctx.SendJson(response)
// }

// func (c *DeleteDoctorController) Delete(ctx raiden.Context) error {

// 	data, err := db.NewQuery(ctx).From(models.Doctor{}).Eq("id", c.Payload.Id).Delete()

// 	if err != nil {
// 		return ctx.SendError("error")
// 	}

// 	response := DoctorResponse{
// 		Success: true,
// 		Data:    string(data),
// 		Message: string(c.Payload.Id),
// 	}

// 	return ctx.SendJson(response)
// }

// func (c *GetSingleDoctorController) Get(ctx raiden.Context) error {
// 	data, err := db.NewQuery(ctx).From(models.Doctor{}).Eq("id", c.Payload.Id).Single()

// 	if err != nil {
// 		return ctx.SendError("error")
// 	}

// 	response := DoctorResponse{
// 		Success: true,
// 		Data:    string(data),
// 	}

// 	return ctx.SendJson(response)
// }
