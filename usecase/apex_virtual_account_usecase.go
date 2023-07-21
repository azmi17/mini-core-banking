package usecase

import (
	"new-apex-api/entities"
	"new-apex-api/entities/err"
	"new-apex-api/repository/apexvirtualaccountrepo"
)

type ApexVirtualAccountUsecase interface {
	CreateApexVirtualAccount(entities.CreateVirtualAccount) (entities.VirtualAccountResponse, error)
	GetListsApexSLAVirtualAccount(limitOffset entities.LimitOffsetLkmUri) (lists []entities.VirtualAccountResponse, er error)
	UpdateApexVirtualAccount(entities.UpdateVirtualAccount) error
	GetListsSLATransactionVirtualAccount(entities.GetListTabtrans, entities.LimitOffsetLkmUri) ([]entities.VirtualAccountTransResponse, error)
	SoftDeleteVAAccount([]string) error
}

type apexVirtualAccountUsecase struct{}

func NewApexVirtualAccountUsecase() ApexVirtualAccountUsecase {
	return &apexVirtualAccountUsecase{}
}

func (va *apexVirtualAccountUsecase) CreateApexVirtualAccount(payload entities.CreateVirtualAccount) (resp entities.VirtualAccountResponse, er error) {
	vaRepo, _ := apexvirtualaccountrepo.NewApexVirtualAccountRepo()

	data, er := vaRepo.CreateApexVirtualAccount(payload)
	if er != nil {
		return resp, er
	}

	resp = data

	return resp, nil

}

func (va *apexVirtualAccountUsecase) GetListsApexSLAVirtualAccount(limitOffset entities.LimitOffsetLkmUri) (lists []entities.VirtualAccountResponse, er error) {
	if limitOffset.Limit <= 0 || limitOffset.Offset < 0 {
		return lists, err.BadRequest
	}

	vaRepo, _ := apexvirtualaccountrepo.NewApexVirtualAccountRepo()

	lists, er = vaRepo.GetListsApexVirtualAccount(limitOffset)
	if er != nil {
		return lists, er
	}

	if len(lists) == 0 {
		return make([]entities.VirtualAccountResponse, 0), nil

	}

	return lists, nil
}

func (va *apexVirtualAccountUsecase) UpdateApexVirtualAccount(payload entities.UpdateVirtualAccount) (er error) {
	vaRepo, _ := apexvirtualaccountrepo.NewApexVirtualAccountRepo()

	if er = vaRepo.UpdateApexVirtualAccount(payload); er != nil {
		return er
	}

	return
}

func (va *apexVirtualAccountUsecase) GetListsSLATransactionVirtualAccount(payload entities.GetListTabtrans, limitOffset entities.LimitOffsetLkmUri) (lists []entities.VirtualAccountTransResponse, er error) {
	if limitOffset.Limit <= 0 || limitOffset.Offset < 0 {
		return lists, err.BadRequest
	}

	vaRepo, _ := apexvirtualaccountrepo.NewApexVirtualAccountRepo()

	lists, er = vaRepo.GetListsSLATransactionVirtualAccount(payload, limitOffset)
	if er != nil {
		return lists, er
	}

	if len(lists) == 0 {
		return make([]entities.VirtualAccountTransResponse, 0), nil

	}

	return lists, nil
}

func (va *apexVirtualAccountUsecase) SoftDeleteVAAccount(vaNumber []string) (er error) {
	vaRepo, _ := apexvirtualaccountrepo.NewApexVirtualAccountRepo()

	if er = vaRepo.SoftDeleteVAAccount(vaNumber...); er != nil {
		return er
	}

	return nil

}
