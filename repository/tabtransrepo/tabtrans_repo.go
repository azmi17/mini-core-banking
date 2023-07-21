package tabtransrepo

import (
	"new-apex-api/entities"
)

type TabtransRepo interface {
	GetNextIDWithUserID() (int, error)
	GetNextNasabahID() (string, error)
	GetNextNoRekeningBiggerThanFour() (string, error)
	GetInformationLKMTransaction(tabtransID int) (entities.GetListTabtransTrx, error)
	GetListsApexTransaction(payload entities.GetListTabtrans, limitOffset entities.LimitOffsetLkmUri) (response []entities.GetListTabtransTrx, er error)
	GetListsApexTransactionBySTAN(stan string) ([]entities.GetListTabtransTrx, error)
	HardDeleteApexTransaction(data ...int) error
	SoftDeleteApexTransaction(data ...string) error
	ChangeDateOnApexTransaction(tabtransID int, tglTrans string) (entities.GetListTabtransTrx, error)
	MultipleChangeDateOnApexTransaction(data entities.MultipleChangeDateTransaction) error
	CalculateSaldoOnRekeningLKM(kodeLKM string) (entities.CalculateSaldoResult, error)
	RepostingSaldoOnRekeningLKM(listOfKodeLKM ...string) error
	doRepostingSaldoProcs(data string) error
	RepostingSaldoOnRekeningLKMByScheduler(PrintRepoResult chan entities.PrintRepo, listOfKodeLKM ...string) error
	// doRepostingSaldoProcsByScheduler(data string) error
	GetSaldo(kodeLKM, tglAwal string) (float64, error)
	GetRekeningKoranLKMDetail(kodeLKM, periodeAwal, periodeAkhir string) ([]entities.RekeningKoran, error)
	GetNominatifDeposit(periode, beginLastMonthDt, endLastmonthDt string, endLastMonthTg int, limitOffset entities.LimitOffsetLkmUri) ([]entities.NominatifDepositResponse, error)
	GetLaporanTransaksi(payload entities.DaftarTransaksiRequest, limitOffset entities.LimitOffsetLkmUri) ([]entities.DaftarTransaksiResponse, error)
	GetListsTransaksiDeposit(payload entities.GetListsDepositTrxReq) ([]entities.GetListsDepositTrxRes, error)
	TransaksiDeposit(data ...entities.Deposit) error
	GetTransaksiDeposit(kuitansi string) (data entities.Deposit, er error)
}
