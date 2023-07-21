package usecase

import (
	"new-apex-api/entities/constants"
	"new-apex-api/entities/err"
	"new-apex-api/repository/approvalrepo"
	"new-apex-api/repository/sysuserrepo"
	"time"
)

type AuthUsecase interface {
	HeaderValidation(token string, userID int) (er error)
}

type authUsecase struct{}

func NewAuthUsecase() AuthUsecase {
	return &authUsecase{}
}

func (auth *authUsecase) HeaderValidation(token string, userID int) (er error) {
	approvalRepo, _ := approvalrepo.NewApprovalRepo()
	userRepo, _ := sysuserrepo.NewSysUserRepo()

	currentTime := time.Now()
	now := currentTime.Format("2006-01-02 15:04:05")

	if (token == "") || userID == 0 {
		return err.HeaderRequired
	}

	data, er := approvalRepo.GetApproval(token)
	if er != nil {
		return er
	}

	_, er = userRepo.GetUserID(userID)
	if er != nil {
		return er
	}

	if token != data.Token {
		return err.InvalidToken
	}

	if userID != data.UserID {
		return err.UserIDDonthMatch
	}

	if now >= data.Expired {
		return err.InvalidToken
	}

	if data.Status != 1 {
		return err.InvalidToken
	}

	er = approvalRepo.UpdateStatusApproval(constants.Used, token)
	if er != nil {
		return er
	}

	return nil
}
