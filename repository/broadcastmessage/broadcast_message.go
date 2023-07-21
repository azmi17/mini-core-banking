package broadcastmessage

type BroadcastMessageRepo interface {
	SendMessage(receiverID string, text string) (bool, error)
}
