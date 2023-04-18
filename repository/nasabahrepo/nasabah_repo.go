package nasabahrepo

import "new-apex-api/entities"

type NasabahRepo interface {
	FindNasabahLkm(nasabahId string) (entities.Nasabah, error)
	CreateNasabah(newNasabah entities.Nasabah) (entities.Nasabah, error)
	UpdateNasabah(updNasabah entities.Nasabah) (entities.Nasabah, error)
	HardDeleteNasabah(kodeLkm string) error
	DeleteNasabah(kodeLkm string) error
}
