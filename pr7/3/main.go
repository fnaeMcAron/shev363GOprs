package main

import (
	"crypto/sha256"
	"fmt"
)

type User struct {
	Username string
	Email    string
	Password []byte
}

func (u *User) SetPassword(password string) {
	hash := sha256.New()
	hash.Write([]byte(password))
	u.Password = hash.Sum(nil)
}

func (u *User) VerifyPassword(password string) bool {
	hash := sha256.New()
	hash.Write([]byte(password))
	enteredPasswordHash := hash.Sum(nil)
	return string(u.Password) == string(enteredPasswordHash)
}

func main() {
	user := &User{
		Username: "JDH",
		Email:    "JDH@lolo.ru",
	}

	user.SetPassword("12345678") //мощный пароль
	fmt.Println("Проверяем первый пароль", user.VerifyPassword("12345678"))
	fmt.Println("Проверяем второй пароль", user.VerifyPassword("11111111")) // еще более мощный пароль

}
