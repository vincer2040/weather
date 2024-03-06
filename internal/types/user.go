package types

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID        int
	Username  string
	FirstName string
	LastName  string
	Password  string
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
