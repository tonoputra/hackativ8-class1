package main

import (
	"fmt"
	"sync"
	"time"
)

type User struct {
	Name string
}

type ListUser struct {
	ListName []*User
}

type UserInterface interface {
	GetUser() []*User
	RegisterUser(u *User)
}

func NewUser() UserInterface {
	return &ListUser{}
}

func (u *ListUser) GetUser() (userlist []*User) {
	return u.ListName
}

func (u *ListUser) RegisterUser(user *User) {
	u.ListName = append(u.ListName, user)
	fmt.Println(user.Name, " berhasil didaftarkan")
	sec := 5
	time.Sleep(time.Duration(sec)*time.Second)
}

var wg sync.WaitGroup

func main() {
	wg.Add(6)
	user := NewUser()
	go func() {
		user.RegisterUser(&User{Name: "tono"})
		wg.Done()
	}()
	go func() {
		user.RegisterUser(&User{Name: "ivan"})
		wg.Done()
	}()
	go func() {
		user.RegisterUser(&User{Name: "adit"})
		wg.Done()
	}()
	go func() {
		user.RegisterUser(&User{Name: "fiqri"})
		wg.Done()
	}()
	go func() {
		user.RegisterUser(&User{Name: "rijal"})
		wg.Done()
	}()
	go func() {
		user.RegisterUser(&User{Name: "clara"})
		wg.Done()
	}()
	wg.Wait()

	fmt.Println("------- hasil get user -----------")

	for _, v := range user.GetUser() {
		fmt.Println(v.Name)
	}
}