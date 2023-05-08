/*
	Keeper - Telegram Bot for saving passwords
	Author: Vasiliy Grachev, 2023
*/

package main

import (
	"Keeper/db"
	"Keeper/handlers"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
)

func main() {
	// Создание соединения с базой данных PostgreSQL
	database := db.GetDB()
	defer func(database *db.DB) {
		err := database.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(database)
	log.Println("Connected to database")

	// Создание таблиц в базе данных
	err := database.Migrate()
	if err != nil {
		log.Fatal(err)
	}

	// Создание соединения с Telegram Bot API
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Запуск бота
	log.Println("Starting bot...")
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}
		handlers.HandleMessage(update, bot, database)
	}
}
