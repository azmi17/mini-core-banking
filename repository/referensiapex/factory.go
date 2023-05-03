package referensiapex

import (
	"database/sql"
	"errors"
	"new-apex-api/repository/databasefactory"
	"new-apex-api/repository/databasefactory/drivers"
)

func NewReferensiApexRepo() (ReferensiApexRepo, error) {
	apexConn := databasefactory.Apex.GetConnection()
	currentDriver := databasefactory.Apex.GetDriverName()
	if currentDriver == drivers.MYSQL {
		return newReferensiApexMysqlImpl(apexConn.(*sql.DB)), nil
	} else {
		return nil, errors.New("unimplemented database driver")
	}
}
