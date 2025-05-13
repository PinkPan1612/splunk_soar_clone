package main

import (
	"log"
	"path/filepath"
	"runtime"
	"splunk_soar_clone/config"
	domain "splunk_soar_clone/internal/domain/entity"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func migrateSchema(db *gorm.DB) error {
	return db.AutoMigrate(
		&domain.Role{},  // Phải gọi Role trước
		&domain.User{},  // Gọi User sau vì phụ thuộc Role
		&domain.Token{}, // Token phụ thuộc User
	)
}

func createDefaultRoles(db *gorm.DB) error {
	roles := []domain.Role{
		{
			RoleID:      "1",
			RoleName:    "Administrator",
			Description: "System administrator with full access",
		},
		{
			RoleID:      "2",
			RoleName:    "Regular User",
			Description: "Regular user with limited access",
		},
	}

	for _, role := range roles {
		if err := db.FirstOrCreate(&role, domain.Role{RoleID: role.RoleID}).Error; err != nil {
			return err
		}
	}
	return nil
}

func runMigration(db *gorm.DB) error {
	// Migrate schema
	if err := migrateSchema(db); err != nil {
		return err
	}
	log.Println("✅ Schema migration completed")

	// Create default roles
	if err := createDefaultRoles(db); err != nil {
		return err
	}
	log.Println("✅ Default roles created")

	return nil
}

func main() {
	_, b, _, _ := runtime.Caller(0)
	projectRoot := filepath.Join(filepath.Dir(b), "../..")

	// Load .env from project root
	if err := godotenv.Load(filepath.Join(projectRoot, ".env")); err != nil {
		log.Fatalf("❌ Error loading .env file: %v", err)
	}

	db := config.ConnectionDatabase()

	if err := runMigration(db); err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}

	log.Println("✅ Migration completed successfully")
}
