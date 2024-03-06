package types

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int
	Username  string
	Email     string
	FirstName string
	LastName  string
	Password  string
}

func NewUser(username, email, firstname, lastname, password string) *User {
	return &User{
		Username:  username,
		Email:     email,
		FirstName: firstname,
		LastName:  lastname,
		Password:  password,
	}
}

func UserFromRequest(c echo.Context) *User {
	return &User{
		Username:  c.FormValue("username"),
		Email:     c.FormValue("email"),
		FirstName: c.FormValue("firstname"),
		LastName:  c.FormValue("firstname"),
		Password:  c.FormValue("password"),
	}
}

func (user *User) IsValid() bool {
	return user.Username != "" && user.Email != "" && user.FirstName != "" && user.LastName != "" && user.Password != ""
}

func (user *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) Authenticate(password string) bool {
	isAuthed := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if isAuthed != nil {
		return false
	}
	return true
}
