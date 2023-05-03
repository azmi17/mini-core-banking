package referensiapex

import (
	"new-apex-api/entities/web"
)

type ReferensiApexRepo interface {
	GetListsScGroup() ([]web.ScGroup, error)
	GetListsJenisTransaksiTabungan() ([]web.JenisTransaksi, error)
	GetListsJenisTransaksiDeposit() ([]web.JenisTransaksi, error)
	GetListsBankGroup() ([]web.BankGroup, error)
	GetlistsProdukTabungan() ([]web.ProdukTabungan, error)
}
