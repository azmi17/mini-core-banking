package broadcastmessage

import (
	"errors"
	"new-apex-api/repository/databasefactory/drivers"
	"os"
)

func NewBroadcastMessage() (BroadcastMessageRepo, error) {
	currentDriver := os.Getenv("app.broadcast_driver_type")
	if currentDriver == drivers.TELEGRAM {
		return newBroadcastMsgTelegramImpl(), nil
	} else if currentDriver == drivers.WHATSAPP {
		return newBroadcastMessageWhatsappImpl(), nil
	} else {
		return nil, errors.New("unimplemented broadcast driver")
	}
}
