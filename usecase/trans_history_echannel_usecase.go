package usecase

import (
	"new-apex-api/entities"
	"new-apex-api/entities/err"
	"new-apex-api/repository/echannelrepo"
)

type TransHisotryEchannelUsecase interface {
	TransHistoriesLists(payload entities.TransHistoryRequest, limitOffset entities.LimitOffsetLkmUri) (list []entities.TransHistoryResponse, er error)
}

type transHisotryEchannelUsecase struct{}

func NewtransHisotryEchannelUsecase() TransHisotryEchannelUsecase {
	return &transHisotryEchannelUsecase{}
}

func (e *transHisotryEchannelUsecase) TransHistoriesLists(payload entities.TransHistoryRequest, limitOffset entities.LimitOffsetLkmUri) (list []entities.TransHistoryResponse, er error) {
	if limitOffset.Limit <= 0 || limitOffset.Offset < 0 {
		return list, err.BadRequest
	}

	payloadCfg := entities.TransHistoryRequest{}
	payloadCfg.Filter = payload.Filter
	payloadCfg.TglAwal = payload.TglAwal + "000000"
	payloadCfg.TglAkhir = payload.TglAwal + "595959"

	echannelRepo, _ := echannelrepo.NewEchannelRepo()
	list, er = echannelRepo.GetEchannelTransHistories(payloadCfg, limitOffset)
	if er != nil {
		return list, er
	}

	if len(list) == 0 {
		return make([]entities.TransHistoryResponse, 0), nil
		// ^ []entities.GetDetailLKMInfo (untuk membuat sebuah slice kosong agar tidak return null di JSON) | err.NoRecord
	}

	return list, nil
}
