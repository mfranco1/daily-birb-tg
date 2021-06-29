package main

import (
    "fmt"
    "log"
    "html"
    "math/rand"
    "os"
    "strconv"
    "time"

    "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
    message := createMessage()
    log.Printf("%s", message)
    chat_id, err := strconv.ParseInt(os.Getenv("TG_CHAT_ID"), 10, 64)
    chat_object := tgbotapi.NewMessage(chat_id, message)

    bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_TOKEN"))
    if err != nil {
        log.Panic(err)
    }
    bot.Debug = true
    log.Printf("Authorized on account %s", bot.Self.UserName)

    bot.Send(chat_object)
}

func init() {
    rand.Seed(time.Now().UnixNano())
}

func createMessage() string {
    return fmt.Sprintf("Time for your daily birb %s", randomBirb())
}

func randomBirb() string {
    bird_list := []string{
        "0x0001f426",
        "0x0001f424",
        "0x0001f414",
        "0x0001f986",
        "0x0001f985",
        "0x0001f9a9",
        "0x0001f425",
        "0x0001f423",
        "0x0001f989",
        "0x0001f99c",
        "0x0001f427",
        "0x0001f9a2",
        "0x0001f983",
    }
    hex_bird := bird_list[rand.Intn(len(bird_list))]
    int_bird, _ := strconv.ParseInt(hex_bird, 0, 64)
    return html.UnescapeString(string(int_bird))
}
