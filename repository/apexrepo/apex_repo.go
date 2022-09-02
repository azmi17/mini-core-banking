package apexrepo

import (
	"apex-ems-integration-clean-arch/entities"
)

type ApexRepo interface {
	SaveNasabah(newNasabah entities.Nasabah) (entities.Nasabah, error)
	SaveTabung(newTabung entities.Tabung) (entities.Tabung, error)
	SaveSysDaftarUser(newSysUser entities.SysDaftarUser) (entities.SysDaftarUser, error)
}
