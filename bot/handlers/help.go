package handlers

import (
	"Keeper/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// SendHelp - отправляет список команд пользователю в виде сообщения в телеграм
func SendHelp(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg := "<b>Список комманд:</b>\n"
	msg += "/help - вывести список команд\n"
	msg += "/set - добавить сервис (vk.ru:88005553535:Qwerty123)\n"
	msg += "/get - получить логин и пароль к сервису (vk.ru)\n"
	msg += "/all - получить все сохранённые сервисы\n"
	msg += "/del - удалить сервис (vk.ru)\n"
	utils.Send(bot, update.Message.Chat.ID, msg)
}
