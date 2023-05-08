package handlers

import (
	"Keeper/db"
	"Keeper/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// GetService - возвращает сервис по имени из базы данных и отправляет его пользователю в виде сообщения в телеграм
func GetService(update tgbotapi.Update, bot *tgbotapi.BotAPI, database *db.DB) error {
	idTg := update.Message.From.ID
	serviceName := update.Message.CommandArguments()
	service, err := database.GetService(idTg, serviceName)
	if err != nil {
		utils.Send(bot, update.Message.Chat.ID, "Сервис "+serviceName+" не найден")
	} else {
		utils.Send(bot, update.Message.Chat.ID, utils.String(service))
	}

	return err
}
