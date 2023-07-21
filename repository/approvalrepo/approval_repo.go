package approvalrepo

import "new-apex-api/entities"

type ApprovalRepo interface {
	CreateNewApproval(payload entities.Approval) (data entities.Approval, er error)
	GetApproval(token string) (data entities.ApprovalResponse, er error)
	GetListsApproval(limitOffset entities.LimitOffsetLkmUri) (data []entities.ApprovalResponse, er error)
	UpdateStatusApproval(status int, token string) (er error)
}
