package broadcastmessage

// type broadcastMessageRepo struct{}

// func NewBroadcastMessageRepo() BroadcastMessageRepo {
// 	return &broadcastMessageRepo{}
// }

// var (
// 	Token     string
// 	ChatId    string
// 	ParseMode string = "html"
// )

// func getUrl() string {
// 	return fmt.Sprintf("https://api.telegram.org/bot%s", Token)
// }

// func (b *broadcastMessageRepo) GetTelegramTokenBotWithChatID(tokenBot, ChatID string) {
// 	Token = tokenBot
// 	ChatId = ChatID
// }

// func (b *broadcastMessageRepo) SendMessageToTelegram(text string) (bool, error) {

// 	// Global variables
// 	var err error
// 	var response *http.Response

// 	// Send the message
// 	url := fmt.Sprintf("%s/sendMessage", getUrl())
// 	body, _ := json.Marshal(map[string]string{
// 		"chat_id":    ChatId,
// 		"text":       text,
// 		"parse_mode": ParseMode,
// 	})
// 	response, err = http.Post(
// 		url,
// 		"application/json",
// 		bytes.NewBuffer(body),
// 	)
// 	if err != nil {
// 		return false, err
// 	}

// 	// Close the request at the end
// 	defer response.Body.Close()

// 	// Body
// 	body, err = io.ReadAll(response.Body)
// 	if err != nil {
// 		return false, err
// 	}

// 	_ = glg.Info("Notifikasi '%s' terkirim", text)
// 	_ = glg.Info("Response JSON: %s", string(body))

// 	// Return
// 	return true, nil
// }
