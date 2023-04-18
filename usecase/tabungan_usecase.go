package usecase

import (
	"new-apex-api/entities/err"
	"new-apex-api/entities/web"
	"new-apex-api/repository/tabunganrepo"
)

type TabunganUsecase interface {
	GetTabScGroup() ([]web.TabSCGroup, error)
	GetTabDetailInfo(Id string) (web.GetDetailLKMInfo, error)
	GetTabInfoList(limitOffset web.LimitOffsetLkmUri) ([]web.GetDetailLKMInfo, error)
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
		return make([]web.GetDetailLKMInfo, 0), nil
		//^ []web.GetDetailLKMInfo (untuk membuat sebuah slice kosong agar tidak return null di JSON) |err.NoRecord
	}

	return lkmTabList, nil
}
