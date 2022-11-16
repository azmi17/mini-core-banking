package sysuserrepo

import "apex-ems-integration-clean-arch/entities"

type SysUserRepo interface {
	CreateSysDaftarUser(newSysUser entities.SysDaftarUser) (entities.SysDaftarUser, error)
	UpdateSysDaftarUser(updNasabah entities.SysDaftarUser) (entities.SysDaftarUser, error)
	HardDeleteSysDaftarUser(kodeLkm string) error
	DeleteSysDaftarUser(kodeLkm string) error
	ResetUserPassword(user entities.SysDaftarUser) (entities.SysDaftarUser, error)
	FindByUserName(userName string) (entities.SysDaftarUser, error)
}
