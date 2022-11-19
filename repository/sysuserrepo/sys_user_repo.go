package sysuserrepo

import (
	"apex-ems-integration-clean-arch/entities"
	"apex-ems-integration-clean-arch/entities/web"
)

type SysUserRepo interface {
	CreateSysDaftarUser(newSysUser entities.SysDaftarUser) (entities.SysDaftarUser, error)
	UpdateSysDaftarUser(updNasabah entities.SysDaftarUser) (entities.SysDaftarUser, error)
	HardDeleteSysDaftarUser(kodeLkm string) error
	DeleteSysDaftarUser(kodeLkm string) error
	ResetUserPassword(user entities.SysDaftarUser) (entities.SysDaftarUser, error)
	FindByUserName(userName string) (entities.SysDaftarUser, error)
	GetListSysApexRoutingRekInduk() ([]web.RoutingRekIndukData, error)
	CreateSysApexRoutingRekInduk(bankCode, norekInduk string) (web.RoutingRekIndukData, error)
	UpdateSysApexRoutingRekInduk(newBankCode, norekInduk, currentBankCode string) (web.RoutingRekIndukData, error)
	DeleteSysApexRoutingRekInduk(kodeLkm string) error
}
