package main

import (
    "log"
    "html"
    "math/rand"
    "os"
    "strconv"
    "time"

    "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
    bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_TOKEN"))
    if err != nil {
        log.Panic(err)
    }

    bot.Debug = true
    log.Printf("Authorized on account %s", bot.Self.UserName)
    log.Printf("%s", randomBirb())
}

func init() {
    rand.Seed(time.Now().UnixNano())
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
