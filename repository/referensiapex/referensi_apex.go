package referensiapex

import "new-apex-api/entities"

type ReferensiApexRepo interface {
	GetListsScGroup() ([]entities.ScGroup, error)
	GetListsJenisTransaksiTabungan() ([]entities.JenisTransaksi, error)
	GetListsJenisTransaksiDeposit() ([]entities.JenisTransaksi, error)
	GetListsBankGroup() ([]entities.BankGroup, error)
	GetlistsProdukTabungan() ([]entities.ProdukTabungan, error)
	GetListsJenisPembayaranSLA() ([]entities.JenisPembayaranSLA, error)
	GetListsTabunganIntegrasi() ([]entities.TabunganIntegrasi, error)
}
