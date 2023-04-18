package tabunganrepo

import (
	"new-apex-api/entities"
	"new-apex-api/entities/web"
)

type TabunganRepo interface {
	FindTabunganLkm(tabunganLkm string) (entities.Tabung, error)
	CreateTabung(newTabung entities.Tabung) (entities.Tabung, error)
	GetTabScGroup() ([]web.TabSCGroup, error)
	GetTabDetailInfo(KodeLkm string) (web.GetDetailLKMInfo, error)
	GetTabInfoList(limitOffset web.LimitOffsetLkmUri) ([]web.GetDetailLKMInfo, error)
	HardDeleteTabung(kodeLkm string) error
	DeleteTabung(kodeLkm string) error
	GetRekeningLKMByStatusActive() ([]string, error)
}
