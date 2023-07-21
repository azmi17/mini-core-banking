package echannelrepo

import (
	"database/sql"
	"errors"
	"new-apex-api/repository/databasefactory"
	"new-apex-api/repository/databasefactory/drivers"
)

func NewEchannelRepo() (EchannelRepo, error) {
	apexConn := databasefactory.Echannel.GetConnection()
	currentDriver := databasefactory.Echannel.GetDriverName()
	if currentDriver == drivers.MYSQL {
		return newEchannelMysqlImpl(apexConn.(*sql.DB)), nil
	} else {
		return nil, errors.New("unimplemented database driver")
	}
}
