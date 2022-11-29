package tabtransrepo

import "apex-ems-integration-clean-arch/entities/web"

type TabtransRepo interface {
	GetListTabtransInfo(TglTrans web.GetListTabtransByDate, limitOffset web.LimitOffsetLkmUri) (
		[]web.GetListTabtransInfo,
		web.GetCountWithSumTabtransTrx,
		error,
	)
	GetTotalTrxWithTotalPokok(TglTrans web.GetListTabtransByDate) (web.GetCountWithSumTabtransTrx, error)
}
