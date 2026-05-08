package seed

import (
	"log"

	"github.com/adrianus123/project-management/config"
	"github.com/adrianus123/project-management/model"
	"github.com/adrianus123/project-management/util"
	"github.com/google/uuid"
)

func SeedAdmin() {
	password, _ := util.HashPassword("admin123")

	admin := model.User{
		Name:     "Admin",
		Email:    "admin@example.com",
		Password: password,
		Role:     "admin",
		PublicID: uuid.New(),
	}

	err := config.DB.FirstOrCreate(&admin, model.User{Email: admin.Email}).Error
	if err != nil {
		log.Fatalf("Failed to seed admin: %v", err)
	}

	log.Println("Admin seeded successfully")
}
