package tabtransrepo

import "apex-ems-integration-clean-arch/entities/web"

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
	CountSaldoAkhirOnNoRekening(kodeLKM string) (web.CalculateRepostingResult, error)
	GetTotalTrxWithTotalPokok(TglTrans web.GetListTabtransByDate) (web.GetCountWithSumTabtransTrx, error)
}
