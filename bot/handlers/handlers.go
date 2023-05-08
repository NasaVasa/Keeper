package handlers

import (
	"Keeper/db"
	"Keeper/models"
	"Keeper/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

// HandleMessage - обрабатывает сообщение от пользователя и вызывает соответствующую функцию для обработки команды
func HandleMessage(update tgbotapi.Update, bot *tgbotapi.BotAPI, database *db.DB) {
	log.Printf("Message %s from @%s (%s)", update.Message.Text, update.Message.From.UserName, strconv.Itoa(update.Message.From.ID))
	userModel := models.User{IdTg: update.Message.From.ID}
	err := database.AddUser(&userModel)
	if err != nil {
		log.Println(err)
		return
	} else {
		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "set":
				err = SetPassword(update, bot, database)
			case "all":
				err = GetServices(update, bot, database)
			case "get":
				err = GetService(update, bot, database)
			case "del":
				err = DeleteService(update, bot, database)
			case "start":
				fallthrough
			case "help":
				fallthrough
			default:
				SendHelp(update, bot)
			}
			if err != nil {
				log.Println(err)
			}
		} else {
			SendHelp(update, bot)
		}
	}
	go utils.DeleteMessageFromUser(bot, update.Message, 30)
}
