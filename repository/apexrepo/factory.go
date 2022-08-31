package apexrepo

import (
	"apex-ems-integration-clean-arch/repository/databasefactory"
	"apex-ems-integration-clean-arch/repository/databasefactory/drivers"
	"database/sql"
	"errors"
)

func NewApexRepo() (ApexRepo, error) {

	conn1 := databasefactory.AppDb1.GetConnection()

	currentDriver := databasefactory.AppDb1.GetDriverName()
	if currentDriver == drivers.MYSQL {
		conn2 := databasefactory.AppDb2.GetConnection()
		return newApexMysqlImpl(conn1.(*sql.DB), conn2.(*sql.DB)), nil
	} else {
		return nil, errors.New("unimplemented database driver")
	}

}
