package sql

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
  _ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

// Package db has specific primitives for database config & connections.
//
// Usage:
// - 	E.g. dbpkg.NewDb(&c), where c must implement ConfigReader and default use case is to just use Config struct.

const (
	// PostgresConnectionDSNFormat is postgres connection path format for gorm.
	// E.g. host=localhost:3306 dbname=stork sslmode=require user=stork password=stork
	PostgresConnectionDSNFormat = "host=%s dbname=%s sslmode=%s user=%s password=%s"
	// MysqlConnectionDSNFormat is mysql connection path format for gorm.
	// E.g. stork:stork@tcp(localhost:3306)/stork?charset=utf8&parseTime=True&loc=Local
	MysqlConnectionDSNFormat = "%s:%s@%s(%s)/%s?charset=utf8&parseTime=True&loc=Local"
)

// ConfigReader interface has methods to read various db configurations.
type ConfigReader interface {
	GetDialect() string
	GetConnectionPath() string
	GetMaxIdleConnections() int
	GetMaxOpenConnections() int
	GetConnMaxLifetime() time.Duration
}

// Config struct holds db configurations and implements ConfigReader.
type Config struct {
	Dialect               string
	Protocol              string
	URL                   string
	Username              string
	Password              string
	SslMode               string
	Name                  string
	MaxOpenConnections    int
	MaxIdleConnections    int
	ConnectionMaxLifetime time.Duration
	BulkUpdateLimit       int
}

// GetDialect return dialect.
func (c *Config) GetDialect() string {
	return c.Dialect
}

// GetConnectionPath returns connection string to be used by gorm basis dialect.
func (c *Config) GetConnectionPath() string {
	switch c.Dialect {
	case "postgres":
		return fmt.Sprintf(PostgresConnectionDSNFormat, c.URL, c.Name, c.SslMode, c.Username, c.Password)
	default:
		return fmt.Sprintf(MysqlConnectionDSNFormat, c.Username, c.Password, c.Protocol, c.URL, c.Name)
	}
}

// GetMaxOpenConnections returns configurable max open connections.
func (c *Config) GetMaxOpenConnections() int {
	return c.MaxOpenConnections
}

// GetMaxIdleConnections returns configurable max idle connections.
func (c *Config) GetMaxIdleConnections() int {
	return c.MaxIdleConnections
}

// GetConnMaxLifetime returns configurable max lifetime.
func (c *Config) GetConnMaxLifetime() time.Duration {
	return c.ConnectionMaxLifetime
}

// Db is the specific wrapper holding gorm db instance.
type Db struct {
	configReader ConfigReader
	instance     *gorm.DB
	// Driver- if exists will be used to open gorm connection instead of
	// connection URL from configuration.
	driver *sql.DB
}

// NewDb instantiates Db and connects to database.
func NewDb(c ConfigReader, driver *sql.DB) (*Db, error) {
	db := &Db{configReader: c, driver: driver}
	if err := db.connect(); err != nil {
		fmt.Println(err)
		return nil, err
	}
	db.replaceGormCallbacks()
	return db, nil
}

// Instance returns underlying gorm db instance.
func (db *Db) Instance() *gorm.DB {
	return db.instance
}

// Alive executes a select query and checks if connection exists and is alive.
func (db *Db) Alive() error {
	_, err := db.instance.DB().Exec("SELECT 1")
	return err
}

// connect actually opens a gorm connection and configures other connection details.
func (db *Db) connect() error {
	var err error

	if db.driver == nil {
		if db.instance, err = gorm.Open(db.configReader.GetDialect(), db.configReader.GetConnectionPath()); err != nil {
			return err
		}
	} else {
		if db.instance, err = gorm.Open(db.configReader.GetDialect(), db.driver); err != nil {
			return err
		}
	}

	db.instance.DB().SetMaxIdleConns(db.configReader.GetMaxIdleConnections())
	db.instance.DB().SetMaxOpenConns(db.configReader.GetMaxOpenConnections())
	db.instance.DB().SetConnMaxLifetime(db.configReader.GetConnMaxLifetime() * time.Second)
	return nil
}

// Ping checks the connectivity to the database server
func (db *Db) Ping() error {
	// dbProvider.instance.DB().Ping()
	// Ping only checks if the connection is available in the pool
	// If not and connection limit is not reached it'll create one connection
	// If the instance of connection is available in the pool and if it was killed by the database server
	// 	- In this case Ping will return the available instance as it doesn't know the connection was closed by server
	//	- To handle this we execute a query, by doing which the connection will re-established with the server

	var err error

	if _, err = db.instance.DB().Exec("SELECT 1"); err != nil {
		return err
	}

  if _, err = db.instance.DB().Exec("SET AUTOCOMMIT off"); err != nil {
    return err
  }

	return err
}

// replaceGormCallbacks replaces registered gorm's callback to touch timestamps
// on create & update in specific way.
func (db *Db) replaceGormCallbacks() {
	db.instance.Callback().Create().Replace("gorm:update_time_stamp", updateTimestampOnCreate)
	db.instance.Callback().Update().Replace("gorm:update_time_stamp", updateTimestampOnUpdate)
	db.instance.Callback().Delete().Replace("gorm:delete", updateTimestampOnDelete)
}

func updateTimestampOnCreate(scope *gorm.Scope) {
	if f, ok := scope.FieldByName("CreatedAt"); ok {
		if f.IsBlank {
			err := scope.SetColumn("CreatedAt", time.Now().Unix())
			if err != nil {
				log.Fatal(fmt.Sprintf("GORM_CREATEAT_CALLBACK_ERROR %v", err))
			}
		}
	}
	updateTimestampOnUpdate(scope)
}

func updateTimestampOnUpdate(scope *gorm.Scope) {
	if _, ok := scope.FieldByName("UpdatedAt"); ok {
		err := scope.SetColumn("UpdatedAt", time.Now().Unix())
		if err != nil {
			log.Fatal(fmt.Sprintf("GORM_CREATEAT_CALLBACK_ERROR %v", err))
		}
	}
}

func updateTimestampOnDelete(scope *gorm.Scope) {
	// Source: gorm.deleteCallback()
	// Change is using unix timestamp value instead of gorm.NowFunc().
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedAtField, hasDeletedAtField := scope.FieldByName("DeletedAt")

		if !scope.Search.Unscoped && hasDeletedAtField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedAtField.DBName),
				time.Now().Unix(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	// Source: gorm.addExtraSpaceIfExist()
	if str != "" {
		return " " + str
	}
	return ""
}
