package usecase

import (
	"fmt"
	"new-apex-api/entities"
	"new-apex-api/entities/err"
	"new-apex-api/helper"
	"new-apex-api/repository/approvalrepo"
	"new-apex-api/repository/broadcastmessage"
	"new-apex-api/repository/broadcastrepo"
	"os"
	"strconv"
	"time"
)

type ApprovalUsecase interface {
	RequestNewTokenCode(payload entities.ApprovalRequest) (er error)
	GetListsApproval(limitOffset entities.LimitOffsetLkmUri) (approvalLists []entities.ApprovalResponse, er error)
}

type approvalUsecase struct{}

func NewApprovalUsecase() ApprovalUsecase {
	return &approvalUsecase{}
}

func (a *approvalUsecase) RequestNewTokenCode(payload entities.ApprovalRequest) (er error) {

	approvalRepo, _ := approvalrepo.NewApprovalRepo()
	broadcastRepo, _ := broadcastrepo.NewBroadcastRepo()
	broadcastMessage, _ := broadcastmessage.NewBroadcastMessage()

	// # Flaging approval status
	flagingTokenStatus := os.Getenv("app.flag_token_status")
	convflagingTokenStatus, _ := strconv.Atoi(flagingTokenStatus)

	// # Flaging expired time
	expiredTokenTime := os.Getenv("app.expired_token_time")
	convExpiredTokenTime, _ := strconv.Atoi(expiredTokenTime)

	// # Time configuration in go
	currentTime := time.Now()
	now := currentTime.Format("2006-01-02 15:04:05")
	duration := int(convExpiredTokenTime)
	nowWithAdd := currentTime.Add(time.Duration(duration) * time.Minute).Format("2006-01-02 15:04:05")

	appv := entities.Approval{}
	appv.UserID = payload.UserID
	appv.OtorisatorID = payload.OtorisatorID
	appv.Description = payload.Description
	appv.Token = helper.String(6)
	appv.Status = convflagingTokenStatus
	appv.Time = now
	appv.Expired = nowWithAdd
	_, er = approvalRepo.CreateNewApproval(appv)
	if er != nil {
		return er
	}

	user, er := broadcastRepo.GetReceiverID(payload.UserID)
	if er != nil {
		return er
	}

	otorisator, er := broadcastRepo.GetReceiverID(payload.OtorisatorID)
	if er != nil {
		return er
	}

	header := "\n =================================\n"
	title := fmt.Sprintf("<b>JUDUL</b>: <i>%s</i> \U0001F4E5", "Approval Request")
	description := fmt.Sprintf("\n<b>DESKRIPSI</b>: <i>%s</i>", "Otorisasi Dari "+user.Name)
	requestType := fmt.Sprintf("\n<b>REQUEST</b>: <i>%s</i>", appv.Description)
	tokenCode := fmt.Sprintf("\n<b>TOKEN</b>: <i>%s</i>", appv.Token)
	expired := fmt.Sprintf("\n<b>EXPIRED</b>: <i>%s</i>", appv.Expired)
	footer := "\n=================================\n\nYour Helper :)\nCT Support Asisstant"
	content := header + title + description + requestType + tokenCode + expired + footer

	_, er = broadcastMessage.SendMessage(otorisator.ReceiverID, content)
	if er != nil {
		entities.PrintError("%s", er)
	}

	return nil

}

func (a *approvalUsecase) GetListsApproval(limitOffset entities.LimitOffsetLkmUri) (approvalLists []entities.ApprovalResponse, er error) {
	if limitOffset.Limit <= 0 || limitOffset.Offset < 0 {
		return approvalLists, err.BadRequest
	}

	approvalRepo, _ := approvalrepo.NewApprovalRepo()

	approvalLists, er = approvalRepo.GetListsApproval(limitOffset)
	if er != nil {
		return approvalLists, er
	}

	if len(approvalLists) == 0 {
		return make([]entities.ApprovalResponse, 0), nil // <= []entities.GetDetailLKMInfo (untuk membuat sebuah slice kosong agar tidak return null di JSON) |err.NoRecord

	}

	return approvalLists, nil
}
