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
	GetDB() *gorm.DB
	Health() map[string]string
	Close() error
}

type gormInstance struct {
	db *gorm.DB
}

var (
	// Environment variables for the first database
	dbHost1     = os.Getenv("DB_HOST")
	dbUsername1 = os.Getenv("DB_USERNAME")
	dbPassword1 = os.Getenv("DB_PASSWORD")
	dbPort1     = os.Getenv("DB_PORT")
	dbSchema1   = os.Getenv("DB_SCHEMA")
	dbName1     = os.Getenv("DB_DATABASE")
	dbName2     = os.Getenv("DB_DATABASE2")

)

var (
	dbInstance1 *gormInstance
	dbInstance2 *gormInstance
)

// NewGorm creates and returns two Gorm instances for the databases.
func NewGorm() (Gorm, Gorm) {
	if dbInstance1 == nil {
		dbInstance1 = createGormInstance(
			dbHost1,
			dbUsername1,
			dbPassword1,
			dbName1,
			dbPort1,
			dbSchema1,
		)
		MigrateEntitiesDB1(dbInstance1.GetDB()) 
	}

	if dbInstance2 == nil {
		dbInstance2 = createGormInstance(
			dbHost1,
			dbUsername1,
			dbPassword1,
			dbName2,
			dbPort1,
			dbSchema1,
		)
		MigrateEntitiesDB2(dbInstance2.GetDB()) 
	}

	return dbInstance1, dbInstance2
}

// createGormInstance handles the creation of a Gorm instance.
func createGormInstance(host, username, password, database, port, schema string) *gormInstance {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable search_path=%s",
		host, username, password, database, port, schema)

	fmt.Println(connStr)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	return &gormInstance{db: db}
}


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
func (g *gormInstance) Close() error {
	sqlDB, err := g.db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database object: %v", err)
	}
	log.Printf("Disconnected from database")
	return sqlDB.Close()
}

// GetDB returns the underlying *gorm.DB object.
func (g *gormInstance) GetDB() *gorm.DB {
	return g.db
}

func MigrateEntitiesDB1(db *gorm.DB) {
	entities := []interface{}{
		&entity.User{},
		&entity.Blog{},
		&entity.Admin{},
		&entity.Category{},
		&entity.Tag{},
	}

	migrateEntities(db, entities)
}


func MigrateEntitiesDB2(db *gorm.DB) {
	// List of entities specific to database 2
	entities := []interface{}{
		&entity.Analytic{},
	}

	migrateEntities(db, entities)
}


func migrateEntities(db *gorm.DB, entities []interface{}) {
	for _, entity := range entities {
		err := db.AutoMigrate(entity)
		if err != nil {
			log.Fatalf("failed to migrate entity %T: %v", entity, err)
		} else {
			log.Printf("Migrated entity %T successfully", entity)
		}
	}
}