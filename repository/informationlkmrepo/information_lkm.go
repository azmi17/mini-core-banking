package informationlkmrepo

import (
	"new-apex-api/entities"
)

type InformationLKMRepo interface {
	RekeningKoranLKMDetailHeader(kodeLKM string) (entities.RekeningKoranHeader, error)
	LKMInformation(kodeLKM string) (entities.LkmInfo, error)
}
