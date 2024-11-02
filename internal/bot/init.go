package bot

import (
	"log"
	"github.com/Resolution-hash/s1-report/config"
	"github.com/Resolution-hash/s1-report/internal/api"
	"github.com/Resolution-hash/s1-report/internal/bot/message"
	"github.com/Resolution-hash/s1-report/internal/parser"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const HOUR_START int = 9
const HOUR_END int = 21

// CHAT_ID int64 = -4599247653
var (
	CHAT_ID int64 = 0
)

type BOT struct {
	TelegramAPIToken string
}

func (b *BOT) SendMessage(bot *tgbotapi.BotAPI, userID int, text string) {
}

func InitBOT(cfg *config.BOTCongig) {
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramAPIToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true
	log.Println("Bot launched!")
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			CHAT_ID = update.Message.Chat.ID
			log.Println(update.Message.Text)
			log.Println(update.Message.Chat.ID)

		}
		err = sendMessageInHour(bot, CHAT_ID, HOUR_START, HOUR_END)
		if err != nil {
			log.Println(err)
		}

	}

}

func sendMessageInHour(bot *tgbotapi.BotAPI, chatID int64, hourStart int, hourEnd int) error {
	sentHours := make(map[int]bool)

	for {
		currentTime := time.Now()
		currentHour := currentTime.Hour()
		log.Println("Текущее время:", currentTime)
		log.Println("Час:", currentHour)
		log.Println("Отправлено в этом часе:", sentHours[currentHour])
		log.Println("В пределах интервала:", currentHour >= hourStart && currentHour <= hourEnd)

		if currentHour >= hourStart && currentHour <= hourEnd && !sentHours[currentHour] {
			var messageText string

			formattedTime := currentTime.Format("15:04")
			TimeText := "<i>Статистика за <b>" + formattedTime + "</b></i>\n\n"

			messageText = TimeText + "<b>Магазины не в сети:</b>\n"
			offlineStores, err := parser.ParseOfflineStores()
			if err != nil {
				log.Println("Error of parsing", err)
				messageText += "Не удалось получить данные от SolarWinds.\n"
			} else {
				for _, store := range offlineStores {
					if store == "" {
						messageText += "🟢 Все магазины в сети\n"
						break
					} else {
						messageText += "🔴 " + store + "\n"
					}
				}
			}

			ticketCount, err := api.GetStatistics()
			if err != nil {
				log.Println("Error of parsing", err)
				messageText += "\n<b>Количество заявок на группе:</b>\nНе удалось получить данные от S1.\n\nОшибка: " + err.Error()
			} else {
				messageText += "\n<b>Количество заявок на группе:</b>\n✍️ " + strconv.FormatFloat(ticketCount, 'f', 0, 64)
			}

			err = message.SendAfterMinutes(bot, chatID, hourStart, hourEnd, messageText)
			if err != nil {
				return err
			}

			sentHours[currentHour] = true
		}

		time.Sleep(1 * time.Minute)
		log.Println("1 minute")
	}
}
