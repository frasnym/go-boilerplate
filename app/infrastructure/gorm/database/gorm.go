package gorm

import (
	"fmt"
	"time"

	"github.com/frasnym/go-boilerplate/app/entity"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func NewConnectionDB(driverDB string, database string, host string, user string, password string, port int) (*gorm.DB, error) {
	var dialect gorm.Dialector
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	if driverDB == "postgres" {
		dsn := fmt.Sprintf(
			"host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
			host,
			port,
			user,
			database,
			password,
			"disable",
		)

		dialect = postgres.Open(dsn)
	} else if driverDB == "mysql" {
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			user,
			password,
			host,
			port,
			database,
		)

		dialect = mysql.Open(dsn)
	} else if driverDB == "sqlite" {
		dialect = sqlite.Open(database)
	}

	db, err := gorm.Open(dialect, gormConfig)
	if err != nil {
		return nil, err
	}

	// Auto migration
	_ = db.AutoMigrate(&entity.Fanuser{})

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(20)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	//pool time
	tm := time.Minute * time.Duration(20)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(tm)

	return db, nil
}
