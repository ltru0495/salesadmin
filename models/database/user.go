package database

import (
	"admin/config"
	"admin/models"
	"admin/utils"

	"errors"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

const USER_COLLECTION = "users"
const USERNAME = "username"

func InsertUser(user *models.User) error {
	err := user.EncryptPassword()
	if err != nil {
		return err
	}
	user.SetRole()
	user.SetRegistrationDate()
	user.SetLastModificationDate()

	err = insert(USER_COLLECTION, user)
	return err
}

func CompareAdminPassword(password string) error {
	session := utils.GetSession().Copy()
	defer session.Close()
	conf := config.GetDatabaseConfig()
	dbname := conf.Database
	C := session.DB(dbname).C(USER_COLLECTION)
	user := models.User{}

	err := C.Find(bson.M{USERNAME: "admin"}).One(&user)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword(user.HashPassword, []byte(password))
	if err != nil {
		return errors.New("Contraseña incorrecta")
	}
	return nil
}

func UpdatePassword(username, oldPassword, newPassword string) error {
	session := utils.GetSession().Copy()
	defer session.Close()
	conf := config.GetDatabaseConfig()
	dbname := conf.Database
	C := session.DB(dbname).C(USER_COLLECTION)

	user := models.User{}

	err := C.Find(bson.M{USERNAME: username}).One(&user)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword(user.HashPassword, []byte(oldPassword))
	if err != nil {
		return err
	}

	user.Password = newPassword

	err = user.EncryptPassword()
	if err != nil {
		return err
	}

	err = C.Update(bson.M{"username": username}, user)
	if err != nil {
		return err
	}
	return nil
}

func AllShops() []models.Shop {
	session := utils.GetSession().Copy()
	defer session.Close()
	conf := config.GetDatabaseConfig()
	dbname := conf.Database
	C := session.DB(dbname).C(USER_COLLECTION)

	var shops []models.Shop
	var users []models.User
	err := C.Find(nil).All(&users)

	for _, user := range users {
		if user.Role != "admin" {
			shops = append(shops, models.Shop{user.Name})
		}
	}

	if err != nil {
		println(err)
	}

	return shops
}

func GetAllUsers() []models.User {
	session := utils.GetSession().Copy()
	defer session.Close()
	conf := config.GetDatabaseConfig()
	dbname := conf.Database
	C := session.DB(dbname).C(USER_COLLECTION)

	var users []models.User
	err := C.Find(nil).All(&users)
	if err != nil {
		println(err)
	}
	return users
}

func Login(user models.User) (u models.User, err error) {
	session := utils.GetSession().Copy()

	defer session.Close()
	conf := config.GetDatabaseConfig()
	dbname := conf.Database
	C := session.DB(dbname).C(USER_COLLECTION)

	err = C.Find(bson.M{"username": user.Username}).One(&u)
	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword(u.HashPassword, []byte(user.Password))
	if err != nil {
		u = models.User{}
	}
	return
}
