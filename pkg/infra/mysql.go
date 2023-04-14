package infra

import (
	"context"
	"fmt"
	"log"
	"order-service/pkg/config"
	"order-service/pkg/constants"
	"order-service/service/model/entity"
	"order-service/service/model/mock"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// PGDB is wrapper of pg.DB
type SQLDB struct {
	gormDB *gorm.DB
}

var sqlSingleton *SQLDB

func InitMySQL() {
	dbConfig := config.DbConfig()
	// https://github.com/go-gorm/postgres
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf(constants.DBConnectionString,
			dbConfig.DBUserName,
			dbConfig.DBPassword,
			dbConfig.DBURL,
			dbConfig.DBPort,
			dbConfig.DBName,
		),
		// DSN:                       "root:1234@tcp(database:3308)/test?charset=utf8&parseTime=True&loc=Local", // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	// }), &gorm.Config{CreateBatchSize: 1000})

	if err != nil {
		log.Fatal(fmt.Errorf("Can't instance db connection error: %v", err))
	}

	// auto migrate db tables
	db.AutoMigrate(
		&entity.User{}, &entity.Category{}, &entity.Product{},
	)

	go mock.MakeMockData(db)

	sqlSingleton = &SQLDB{db}
}

func ClosePostgresql() error {
	db, err := sqlSingleton.gormDB.DB()
	db.Close()
	if err != nil {
		log.Fatal(fmt.Errorf("Close connection db error: %v", err))
	}
	return nil
}

func GetDB() *gorm.DB {
	if sqlSingleton == nil {
		log.Fatal("Connection to database Postgres is not setup")
	}
	return sqlSingleton.gormDB
}

// BeginTransaction start an Transaction, require defer ReleaseTransaction instantly
func BeginTransaction() (*gorm.DB, error) {
	tx := sqlSingleton.gormDB.Begin()
	return tx, nil
}

func CommitTransaction(ctx context.Context, tx *gorm.DB) {
	tx.Commit()
}

func RollbackTransaction(ctx context.Context, tx *gorm.DB) {
	tx.Rollback()
}

func ReleaseTransaction(tx *gorm.DB, err error) {
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
}
