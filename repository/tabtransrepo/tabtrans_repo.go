package tabtransrepo

import (
	"new-apex-api/entities"
	"new-apex-api/entities/web"
)

type TabtransRepo interface {
	GetNextIDWithUserID() (int, error)
	GetSingleTabtransTrx(tabtransID int) (web.GetListTabtransTrx, error)
	GetListsTabtransTrx(TglTrans web.GetListTabtransByDate, limitOffset web.LimitOffsetLkmUri) (
		[]web.GetListTabtransTrx,
		web.GetCountWithSumTabtransTrx,
		error,
	)
	GetListsTabtransTrxBySTAN(stan string) ([]web.GetListTabtransTrx, error)
	DeleteTabtransTrx(tabtransID int) error
	ChangeDateOnTabtransTrx(tabtransID int, tglTrans string) (web.GetListTabtransTrx, error)

	CalculateSaldoOnRekeningLKM(kodeLKM string) (web.CalculateSaldoResult, error)
	RepostingSaldoOnRekeningLKM(listOfKodeLKM ...string) error
	doRepostingSaldoProcs(data string) error
	RepostingSaldoOnRekeningLKMByScheduler(PrintRepoResult chan entities.PrintRepo, listOfKodeLKM ...string) error
	// doRepostingSaldoProcsByScheduler(data string) error

	GetSaldo(kodeLKM, tglAwal string) (float64, error)
	GetRekeningKoranLKMDetail(kodeLKM, periodeAwal, periodeAkhir string) ([]web.RekeningKoran, error)
	GetNominatifDeposit(periode, beginLastMonthDt, endLastmonthDt string, endLastMonthTg int, limitOffset web.LimitOffsetLkmUri) ([]web.NominatifDepositResponse, error)
	GetLaporanTransaksi(payload web.DaftarTransaksiRequest, limitOffset web.LimitOffsetLkmUri) ([]web.DaftarTransaksiResponse, error)
	GetListsTransaksiDeposit(payload web.GetListsDepositTrxReq) ([]web.GetListsDepositTrxRes, error)
}
