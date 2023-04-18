package informationlkmrepo

import "new-apex-api/entities/web"

type InformationLKMRepo interface {
	RekeningKoranLKMDetailHeader(kodeLKM string) (web.RekeningKoranHeader, error)
}
