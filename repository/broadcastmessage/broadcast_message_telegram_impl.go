package broadcastmessage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/kpango/glg"
)

func newBroadcastMsgTelegramImpl() BroadcastMessageRepo {
	return &telegramImpl{}
}

type telegramImpl struct{}

func (b *telegramImpl) SendMessage(receiverID string, text string) (bool, error) {
	var err error
	var response *http.Response

	Token := os.Getenv("app.bot_telegram_token")
	parseMode := os.Getenv("app.telegram_parse_mode")
	baseUrl := fmt.Sprintf("https://api.telegram.org/bot%s", Token)

	// Send the message
	url := fmt.Sprintf("%s/sendMessage", baseUrl)
	body, _ := json.Marshal(map[string]string{
		"chat_id":    receiverID,
		"text":       text,
		"parse_mode": parseMode,
	})

	response, err = http.Post(
		url,
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return false, err
	}
	defer response.Body.Close() // Close the request at the end

	// Body
	body, err = io.ReadAll(response.Body)
	if err != nil {
		return false, err
	}

	_ = glg.Info("Notifikasi '%s' terkirim", text)
	_ = glg.Info("Response JSON: %s", string(body))

	return true, nil
}
