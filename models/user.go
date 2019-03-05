package models

import "log"

type User struct {
	ID int `gorm:"primary_key": json:"id"`
	Username string `json:"username"`
	PasswordHash string `json:"password_hash"`
	Email string `json:"email"`
	IsAdmin int `json:"is_admin"`
	Name string `json:"name"`
	Location string `json:"location"`
	Bio string `json:"bio"`
	MemberSince string `json:"member_since"`
	AvatarHash string `json:"avatar_hast"`
}

func CreateUser(user User) bool {
	result := db.Create(&user)
	if result.Error != nil {
		log.Fatal("create user failed", result.Error.Error())
		return false
	}
	return true

}