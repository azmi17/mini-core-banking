package usecase

import (
	"apex-ems-integration-clean-arch/entities/err"
	"apex-ems-integration-clean-arch/entities/web"
	"apex-ems-integration-clean-arch/repository/sysuserrepo"
	"apex-ems-integration-clean-arch/repository/tabunganrepo"
)

type TabunganUsecase interface {
	GetTabScGroup() ([]web.TabSCGroup, error)
	GetTabDetailInfo(Id string) (web.GetDetailLKMInfo, error)
	GetTabInfoList(limitOffset web.LimitOffsetLkmUri) ([]web.GetDetailLKMInfo, error)
	GetListRoutingRekInduk() ([]web.RoutingRekIndukData, error)
	CreateSysApexRoutingRekInduk(web.SaveRoutingRekInduk) (web.RoutingRekIndukData, error)
	UpdateSysApexRoutingRekInduk(web.SaveRoutingRekInduk) (web.RoutingRekIndukData, error)
	DeleteSysApexRoutingRekInduk(kodeLkm string) error
}

type tabunganUsecase struct{}

func NewTabunganUsecase() TabunganUsecase {
	return &tabunganUsecase{}
}

func (t *tabunganUsecase) GetTabScGroup() (detailTabScGroup []web.TabSCGroup, er error) {
	tabunganRepo, _ := tabunganrepo.NewTabunganRepo()

	detailTabScGroup, er = tabunganRepo.GetTabScGroup()
	if er != nil {
		return detailTabScGroup, er
	}

	if len(detailTabScGroup) == 0 {
		return detailTabScGroup, err.NoRecord
	}

	return detailTabScGroup, nil
}

func (t *tabunganUsecase) GetTabDetailInfo(Id string) (detailTabInfo web.GetDetailLKMInfo, er error) {
	tabunganRepo, _ := tabunganrepo.NewTabunganRepo()

	if detailTabInfo, er = tabunganRepo.GetTabDetailInfo(Id); er != nil {
		return detailTabInfo, er
	}

	return detailTabInfo, nil
}

func (t *tabunganUsecase) GetTabInfoList(limitOffset web.LimitOffsetLkmUri) (lkmTabList []web.GetDetailLKMInfo, er error) {
	if limitOffset.Limit <= 0 || limitOffset.Offset < 0 {
		return lkmTabList, err.BadRequest
	}

	tabunganRepo, _ := tabunganrepo.NewTabunganRepo()

	lkmTabList, er = tabunganRepo.GetTabInfoList(limitOffset)
	if er != nil {
		return lkmTabList, er
	}

	if len(lkmTabList) == 0 {
		return make([]web.GetDetailLKMInfo, 0), nil // []web.GetDetailLKMInfo (untuk membuat sebuah slice kosong agar tidak return null di JSON) |err.NoRecord
	}

	return lkmTabList, nil
}

func (t *tabunganUsecase) GetListRoutingRekInduk() (routingList []web.RoutingRekIndukData, er error) {
	sysApexRepo, _ := sysuserrepo.NewSysUserRepo()

	routingList, er = sysApexRepo.GetListSysApexRoutingRekInduk()
	if er != nil {
		return routingList, er
	}

	if len(routingList) == 0 {
		return routingList, err.NoRecord
	}

	return routingList, nil
}

func (t *tabunganUsecase) CreateSysApexRoutingRekInduk(payload web.SaveRoutingRekInduk) (data web.RoutingRekIndukData, er error) {
	sysApexRepo, _ := sysuserrepo.NewSysUserRepo()

	data.KodeLkm = payload.KodeLkm
	data.NorekInduk = payload.NorekInduk

	if data, er = sysApexRepo.CreateSysApexRoutingRekInduk(data.KodeLkm, data.NorekInduk); er != nil {
		return data, er
	}

	return
}

func (t *tabunganUsecase) UpdateSysApexRoutingRekInduk(payload web.SaveRoutingRekInduk) (data web.RoutingRekIndukData, er error) {
	sysApexRepo, _ := sysuserrepo.NewSysUserRepo()

	data.KodeLkm = payload.KodeLkm
	data.NorekInduk = payload.NorekInduk

	if data, er = sysApexRepo.UpdateSysApexRoutingRekInduk(data.KodeLkm, data.NorekInduk, payload.KodeLkmTarget); er != nil {
		return data, er
	}

	return
}

func (t *tabunganUsecase) DeleteSysApexRoutingRekInduk(kodeLkm string) (er error) {
	sysApexRepo, _ := sysuserrepo.NewSysUserRepo()

	if er = sysApexRepo.DeleteSysApexRoutingRekInduk(kodeLkm); er != nil {
		return er
	}

	return nil

}
