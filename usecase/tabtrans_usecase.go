package usecase

import (
	"apex-ems-integration-clean-arch/entities/err"
	"apex-ems-integration-clean-arch/entities/web"
	"apex-ems-integration-clean-arch/repository/tabtransrepo"
)

type TabtransUsecase interface {
	GetListTabtransInfo(tglTrans web.GetListTabtransByDate, limitOffset web.LimitOffsetLkmUri) (
		[]web.GetListTabtransInfo,
		web.GetCountWithSumTabtransTrx,
		error,
	)
}

type tabtransUsecase struct{}

func NewTabtransUsecase() TabtransUsecase {
	return &tabtransUsecase{}
}

func (t *tabtransUsecase) GetListTabtransInfo(tglTrans web.GetListTabtransByDate, limitOffset web.LimitOffsetLkmUri) (
	tabtransTxList []web.GetListTabtransInfo,
	total web.GetCountWithSumTabtransTrx,
	er error,
) {
	if limitOffset.Limit <= 0 || limitOffset.Offset < 0 {
		return tabtransTxList, total, err.BadRequest
	}

	tabunganRepo, _ := tabtransrepo.NewTabtransRepo()
	tabtransTxList, total, er = tabunganRepo.GetListTabtransInfo(tglTrans, limitOffset)
	if er != nil {
		return tabtransTxList, total, er
	}

	if len(tabtransTxList) == 0 {
		return make([]web.GetListTabtransInfo, 0), total, nil
		// ^ []web.GetDetailLKMInfo (untuk membuat sebuah slice kosong agar tidak return null di JSON) | err.NoRecord
	}

	return tabtransTxList, total, nil
}
