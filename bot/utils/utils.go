package utils

import (
	"Keeper/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
	"time"
)

// String - возвращает строку переданного сервиса с логином и паролем
func String(service *models.Service) string {
	return "Сервис: " + service.ServiceName + "\n" +
		"Логин: <code>" + service.Login + "</code>\n" +
		"Пароль: <code>" + service.Password + "</code>\n"
}

// Send - отправляет сообщение в чат и удаляет его через 30 секунд
func Send(bot *tgbotapi.BotAPI, chatId int64, message string) {
	messageConfig := tgbotapi.NewMessage(chatId, message)
	messageConfig.ParseMode = "HTML"
	sendMessage, err := bot.Send(messageConfig)
	if err != nil {
		log.Println(err)
	}
	go DeleteMessageFromBot(bot, &sendMessage, 30)
}

// DeleteMessageFromBot - удаляет сообщение прямо сейчас
func DeleteMessageFromBot(bot *tgbotapi.BotAPI, Message *tgbotapi.Message, seconds int) {
	text := Message.Text
	for i := seconds; i > 0; i-- {
		msg := tgbotapi.NewEditMessageText(Message.Chat.ID, Message.MessageID,
			"-----[ "+strconv.Itoa(i)+" ]-----\n"+text)
		_, err := bot.Send(msg)
		if err != nil {
			log.Println(err)
		}
		time.Sleep(time.Second)
	}
	msg := tgbotapi.NewDeleteMessage(Message.Chat.ID, Message.MessageID)
	_, err := bot.Send(msg)
	if err != nil {
		log.Println(err)
	}

}

// DeleteMessageFromUser - удаляет сообщение через заданное количество секунд
func DeleteMessageFromUser(bot *tgbotapi.BotAPI, Message *tgbotapi.Message, seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
	msg := tgbotapi.NewDeleteMessage(Message.Chat.ID, Message.MessageID)
	_, err := bot.Send(msg)
	if err != nil {
		log.Println(err)
	}

}
