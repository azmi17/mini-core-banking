package sysuserrepo

import (
	"apex-ems-integration-clean-arch/repository/databasefactory"
	"apex-ems-integration-clean-arch/repository/databasefactory/drivers"
	"database/sql"
	"errors"
)

func NewSysUserRepo() (SysUserRepo, error) {
	apexConn := databasefactory.SysApex.GetConnection()
	currentDriver := databasefactory.SysApex.GetDriverName()
	if currentDriver == drivers.MYSQL {
		return newSysUserMysqlImpl(apexConn.(*sql.DB)), nil
	} else {
		return nil, errors.New("unimplemented database driver")
	}
}
