package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupDatabase() *gorm.DB {
	dsn := "host=localhost user=your_db_user password=your_db_password dbname=your_db_name port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	migrate()

	return db
}

func migrate() {
	err := db.AutoMigrate(&User{}, &Photo{})
	if err != nil {
		panic("Failed to migrate database")
	}
}
