package apexrepo

import (
	"apex-ems-integration-clean-arch/entities"
	"apex-ems-integration-clean-arch/entities/web"
)

type ApexRepo interface {
	CreateNasabah(newNasabah entities.Nasabah) (entities.Nasabah, error)
	CreateTabung(newTabung entities.Tabung) (entities.Tabung, error)
	CreateSysDaftarUser(newSysUser entities.SysDaftarUser) (entities.SysDaftarUser, error)

	UpdateNasabah(updNasabah entities.Nasabah) (entities.Nasabah, error)
	UpdateSysDaftarUser(updNasabah entities.SysDaftarUser) (entities.SysDaftarUser, error)

	DeleteNasabah(kodeLkm string) error
	DeleteTabung(kodeLkm string) error
	DeleteSysDaftarUser(kodeLkm string) error

	GetScGroup() ([]web.SCGroup, error)
	GetLkmDetailInfo(KodeLkm string) (web.GetDetailLKMInfo, error)
	GetLkmInfoList(limitOffset web.LimitOffsetLkmUri) ([]web.GetDetailLKMInfo, error)

	ResetApexPassword(user entities.SysDaftarUser) (entities.SysDaftarUser, error)
}
