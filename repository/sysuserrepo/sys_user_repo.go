package sysuserrepo

import (
	"apex-ems-integration-clean-arch/entities"
	"apex-ems-integration-clean-arch/entities/web"
)

type SysUserRepo interface {
	GetSingleUserByUserName(userName string) (web.ManajemenUserDataResponse, error)
	GetListOfUsers(limitOffset web.LimitOffsetLkmUri) ([]web.ManajemenUserDataResponse, error)
	CreateSysDaftarUser(newSysUser entities.SysDaftarUser) (entities.SysDaftarUser, error)
	UpdateSysDaftarUser(updNasabah entities.SysDaftarUser) (entities.SysDaftarUser, error)
	HardDeleteSysDaftarUser(kodeLkm string) error
	DeleteSysDaftarUser(kodeLkm string) error
	ResetUserPassword(user entities.SysDaftarUser) (entities.SysDaftarUser, error)
	FindByUserName(userName string) (entities.SysDaftarUser, error)
}
