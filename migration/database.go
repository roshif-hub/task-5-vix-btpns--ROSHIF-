package migrations

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Username  string    `json:"username" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Photo struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func SetupDatabase() *gorm.DB {
	dsn := "host=localhost user=your_db_user password=your_db_password dbname=your_db_name port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	Migrate()

	return db
}

func Migrate() {
	err := db.AutoMigrate(&User{}, &Photo{})
	if err != nil {
		panic("Failed to migrate database")
	}

	// Anda dapat menambahkan lebih banyak migrasi atau perubahan di sini
}

func main() {
	db = SetupDatabase()
}
