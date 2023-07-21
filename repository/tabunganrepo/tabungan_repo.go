package tabunganrepo

import (
	"new-apex-api/entities"
)

type TabunganRepo interface {
	FindTabunganLkm(tabunganLkm string) (entities.Tabung, error)
	// MultipleFindTabunganLkm(tabunganLkm ...string) ([]entities.Tabung, error)
	CreateTabung(newTabung entities.Tabung) (entities.Tabung, error)
	GetTabDetailInfo(KodeLkm string) (entities.GetDetailLKMInfo, error)
	GetTabInfoList(payload entities.GlobalFilter, limitOffset entities.LimitOffsetLkmUri) ([]entities.GetDetailLKMInfo, error)
	HardDeleteTabung(kodeLkm ...string) error
	SoftDeleteTabung(kodeLkm ...string) error
	GetRekeningLKMByStatusActive() ([]string, error)
	EditRekeningLKM(data entities.UpdateRekeningLKM) (resp entities.UpdateRekeningLKM, er error)
}
