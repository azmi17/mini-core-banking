package broadcastrepo

import "new-apex-api/entities"

type BroadcastRepo interface {
	GetReceiverID(userID int) (data entities.Broadcast, er error)
}
