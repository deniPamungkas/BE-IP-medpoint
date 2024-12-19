package controllers

import (
	"fmt"

	"github.com/sev-2/raiden"
	"github.com/supabase-community/gotrue-go/types"
	"github.com/supabase-community/supabase-go"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Success bool `json:"success"`
	Data    any  `json:"data"`
	Message string
}

type SignUpController struct {
	raiden.ControllerBase
	Http    string `path:"/auth/v1/sign-up" type:"custom"`
	Payload *AuthRequest
}

type LoginController struct {
	raiden.ControllerBase
	Http    string `path:"/auth/v1/login" type:"custom"`
	Payload *AuthRequest
}

type LogoutController struct {
	raiden.ControllerBase
	Http    string `path:"/auth/v1/logout" type:"custom"`
	Payload *AuthRequest
}

func (c *SignUpController) Post(ctx raiden.Context) error {
	client, err := supabase.NewClient(API_URL, API_KEY, &supabase.ClientOptions{})
	if err != nil {
		fmt.Println("cannot initalize client", err)
		return ctx.SendError(err.Error())
	}
	signUpRequest := types.SignupRequest{
		Email:    c.Payload.Email,
		Password: c.Payload.Password,
	}

	user, errorSignUp := client.Auth.Signup(signUpRequest)

	if errorSignUp != nil {
		return ctx.SendError(errorSignUp.Error())
	}

	response := AuthResponse{
		Success: true,
		Data:    user,
		Message: "success to Sign Up",
	}

	return ctx.SendJson(response)

}

var API_URL string = "https://gamitsljjtdiitmhxzuz.supabase.co"
var API_KEY string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImdhbWl0c2xqanRkaWl0bWh4enV6Iiwicm9sZSI6ImFub24iLCJpYXQiOjE3MzM5MzIzODIsImV4cCI6MjA0OTUwODM4Mn0.fUHxrAQHmJ1UPBowZXGJq8uAIztvh9qJeQ_gpt81nek"

func (c *LoginController) Post(ctx raiden.Context) error {

	client, err := supabase.NewClient(API_URL, API_KEY, &supabase.ClientOptions{})
	if err != nil {
		fmt.Println("cannot initalize client", err)
		return ctx.SendError(err.Error())
	}

	user, errorLogin := client.Auth.SignInWithEmailPassword(c.Payload.Email, c.Payload.Password)

	if errorLogin != nil {
		return ctx.SendError(errorLogin.Error())
	}

	response := AuthResponse{
		Success: true,
		Data:    user,
		Message: "success to login",
	}

	return ctx.SendJson(response)

}

func (c *LogoutController) Post(ctx raiden.Context) error {

	client, err := supabase.NewClient(API_URL, API_KEY, &supabase.ClientOptions{})
	if err != nil {
		fmt.Println("cannot initalize client", err)
		return ctx.SendError(err.Error())
	}

	user := client.Auth.Logout()
	response := AuthResponse{
		Success: true,
		Data:    user,
		Message: "success to logout",
	}

	return ctx.SendJson(response)

}
