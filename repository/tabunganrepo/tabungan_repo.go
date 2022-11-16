package tabunganrepo

import (
	"apex-ems-integration-clean-arch/entities"
	"apex-ems-integration-clean-arch/entities/web"
)

type TabunganRepo interface {
	CreateTabung(newTabung entities.Tabung) (entities.Tabung, error)
	HardDeleteTabung(kodeLkm string) error
	DeleteTabung(kodeLkm string) error
	GetTabScGroup() ([]web.TabSCGroup, error)
	GetTabDetailInfo(KodeLkm string) (web.GetDetailLKMInfo, error)
	GetTabInfoList(limitOffset web.LimitOffsetLkmUri) ([]web.GetDetailLKMInfo, error)
}
