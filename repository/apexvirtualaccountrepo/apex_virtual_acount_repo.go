package apexvirtualaccountrepo

import "new-apex-api/entities"

type ApexVirtualAccountRepo interface {
	CreateApexVirtualAccount(entities.CreateVirtualAccount) (entities.VirtualAccountResponse, error)
	GetListsApexVirtualAccount(entities.LimitOffsetLkmUri) ([]entities.VirtualAccountResponse, error)
	UpdateApexVirtualAccount(entities.UpdateVirtualAccount) error
	GetListsSLATransactionVirtualAccount(entities.GetListTabtrans, entities.LimitOffsetLkmUri) ([]entities.VirtualAccountTransResponse, error)
	SoftDeleteVAAccount(...string) error
}
