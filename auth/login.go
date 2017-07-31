package auth

import (
	"github.com/Greckas/SSU-Golang-252-Chat/loger"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UserAuth struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (ug *UserGorm) ByName(UserName string) (*User, error) {
	res, err := ug.byQuery(ug.DB.Where("username=?", UserName))
	if err != nil{
		// error handling code
		return nil, err
	}
	return res, err
}

func (ug *UserGorm) Login(username, password string) error {
	foundUser, err := ug.ByName(username)
	if err != nil {
		loger.Log.Errorf("No user with that Username")
		return err
	}

	err2 := bcrypt.CompareHashAndPassword(
		[]byte(foundUser.Password),
		[]byte(password))
	if err2 != nil {
		loger.Log.Errorf("Invalid password", err2)
		return err2
	}

	return nil
}

func (ug *UserGorm) byQuery(query *gorm.DB) (*User, error) {
	ret := &User{}
	err := ug.DB.First(ret).Error
	return ret, err
}
