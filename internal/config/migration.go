package config

import (
	"log"
	"simulation/internal/entity"

	"gorm.io/gorm"
)

func MigrateEntities(db *gorm.DB) {
    // List of entities to migrate
    entities := []interface{}{
        &entity.User{},
        &entity.Admin{},
        &entity.Blog{},
        &entity.Category{},
        &entity.Tag{},
    }

    // Loop through each entity and migrate it
    for _, entity := range entities {
        err := db.AutoMigrate(entity)
        if err != nil {
            log.Fatalf("failed to migrate entity %T: %v", entity, err)
        } else {
            log.Printf("Migrated entity %T successfully", entity)
        }
    }
}