package usecase

import (
	"new-apex-api/entities"
	"new-apex-api/entities/err"
	"new-apex-api/repository/referensiapex"
)

type ReferensiApexUsecase interface {
	GetListsScGroup() ([]entities.ScGroup, error)
	GetListsJenisTransaksiTabungan() ([]entities.JenisTransaksi, error)
	GetListsJenisTransaksiDeposit() ([]entities.JenisTransaksi, error)
	GetListsBankGroup() ([]entities.BankGroup, error)
	GetlistsProdukTabungan() ([]entities.ProdukTabungan, error)
	GetListsJenisPembayaranSLA() ([]entities.JenisPembayaranSLA, error)
	GetListsTabunganIntegrasi() ([]entities.TabunganIntegrasi, error)
}

type referensiApexUsecase struct{}

func NewReferensiApexUsecase() ReferensiApexUsecase {
	return &referensiApexUsecase{}
}

func (r *referensiApexUsecase) GetListsScGroup() (lists []entities.ScGroup, er error) {
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

func (r *referensiApexUsecase) GetListsJenisTransaksiTabungan() (lists []entities.JenisTransaksi, er error) {
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

func (r *referensiApexUsecase) GetListsJenisTransaksiDeposit() (lists []entities.JenisTransaksi, er error) {
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

func (r *referensiApexUsecase) GetListsBankGroup() (lists []entities.BankGroup, er error) {
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

func (r *referensiApexUsecase) GetlistsProdukTabungan() (lists []entities.ProdukTabungan, er error) {
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

func (r *referensiApexUsecase) GetListsJenisPembayaranSLA() (lists []entities.JenisPembayaranSLA, er error) {
	ref, _ := referensiapex.NewReferensiApexRepo()

	lists, er = ref.GetListsJenisPembayaranSLA()
	if er != nil {
		return lists, er
	}

	if len(lists) == 0 {
		return lists, err.NoRecord
	}

	return lists, nil
}

func (r *referensiApexUsecase) GetListsTabunganIntegrasi() (lists []entities.TabunganIntegrasi, er error) {
	ref, _ := referensiapex.NewReferensiApexRepo()

	lists, er = ref.GetListsTabunganIntegrasi()
	if er != nil {
		return lists, er
	}

	if len(lists) == 0 {
		return lists, err.NoRecord
	}

	return lists, nil
}
