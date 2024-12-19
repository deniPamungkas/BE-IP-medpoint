// Code generated by raiden-cli; DO NOT EDIT.
package bootstrap

import (
	"github.com/sev-2/raiden"
	"ipmedpointsistem/internal/controllers"
	"ipmedpointsistem/internal/models"
	"github.com/valyala/fasthttp"
)

func RegisterRoute(server *raiden.Server) {
	server.RegisterRoute([]*raiden.Route{
		{
			Type:       raiden.RouteTypeCustom,
			Path:       "/auth/v1/sign-up",
			Methods:    []string{fasthttp.MethodPost},
			Controller: &controllers.SignUpController{},
		},
		{
			Type:       raiden.RouteTypeCustom,
			Path:       "/auth/v1/login",
			Methods:    []string{fasthttp.MethodPost},
			Controller: &controllers.LoginController{},
		},
		{
			Type:       raiden.RouteTypeCustom,
			Path:       "/auth/v1/logout",
			Methods:    []string{fasthttp.MethodPost},
			Controller: &controllers.LogoutController{},
		},
		{
			Type:       raiden.RouteTypeCustom,
			Path:       "/doctor/{id}",
			Methods:    []string{fasthttp.MethodDelete},
			Controller: &controllers.DeleteDoctorController{},
			Model:      models.Doctor{},
		},
		{
			Type:       raiden.RouteTypeCustom,
			Path:       "/doctor/{id}",
			Methods:    []string{fasthttp.MethodGet},
			Controller: &controllers.GetSingleDoctorController{},
			Model:      models.Doctor{},
		},
		{
			Type:       raiden.RouteTypeCustom,
			Path:       "/doctor",
			Methods:    []string{fasthttp.MethodGet, fasthttp.MethodPost, fasthttp.MethodPatch},
			Controller: &controllers.DoctorController{},
			Model:      models.Doctor{},
		},
		{
			Type:       raiden.RouteTypeCustom,
			Path:       "/schedule",
			Methods:    []string{fasthttp.MethodGet},
			Controller: &controllers.ScheduleController{},
			Model:      models.DoctorSchedule{},
		},
		{
			Type:       raiden.RouteTypeCustom,
			Path:       "/hello/{name}",
			Methods:    []string{fasthttp.MethodGet},
			Controller: &controllers.HelloWordController{},
		},
	})
}
