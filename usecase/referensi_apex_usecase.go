package usecase

import (
	"new-apex-api/entities/err"
	"new-apex-api/entities/web"
	"new-apex-api/repository/referensiapex"
)

type ReferensiApexUsecase interface {
	GetListsScGroup() ([]web.ScGroup, error)
	GetListsJenisTransaksiTabungan() ([]web.JenisTransaksi, error)
	GetListsJenisTransaksiDeposit() ([]web.JenisTransaksi, error)
	GetListsBankGroup() ([]web.BankGroup, error)
	GetlistsProdukTabungan() ([]web.ProdukTabungan, error)
}

type referensiApexUsecase struct{}

func NewReferensiApexUsecase() ReferensiApexUsecase {
	return &referensiApexUsecase{}
}

func (r *referensiApexUsecase) GetListsScGroup() (lists []web.ScGroup, er error) {
	ref, _ := referensiapex.NewReferensiApexRepo()

	lists, er = ref.GetListsScGroup()
	if er != nil {
		return lists, er
	}

	if len(lists) == 0 {
		return lists, err.NoRecord
	}

	return lists, nil
}

func (r *referensiApexUsecase) GetListsJenisTransaksiTabungan() (lists []web.JenisTransaksi, er error) {
	ref, _ := referensiapex.NewReferensiApexRepo()

	lists, er = ref.GetListsJenisTransaksiTabungan()
	if er != nil {
		return lists, er
	}

	if len(lists) == 0 {
		return lists, err.NoRecord
	}

	return lists, nil
}

func (r *referensiApexUsecase) GetListsJenisTransaksiDeposit() (lists []web.JenisTransaksi, er error) {
	ref, _ := referensiapex.NewReferensiApexRepo()

	lists, er = ref.GetListsJenisTransaksiDeposit()
	if er != nil {
		return lists, er
	}

	if len(lists) == 0 {
		return lists, err.NoRecord
	}

	return lists, nil
}

func (r *referensiApexUsecase) GetListsBankGroup() (lists []web.BankGroup, er error) {
	ref, _ := referensiapex.NewReferensiApexRepo()

	lists, er = ref.GetListsBankGroup()
	if er != nil {
		return lists, er
	}

	if len(lists) == 0 {
		return lists, err.NoRecord
	}

	return lists, nil
}

func (r *referensiApexUsecase) GetlistsProdukTabungan() (lists []web.ProdukTabungan, er error) {
	ref, _ := referensiapex.NewReferensiApexRepo()

	lists, er = ref.GetlistsProdukTabungan()
	if er != nil {
		return lists, er
	}

	if len(lists) == 0 {
		return lists, err.NoRecord
	}

	return lists, nil
}
