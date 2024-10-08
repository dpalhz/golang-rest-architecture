package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"simulation/internal/entity"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Gorm represents a struct that interacts with a database using GORM.
type Gorm interface {
	// Health returns a map of health status information.
	// The keys and values in the map are service-specific.
	GetDB() *gorm.DB
	Health() map[string]string

	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error
}

type gormInstance struct {
	db *gorm.DB
}

var (
	database   = os.Getenv("DB_DATABASE")
	password   = os.Getenv("DB_PASSWORD")
	username   = os.Getenv("DB_USERNAME")
	port       = os.Getenv("DB_PORT")
	host       = os.Getenv("DB_HOST")
	schema     = os.Getenv("DB_SCHEMA")
	dbInstance *gormInstance
)


func NewGorm() Gorm {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable search_path=%s",
		host, username, password, database, port, schema)

	// GORM with PostgreSQL driver
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Adjust log level as needed
	})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	MigrateEntities(db)

	dbInstance = &gormInstance{
		db: db,
	}

	return dbInstance
}

// Health checks the health of the database connection by pinging the database.
// It returns a map with keys indicating various health statistics.
func (g *gormInstance) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	// Ping the database
	sqlDB, err := g.db.DB()
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("failed to get database object: %v", err)
		log.Fatalf("failed to get database object: %v", err)
		return stats
	}

	err = sqlDB.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		log.Fatalf("db down: %v", err)
		return stats
	}

	// Database is up, add more statistics
	stats["status"] = "up"
	stats["message"] = "It's healthy"

	// Get database stats (like open connections, in use, idle, etc.)
	dbStats := sqlDB.Stats()
	stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
	stats["in_use"] = strconv.Itoa(dbStats.InUse)
	stats["idle"] = strconv.Itoa(dbStats.Idle)
	stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
	stats["wait_duration"] = dbStats.WaitDuration.String()
	stats["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
	stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

	// Evaluate stats to provide a health message
	if dbStats.OpenConnections > 40 { // Assuming 50 is the max for this example
		stats["message"] = "The database is experiencing heavy load."
	}

	if dbStats.WaitCount > 1000 {
		stats["message"] = "The database has a high number of wait events, indicating potential bottlenecks."
	}

	if dbStats.MaxIdleClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many idle connections are being closed, consider revising the connection pool settings."
	}

	if dbStats.MaxLifetimeClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."
	}

	return stats
}

// Close closes the database connection.
// It logs a message indicating the disconnection from the specific database.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
func (g *gormInstance) Close() error {
	sqlDB, err := g.db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database object: %v", err)
	}
	log.Printf("Disconnected from database: %s", database)
	return sqlDB.Close()
}

// GetDB returns the underlying *gorm.DB object.
func (g *gormInstance) GetDB() *gorm.DB {
	return g.db
}



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