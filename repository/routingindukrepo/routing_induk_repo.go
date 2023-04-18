package routingindukrepo

import "new-apex-api/entities/web"

type RoutingIndukRepo interface {
	GetRoutingRekInduk(kodeLkm string) (web.RoutingRekIndukData, error)
	GetListSysApexRoutingRekInduk(limitOffset web.LimitOffsetLkmUri) ([]web.RoutingRekIndukData, error)
	CreateSysApexRoutingRekInduk(bankCode, norekInduk string) (web.RoutingRekIndukData, error)
	UpdateSysApexRoutingRekInduk(newBankCode, norekInduk, currentBankCode string) (web.RoutingRekIndukData, error)
	DeleteSysApexRoutingRekInduk(kodeLkm string) error
}
