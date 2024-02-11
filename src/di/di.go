package di

import (
	"fmt"
	"github.com/hamidteimouri/gommon/htenvier"
	"github.com/sarmayex/phishdefender/data/postgres"
	"github.com/sarmayex/phishdefender/domain"
	"github.com/sarmayex/phishdefender/presentation/handlers"
	"github.com/sirupsen/logrus"
	gormPostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db              *gorm.DB
	dbSource        *postgres.Postgres
	defenderHandler *handlers.DefenderHandler
	defenderService *domain.PhishDefenderService
)

func dbInitializer() *gorm.DB {
	if db != nil {
		return db
	}

	var err error
	dbHost := htenvier.Env("DB_HOST")
	dbPort := htenvier.Env("DB_PORT")
	dbName := htenvier.Env("DB_NAME")
	dbUsername := htenvier.Env("DB_USERNAME")
	dbPassword := htenvier.Env("DB_PASSWORD")
	dbTimezone := htenvier.Env("DB_TIMEZONE")

	// logger of gorm
	gormLogger := logger.Default.LogMode(logger.Silent)

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=%v",
		dbHost, dbUsername, dbPassword, dbName, dbPort, dbTimezone)
	db, err = gorm.Open(gormPostgres.Open(dsn), &gorm.Config{Logger: gormLogger, SkipDefaultTransaction: true})

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Panic("database connection error")
	}
	d, err := db.DB()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Panic("failed to get DB instance")
	}
	if err = d.Ping(); err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Panic("database has no ping")
	}

	return db
}

func dbDatasource() *postgres.Postgres {
	if dbSource == nil {
		dbSource = postgres.NewPostgres(dbInitializer())
	}
	return dbSource
}

func DefenderService() *domain.PhishDefenderService {
	if defenderService == nil {
		defenderService = domain.NewPhishDefenderService(dbDatasource())
	}
	return defenderService
}
func DefenderHandler() *handlers.DefenderHandler {
	if defenderHandler == nil {
		defenderHandler = handlers.NewDefenderHandler(DefenderService())
	}
	return defenderHandler
}
