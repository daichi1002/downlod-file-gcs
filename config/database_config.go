package config

import (
	"downlod-file-gcs/util"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	// DBHost is environment variables key for `DB_HOST`.
	DBHost = "DB_HOST"
	// DBName is environment variables key for `DB_NAME`.
	DBName = "DB_NAME"
	// DBUser is environment variables key for `DB_USER`.
	DBUser = "DB_USER"
	// DBPassword is environment variables key for `DB_PASSWORD`.
	DBPassword = "DB_PASSWORD"
	// DBPort is environment variables key for `DB_PORT`.
	DBPort = "DB_PORT"
	// DBIdleConnection is environment variables key for `DB_IDLE_CONNECTION`.
	DBIdleConnection = "DB_IDLE_CONNECTION"
	// DBMaxConnection is environment variables key for `DB_MAX_CONNECTION`.
	DBMaxConnection = "DB_MAX_CONNECTION"
	// DBConnMaxLifeTime is environment variables key for `DB_CONN_MAX_LIFE_TIME`.
	DBConnMaxLifeTime = "DB_CONN_MAX_LIFE_TIME"
	// DBTimeZone is environment variables key for `DB_TIME_ZONE`
	DBTimeZone = "DB_TIME_ZONE"

	dbConnectRetryWaitTime = 3 * time.Second
)

var (
	databaseViper = viper.New()
	logger        = util.NewLogger()
)

type DatabaseConfig struct {
	DBHost            string
	DBPort            int
	DBName            string
	DBUser            string
	DBPassword        string
	DBIdleConnection  int
	DBMaxConnection   int
	DBConnMaxLifeTime time.Duration
	DBTimeZone        string
}

func init() {
	databaseViper.AutomaticEnv()
	databaseViper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	databaseViper.SetDefault(DBPort, 3306)
	databaseViper.SetDefault(DBName, "dev")
	databaseViper.SetDefault(DBHost, "127.0.0.1")
	databaseViper.SetDefault(DBUser, "root")
	databaseViper.SetDefault(DBIdleConnection, 100)
	databaseViper.SetDefault(DBMaxConnection, 100)
	databaseViper.SetDefault(DBConnMaxLifeTime, 100)
	databaseViper.SetDefault(DBTimeZone, "UTC")
}

// ConnectDatabaseWithGorm is connect database with gorm.
// And returns gorm.DB instance.
func (d DatabaseConfig) ConnectDatabaseWithGorm(retry int) (*gorm.DB, error) {
	if retry <= 0 {
		return nil, errors.New("connect retry over")
	}
	var url string
	if len(d.DBPassword) != 0 {
		url = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", d.DBUser, d.DBPassword, d.DBHost, d.DBPort, d.DBName)
	} else {
		url = fmt.Sprintf("%v@tcp(%v:%v)/%v", d.DBUser, d.DBHost, d.DBPort, d.DBName)
	}
	url += fmt.Sprintf("?charset=utf8&parseTime=True&loc=%s", d.DBTimeZone)

	db, gormErr := gorm.Open(mysql.Open(url), &gorm.Config{})
	if gormErr != nil {
		time.Sleep(dbConnectRetryWaitTime)
		logger.Warnf("failed to connect database, error: %v connect retry...", gormErr)
		return d.ConnectDatabaseWithGorm(retry - 1)
	}

	sqlDB, sqlErr := db.DB()
	if sqlErr != nil {
		logger.Errorf("failed to get sqldb, error: %v", sqlErr)
		return nil, sqlErr
	}
	sqlDB.SetMaxOpenConns(d.DBMaxConnection)
	sqlDB.SetMaxIdleConns(d.DBIdleConnection)
	sqlDB.SetConnMaxLifetime(d.DBConnMaxLifeTime)
	return db, nil
}

// GetDatabaseConfig is get DatabaseConfig from env.
func GetDatabaseConfig() DatabaseConfig {
	config := DatabaseConfig{
		DBHost:            databaseViper.GetString(DBHost),
		DBPort:            databaseViper.GetInt(DBPort),
		DBName:            databaseViper.GetString(DBName),
		DBUser:            databaseViper.GetString(DBUser),
		DBPassword:        databaseViper.GetString(DBPassword),
		DBIdleConnection:  databaseViper.GetInt(DBIdleConnection),
		DBMaxConnection:   databaseViper.GetInt(DBMaxConnection),
		DBConnMaxLifeTime: time.Duration(databaseViper.GetInt(DBConnMaxLifeTime)) * time.Second,
		DBTimeZone:        databaseViper.GetString(DBTimeZone),
	}
	return config
}
