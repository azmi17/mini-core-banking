package tabunganrepo

import (
	"apex-ems-integration-clean-arch/repository/databasefactory"
	"apex-ems-integration-clean-arch/repository/databasefactory/drivers"
	"database/sql"
	"errors"
)

func NewTabunganRepo() (TabunganRepo, error) {
	apexConn := databasefactory.Apex.GetConnection()
	currentDriver := databasefactory.Apex.GetDriverName()
	if currentDriver == drivers.MYSQL {
		return newTabunganMysqlImpl(apexConn.(*sql.DB)), nil
	} else {
		return nil, errors.New("unimplemented database driver")
	}
}
