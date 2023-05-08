package handlers

import (
	"Keeper/db"
	"Keeper/models"
	"Keeper/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
)

// SetPassword - сохраняет сервис в базу данных и отправляет сообщение пользователю в телеграм о результате сохранения
func SetPassword(update tgbotapi.Update, bot *tgbotapi.BotAPI, database *db.DB) error {
	idTg := update.Message.From.ID
	commandArguments := strings.Split(update.Message.CommandArguments(), ":")
	if len(commandArguments) != 3 {
		utils.Send(bot, update.Message.Chat.ID, "Неверный формат команды")
		return nil
	}
	serviceModel := models.Service{
		IdTg:        idTg,
		ServiceName: commandArguments[0],
		Login:       commandArguments[1],
		Password:    commandArguments[2]}
	err := database.AddService(&serviceModel)
	if err != nil {
		if err.Error() == "service already exists" {
			utils.Send(bot, update.Message.Chat.ID, "У вас уже есть сервис с таким именем")
		} else {
			utils.Send(bot, update.Message.Chat.ID, "Ну удалось сохранить сервис")
		}
	} else {
		utils.Send(bot, update.Message.Chat.ID, "Сервис успешно сохранен")
	}
	return err
}
