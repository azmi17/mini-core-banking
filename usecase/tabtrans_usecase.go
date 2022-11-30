package usecase

import (
	"apex-ems-integration-clean-arch/entities/err"
	"apex-ems-integration-clean-arch/entities/web"
	"apex-ems-integration-clean-arch/repository/tabtransrepo"
)

type TabtransUsecase interface {
	GetListsTabtransTrx(tglTrans web.GetListTabtransByDate, limitOffset web.LimitOffsetLkmUri) (
		[]web.GetListTabtransTrx,
		web.GetCountWithSumTabtransTrx,
		error,
	)
	GetListsTabtransTrxBySTAN(stan string) ([]web.GetListTabtransTrx, error)
	DeleteTabtransTrx(tabtransID int) error
	ChangeDateOnTabtransTrx(tabtransID int, tglTrans string) (web.GetListTabtransTrx, error)
}

type tabtransUsecase struct{}

func NewTabtransUsecase() TabtransUsecase {
	return &tabtransUsecase{}
}

func (t *tabtransUsecase) GetListsTabtransTrx(tglTrans web.GetListTabtransByDate, limitOffset web.LimitOffsetLkmUri) (
	tabtransTxList []web.GetListTabtransTrx,
	total web.GetCountWithSumTabtransTrx,
	er error,
) {
	if limitOffset.Limit <= 0 || limitOffset.Offset < 0 {
		return tabtransTxList, total, err.BadRequest
	}

	tabtransRepo, _ := tabtransrepo.NewTabtransRepo()
	tabtransTxList, total, er = tabtransRepo.GetListsTabtransTrx(tglTrans, limitOffset)
	if er != nil {
		return tabtransTxList, total, er
	}

	if len(tabtransTxList) == 0 {
		return make([]web.GetListTabtransTrx, 0), total, nil
		// ^ []web.GetDetailLKMInfo (untuk membuat sebuah slice kosong agar tidak return null di JSON) | err.NoRecord
	}

	return tabtransTxList, total, nil
}

func (t *tabtransUsecase) GetListsTabtransTrxBySTAN(stan string) (tx []web.GetListTabtransTrx, er error) {
	tabtransRepo, _ := tabtransrepo.NewTabtransRepo()

	tx, er = tabtransRepo.GetListsTabtransTrxBySTAN(stan)
	if er != nil {
		return tx, er
	}

	if len(tx) == 0 {
		return tx, err.NoRecord
	}

	return tx, nil
}

func (t *tabtransUsecase) DeleteTabtransTrx(tabtransID int) (er error) {
	tabtransRepo, _ := tabtransrepo.NewTabtransRepo()

	if er = tabtransRepo.DeleteTabtransTrx(tabtransID); er != nil {
		return er
	}

	return nil
}

func (t *tabtransUsecase) ChangeDateOnTabtransTrx(tabtransID int, tglTrans string) (data web.GetListTabtransTrx, er error) {
	tabtransRepo, _ := tabtransrepo.NewTabtransRepo()

	if data, er = tabtransRepo.ChangeDateOnTabtransTrx(tabtransID, tglTrans); er != nil {
		return data, er
	}

	return
}
