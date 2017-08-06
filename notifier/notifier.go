package notifier

import (
	"encoding/json"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type Notifier interface {
	Push(msg string) error
}

func NewNotifier(rawConfig string) (Notifier, error) {
	config := make(map[string]interface{})
	json.Unmarshal([]byte(rawConfig), &config)

	if config["type"] == "telegram" {
		n := &TelegramNotifier{
			ApiToken:  config["token"].(string),
			Listeners: make(map[int64]tgbotapi.User),
		}
		if err := n.Init(); err != nil {
			return nil, err
		}
		go n.Listen()
		return n, nil
	}
	return nil, fmt.Errorf("Unsupported notifier type: %s", config["type"])
}
