package message

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func SendAfterHour(bot *tgbotapi.BotAPI, chatID int64, hourStart int, hourEnd int, text string) error {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = "HTML"
	_, err := bot.Send(msg)
	if err != nil {
		return err
	}
	log.Println("Сообщение отправлено успешно!")
	return nil
}

func SendAfterMinutes(bot *tgbotapi.BotAPI, chatID int64, hourStart int, hourEnd int, text string) error {

	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = "HTML"
	_, err := bot.Send(msg)
	if err != nil {
		return err
	}
	log.Println("Сообщение отправлено успешно!")
	return nil

}

