package sysuserrepo

import (
	"new-apex-api/entities"
)

type SysUserRepo interface {
	GetSingleUserByUserName(userName string) (entities.ManajemenUserDataResponse, error)
	FindByUserName(userName string) (entities.SysDaftarUser, error) // login check
	GetUserID(userID int) (entities.SysDaftarUser, error)
	GetListOfUsers(payload entities.GlobalFilter, limitOffset entities.LimitOffsetLkmUri) ([]entities.ManajemenUserDataResponse, error)
	CreateSysDaftarUser(newSysUser entities.SysDaftarUser) (entities.SysDaftarUser, error)
	UpdateSysDaftarUser(updNasabah entities.SysDaftarUser) (entities.SysDaftarUser, error)
	HardDeleteSysDaftarUser(kodeLkm ...string) error
	SoftDeleteSysDaftarUser(kodeLkm ...string) error
	ResetUserPassword(user entities.SysDaftarUser) (entities.SysDaftarUser, error)
	UpdateLKMName(updNasabah entities.SysDaftarUser) error
	GetListsOtorisator() ([]entities.Otorisators, error)
}
