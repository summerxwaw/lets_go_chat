package user

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserDB struct {
	gorm.Model
	ID       string
	Username string `gorm:"type:varchar(100);unique"`
}

func (UserDB) TableName() string {
	return "users"
}

type UserRepository struct {
	db *gorm.DB
}

var UsersRepository, RepositoryErr = NewUserRepository()

func NewUserRepository() (*UserRepository, error) {
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_ROOT_PASSWORD")
	hostname := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, hostname, port, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect to database")
		fmt.Println(err)

		return nil, err
	} else {
		fmt.Println("Connected to database")
	}

	if err := db.AutoMigrate(&UserDB{}); err != nil {
		log.Fatalf("Failed to migrate UserDB schema: %v\n", err)
	}

	return &UserRepository{db: db}, nil
}
