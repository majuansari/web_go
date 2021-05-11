package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"web/config"
	"web/pkg/telemetry/otelgorm"
)

func NewDBConnection(cfg config.DBConfig) (*gorm.DB, func(), error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       cfg.Dsn(), // data source name
		DefaultStringSize:         256,       // default size for string fields
		DisableDatetimePrecision:  true,      // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,      // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,      // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,     // auto configure based on currently MySQL version
	}), &gorm.Config{})

	if err != nil {
		return nil, nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, err
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(cfg.ConMaxLifeTime)

	// Initialize otel plugin with options
	plugin := otelgorm.NewPlugin(
	// include any options here
	)
	err = db.Use(plugin)
	if err != nil {
		panic("failed configuring plugin")
	}

	return db, func() { sqlDB.Close() }, nil
}
