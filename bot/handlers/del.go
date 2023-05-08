package handlers

import (
	"Keeper/db"
	"Keeper/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// DeleteService - удаляет сервис из базы данных и возвращает сообщение в телеграм
func DeleteService(update tgbotapi.Update, bot *tgbotapi.BotAPI, database *db.DB) error {
	idTg := update.Message.From.ID
	serviceName := update.Message.CommandArguments()
	service, err := database.DeleteService(idTg, serviceName)
	if err != nil {
		utils.Send(bot, update.Message.Chat.ID, "Сервис "+serviceName+" не найден")
	} else {
		utils.Send(bot, update.Message.Chat.ID, "Сервис "+service.ServiceName+" удален")
	}
	return err
}
