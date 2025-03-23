package db

import (
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"services-api/services-api/models"
)

// DB is the global database connection
var DB *gorm.DB

// InitDB initializes the SQLite database and seeds it with sample data
func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("services.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	DB.AutoMigrate(&models.Service{}, &models.Version{})

	// Seed the database with sample data
	seedData()
}

// seedData adds sample services and versions if the database is empty
func seedData() {
	var count int64
	DB.Model(&models.Service{}).Count(&count)
	if count > 0 {
		return // Skip seeding if data exists
	}

	// Sample
	service1 := models.Service{
		ID:          uuid.New().String(),
		Name:        "Service One",
		Description: "This is service one",
	}
	service2 := models.Service{
		ID:          uuid.New().String(),
		Name:        "Service Two",
		Description: "This is service two",
	}
	service3 := models.Service{
		ID:          uuid.New().String(),
		Name:        "Service Three",
		Description: "This is service three",
	}
	service4 := models.Service{
		ID:          uuid.New().String(),
		Name:        "Service Four",
		Description: "This is service four",
	}
	service5 := models.Service{
		ID:          uuid.New().String(),
		Name:        "Service Five",
		Description: "This is service five",
	}

	DB.Create(&service1)
	DB.Create(&service2)
	DB.Create(&service3)
	DB.Create(&service4)
	DB.Create(&service5)

	// Sample versions for service1
	DB.Create(&models.Version{
		ID:          uuid.New().String(),
		ServiceID:   service1.ID,
		Version:     "1.0.0",
		ReleaseDate: "2023-01-01",
	})
	DB.Create(&models.Version{
		ID:          uuid.New().String(),
		ServiceID:   service1.ID,
		Version:     "1.1.0",
		ReleaseDate: "2023-02-01",
	})

	// Sample version for service2
	DB.Create(&models.Version{
		ID:          uuid.New().String(),
		ServiceID:   service2.ID,
		Version:     "1.0.0",
		ReleaseDate: "2023-03-01",
	})

	// Sample version for service3
	DB.Create(&models.Version{
		ID:          uuid.New().String(),
		ServiceID:   service3.ID,
		Version:     "1.0.0",
		ReleaseDate: "2023-04-01",
	})

	// Sample version for service4
	DB.Create(&models.Version{
		ID:          uuid.New().String(),
		ServiceID:   service4.ID,
		Version:     "1.0.0",
		ReleaseDate: "2023-05-01",
	})

	// Sample version for service5
	DB.Create(&models.Version{
		ID:          uuid.New().String(),
		ServiceID:   service5.ID,
		Version:     "1.0.0",
		ReleaseDate: "2023-05-01",
	})
}
