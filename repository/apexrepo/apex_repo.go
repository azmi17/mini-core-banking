package apexrepo

import (
	"apex-ems-integration-clean-arch/entities"
)

type ApexRepo interface {
	CreateNasabah(newNasabah entities.Nasabah) (entities.Nasabah, error)
	CreateTabung(newTabung entities.Tabung) (entities.Tabung, error)
	CreateSysDaftarUser(newSysUser entities.SysDaftarUser) (entities.SysDaftarUser, error)

	UpdateNasabah(updNasabah entities.Nasabah) (entities.Nasabah, error)
	UpdateSysDaftarUser(updNasabah entities.SysDaftarUser) (entities.SysDaftarUser, error)

	DeleteNasabah(id string) error
	DeleteTabung(id string) error
	DeleteSysDaftarUser(id string) error
}
