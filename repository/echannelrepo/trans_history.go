package echannelrepo

import "new-apex-api/entities"

type EchannelRepo interface {
	GetEchannelTransHistories(entities.TransHistoryRequest, entities.LimitOffsetLkmUri) ([]entities.TransHistoryResponse, error)
}
