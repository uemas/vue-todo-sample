package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	var (
		host = os.Getenv("POSTGRES_HOSTNAME")
		user = os.Getenv("POSTGRES_USER")
		pass = os.Getenv("POSTGRES_PASSWORD")
		name = os.Getenv("POSTGRES_DB")
		port = os.Getenv("POSTGRES_POART")
		dsn  = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
			host,
			user,
			pass,
			name,
			port)
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableAutomaticPing: true,
	})
	if err != nil {
		panic(err)
	}
	return db
}
