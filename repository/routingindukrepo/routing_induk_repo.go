package routingindukrepo

import (
	"new-apex-api/entities"
)

type RoutingIndukRepo interface {
	GetRoutingRekInduk(kodeLkm string) (entities.RoutingRekIndukData, error)
	GetListSysApexRoutingRekInduk(payload entities.GlobalFilter, limitOffset entities.LimitOffsetLkmUri) ([]entities.RoutingRekIndukData, error)
	CreateSysApexRoutingRekInduk(bankCode, norekInduk string) (entities.RoutingRekIndukData, error)
	UpdateSysApexRoutingRekInduk(newBankCode, norekInduk, currentBankCode string) (entities.RoutingRekIndukData, error)
	DeleteSysApexRoutingRekInduk(kodeLkm ...string) error
}
