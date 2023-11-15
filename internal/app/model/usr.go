package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Pass     string `json:"-"`
	Text     string `json:"name"`
	Password string `json:"password,omitempty"`
	Status   int    `json:"-"`
}

func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}
		u.Pass = enc
	}
	return nil
}

func (u *User) Sanitaze() {
	u.Password = ""
}

func (u *User) ComparePassord(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Pass), []byte(password)) == nil
}
func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Login, validation.Required),
		validation.Field(&u.Password, validation.By(requiredIf(u.Pass == "")), validation.Length(6, 100)),
	)
}
