package routingindukrepo

import (
	"apex-ems-integration-clean-arch/repository/databasefactory"
	"apex-ems-integration-clean-arch/repository/databasefactory/drivers"
	"database/sql"
	"errors"
)

func NewRoutingIndukRepo() (RoutingIndukRepo, error) {
	apexConn := databasefactory.SysApex.GetConnection()
	currentDriver := databasefactory.SysApex.GetDriverName()
	if currentDriver == drivers.MYSQL {
		return newRoutingIndukMysqlImpl(apexConn.(*sql.DB)), nil
	} else {
		return nil, errors.New("unimplemented database driver")
	}
}
