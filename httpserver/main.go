package main

import (
	"projectone/httpserver/internal"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	var user []*internal.User
	serviceUser := internal.CreateUser(user)

	e.GET("user/list", serviceUser.GetMemberRoute)
	e.GET("user/:id", serviceUser.GetMemberRoute)
	e.POST("user/register", serviceUser.MemberRegistrationRoute)

	e.Logger.Fatal(e.Start(":8888"))
}