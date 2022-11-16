package nasabahrepo

import (
	"apex-ems-integration-clean-arch/repository/databasefactory"
	"apex-ems-integration-clean-arch/repository/databasefactory/drivers"
	"database/sql"
	"errors"
)

func NewNasabahRepo() (NasabahRepo, error) {
	apexConn := databasefactory.Apex.GetConnection()
	currentDriver := databasefactory.Apex.GetDriverName()
	if currentDriver == drivers.MYSQL {
		return newNasbahMysqlImpl(apexConn.(*sql.DB)), nil
	} else {
		return nil, errors.New("unimplemented database driver")
	}
}
