package nasabahrepo

import "apex-ems-integration-clean-arch/entities"

type NasabahRepo interface {
	CreateNasabah(newNasabah entities.Nasabah) (entities.Nasabah, error)
	UpdateNasabah(updNasabah entities.Nasabah) (entities.Nasabah, error)
	HardDeleteNasabah(kodeLkm string) error
	DeleteNasabah(kodeLkm string) error
}
