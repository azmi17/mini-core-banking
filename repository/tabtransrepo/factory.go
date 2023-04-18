package tabtransrepo

import (
	"database/sql"
	"errors"
	"new-apex-api/repository/databasefactory"
	"new-apex-api/repository/databasefactory/drivers"
)

func NewTabtransRepo() (TabtransRepo, error) {
	apexConn := databasefactory.Apex.GetConnection()
	currentDriver := databasefactory.Apex.GetDriverName()
	if currentDriver == drivers.MYSQL {
		return newTabtransMysqlImpl(apexConn.(*sql.DB)), nil
	} else {
		return nil, errors.New("unimplemented database driver")
	}
}
