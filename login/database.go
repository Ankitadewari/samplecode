package login

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB(config ConfigDB) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
		config.Host, config.User, config.Password, config.Name, config.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil

}

func CreateUser(username, password, email string) error {
	u := User{Username: username, Password: password, Email: email}
	result := db.Create(&u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetUserByName(name string) (User, error) {
	var user User
	result := db.First(&user, "name = ?", name)
	if result.Error != nil {
		return User{}, result.Error
	}

	return user, nil
}
