package informationlkmrepo

import (
	"new-apex-api/entities"
	"new-apex-api/entities/web"
)

type InformationLKMRepo interface {
	RekeningKoranLKMDetailHeader(kodeLKM string) (web.RekeningKoranHeader, error)
	LKMInformation(kodeLKM string) (entities.LkmInfo, error)
}
