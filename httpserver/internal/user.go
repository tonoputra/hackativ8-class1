package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name string `json:"name"`
}

type UserList struct {
	List []*User
}

type MemberInterface interface {
	GetMemberRoute(ctx echo.Context) error
	MemberRegistrationRoute(ctx echo.Context) error
	GetMember() []*User
	MemberRegister(u *User) string
}

func CreateUser(u []*User) MemberInterface {
	return &UserList{}
}

func (m *UserList) GetMember() (list []*User) {
	return m.List
}

func (m *UserList) MemberRegister(userdata *User) (res string) {
	m.List = append(m.List, userdata)
	res = fmt.Sprintf("%s berhasil didaftarkan.", userdata.Name)
	return
}

func (m *UserList) GetMemberRoute(ctx echo.Context) error {
	param := ctx.Param("id")
	list := m.GetMember()
	countlist := len(list)
	if param != "" {
		idInt, _ := strconv.Atoi(param)
		if countlist > 0 && idInt <= countlist {
			data := list[idInt]
			return ctx.JSON(http.StatusOK, data)
		} else {
			return ctx.JSON(http.StatusOK, "member tidak ditemukan")
		}
	} else {
		return ctx.JSON(http.StatusOK, list)
	}
}

func (m *UserList) MemberRegistrationRoute(ctx echo.Context) error {
	var getuser User
	if data := json.NewDecoder(ctx.Request().Body).Decode(&getuser); data != nil {
		return ctx.JSON(http.StatusBadRequest, data)
	}
	register := m.MemberRegister(&getuser)
	return ctx.JSON(http.StatusOK, register)
}
