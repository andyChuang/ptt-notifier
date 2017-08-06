package notifier

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"fmt"
)

type TelegramNotifier struct {
	ApiToken  string
	Bot       *tgbotapi.BotAPI
	Listeners map[int64]tgbotapi.User
}

func (tn *TelegramNotifier) Init() error {
	bot, err := tgbotapi.NewBotAPI(tn.ApiToken)
	if err != nil {
		log.Panic(err)
	}
	tn.Bot = bot
	log.Printf("Authorized on account %s", bot.Self.UserName)
	return nil
}

func (tn *TelegramNotifier) Listen() error {
	log.Printf("Waiting for any recipient's registration...")
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := tn.Bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}
		chatId := update.Message.Chat.ID
		user := update.Message.From
		log.Printf("[%d(%s)] is registered by chat id %d", user.ID, user.String(), chatId)
		tn.Listeners[chatId] = *user
		tn.sendWelcomeMsg(chatId, *user)
	}
	return nil
}

func (tn *TelegramNotifier) Push(msg string) error {
	for chatId := range tn.Listeners {
		msgConfig := tgbotapi.NewMessage(chatId, msg)
		tn.Bot.Send(msgConfig)
	}
	return nil
}

func (tn *TelegramNotifier) sendWelcomeMsg(chatId int64, user tgbotapi.User) error {
	msg := tgbotapi.NewMessage(chatId, fmt.Sprintf("Hello, %d(%s)! 目前測試中，我會一直幫你搜尋Mabinogi板有\"贈送\"兩字的文章，之後的版本就可以讓你自行設定囉", user.ID, user.String()))
	tn.Bot.Send(msg)
	return nil
}
