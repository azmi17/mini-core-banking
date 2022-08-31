package databasefactory

import (
	"apex-ems-integration-clean-arch/repository/databasefactory/drivers"
	"errors"
	"os"
)

func GetDatabase() (db Database, err error) {

	driverName := os.Getenv("app.database_driver")

	// Default driver will return MYSQL-Driver
	if driverName == "" {
		driverName = drivers.MYSQL
	}

	if driverName == drivers.MYSQL {
		return newMysqlImpl(), nil
	} else {
		return db, errors.New("unimplement database driver")
	}
}
