package models

import (
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type User struct {
	Id           bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Username     string        `json:"username" bson:"username"`
	Name         string        `json:"name" bson:"name"`
	Password     string        `json:"password,omitempty" bson:"password,omitempty"`
	HashPassword []byte        `json:"hashpassword,omitempty" bson:"hashpassword,omitempty"`
	Role         string        `json:"role" bson:"role"`
	RegDate      time.Time     `json:"regdate" bson:"registration_date"`
	LastMod      time.Time     `json:"lastmod" bson:"last_modification"`

	Context map[string]interface{} `json:"context,omitempty" bson:"context,omitempty"`
}

type Shop struct {
	Name string `json:"name" bson:"name"`
}

func (u User) String() string {
	return u.Username + "\n" + u.Role + "\n"
}

func (user *User) IsAdmin() bool {
	if user.Role == "admin" {
		return true
	}
	return false

}

func (user *User) NewObjectId() {
	user.Id = bson.NewObjectId()
}

func (user *User) EncryptPassword() error {
	hpass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.HashPassword = hpass
	user.Password = ""
	return nil
}

func (user *User) SetRole() {
	if user.Role != "admin" {
		user.Role = "usr"
	}
}
func (user *User) SetPassword(passwd string) {
	user.Password = passwd
}

func (user *User) SetRegistrationDate() {
	user.RegDate = time.Now()
}

func (user *User) SetLastModificationDate() {
	user.LastMod = time.Now()
}
