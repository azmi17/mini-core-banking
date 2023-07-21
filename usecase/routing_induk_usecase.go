package usecase

import (
	"new-apex-api/entities"
	"new-apex-api/entities/err"
	"new-apex-api/repository/routingindukrepo"
)

type RoutingIndukUsecase interface {
	GetRoutingRekInduk(kodeLkm string) (entities.RoutingRekIndukData, error)
	GetListRoutingRekInduk(payload entities.GlobalFilter, limitOffset entities.LimitOffsetLkmUri) ([]entities.RoutingRekIndukData, error)
	CreateSysApexRoutingRekInduk(entities.CreateRoutingRekInduk) (entities.RoutingRekIndukData, error)
	UpdateSysApexRoutingRekInduk(entities.UpdateRoutingRekInduk) (entities.RoutingRekIndukData, error)
	DeleteSysApexRoutingRekInduk(kodeLkm []string) error
}

type routingIndukUsecase struct{}

func NewRoutingIndukUsecase() RoutingIndukUsecase {
	return &routingIndukUsecase{}
}

func (r *routingIndukUsecase) GetRoutingRekInduk(kodeLkm string) (routingInfo entities.RoutingRekIndukData, er error) {
	routingIndukRepo, _ := routingindukrepo.NewRoutingIndukRepo()

	if routingInfo, er = routingIndukRepo.GetRoutingRekInduk(kodeLkm); er != nil {
		return routingInfo, er
	}

	return routingInfo, nil
}

func (r *routingIndukUsecase) GetListRoutingRekInduk(payload entities.GlobalFilter, limitOffset entities.LimitOffsetLkmUri) (routingList []entities.RoutingRekIndukData, er error) {
	if limitOffset.Limit <= 0 || limitOffset.Offset < 0 {
		return routingList, err.BadRequest
	}

	routingIndukRepo, _ := routingindukrepo.NewRoutingIndukRepo()
	routingList, er = routingIndukRepo.GetListSysApexRoutingRekInduk(payload, limitOffset)
	if er != nil {
		return routingList, er
	}

	if len(routingList) == 0 {
		return routingList, err.NoRecord
	}

	return routingList, nil
}

func (r *routingIndukUsecase) CreateSysApexRoutingRekInduk(payload entities.CreateRoutingRekInduk) (data entities.RoutingRekIndukData, er error) {
	routingIndukRepo, _ := routingindukrepo.NewRoutingIndukRepo()

	data.KodeLkm = payload.KodeLkm
	data.NorekInduk = payload.NorekInduk

	if data, er = routingIndukRepo.CreateSysApexRoutingRekInduk(data.KodeLkm, data.NorekInduk); er != nil {
		return data, er
	}

	return
}

func (r *routingIndukUsecase) UpdateSysApexRoutingRekInduk(payload entities.UpdateRoutingRekInduk) (data entities.RoutingRekIndukData, er error) {
	routingIndukRepo, _ := routingindukrepo.NewRoutingIndukRepo()

	data.KodeLkm = payload.KodeLkm
	data.NorekInduk = payload.NorekInduk

	if data, er = routingIndukRepo.UpdateSysApexRoutingRekInduk(data.KodeLkm, data.NorekInduk, payload.KodeLkmTarget); er != nil {
		return data, er
	}

	return
}

func (r *routingIndukUsecase) DeleteSysApexRoutingRekInduk(kodeLkm []string) (er error) {
	routingIndukRepo, _ := routingindukrepo.NewRoutingIndukRepo()

	if er = routingIndukRepo.DeleteSysApexRoutingRekInduk(kodeLkm...); er != nil {
		return er
	}

	return nil
}
