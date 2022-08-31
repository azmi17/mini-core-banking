package apexrepo

import (
	"apex-ems-integration-clean-arch/entities"
)

type ApexRepo interface {
	CreateNasabah(newNasabah entities.Nasabah) (entities.Nasabah, error)
	CreateTabung(newTabung entities.Tabung) (entities.Tabung, error)
	CreteSysDaftarUser(newSysUser entities.SysDaftarUser) (entities.SysDaftarUser, error)
}
