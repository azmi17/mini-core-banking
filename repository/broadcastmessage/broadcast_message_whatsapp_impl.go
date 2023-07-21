package broadcastmessage

func newBroadcastMessageWhatsappImpl() BroadcastMessageRepo {
	return &whatsappImpl{}
}

type whatsappImpl struct{}

func (b *whatsappImpl) SendMessage(receiverId string, text string) (bool, error) {

	// TODO :
	// implements here..

	return true, nil
}
