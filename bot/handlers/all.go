package handlers

import (
	"Keeper/db"
	"Keeper/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

// GetServices - возвращает список сервисов пользователя в виде сообщения в телеграм
func GetServices(update tgbotapi.Update, bot *tgbotapi.BotAPI, database *db.DB) error {
	idTg := update.Message.From.ID
	services, err := database.GetServices(idTg)
	if err != nil {
		return err
	}
	if len(services) == 0 {
		utils.Send(bot, update.Message.Chat.ID, "Список сервисов пуст")
		return nil
	}
	msg := "<b>Список сервисов:</b>\n"
	for current := range services {
		msg += "<b>##### " + strconv.Itoa(current+1) + " #####</b>" + "\n"
		msg += utils.String(&(services[current]))
		msg += "\n"
	}
	utils.Send(bot, update.Message.Chat.ID, msg)
	return nil
}
