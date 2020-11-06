package basics

import (
	"context"
	"encoding/json"
	tb "gopkg.in/tucnak/telebot.v2"
	"io/ioutil"
	"log"
	"mr-bot/pkg/constants"
	"mr-bot/pkg/datastruct"
	"mr-bot/pkg/helpers"
	"net/http"
	"os"
	"time"
)

type Bot struct {
	BotToken string
	Bot *tb.Bot
}

func (b *Bot) StartBot(ctx context.Context) error {
	select {
	case <-ctx.Done():
		log.Println("Context is done.")
	default:
		settings := tb.Settings{
			Token:       b.BotToken,
			Poller: 	 &tb.LongPoller{Timeout: time.Second},
			ParseMode:   constants.ParseMode,
		}
		bot, err := tb.NewBot(settings)
		if err != nil {
			log.Fatalf("NewBot failed: %s", err)
		}
		b.Bot = bot

		bot.Handle(tb.OnAddedToGroup, b.CheckMR)
		bot.Start()
	}
	return nil
}

func (b *Bot) CheckMR(msg *tb.Message) {
	log.Printf("Start with user %s in chat %s\n", msg.Sender.Username, msg.Chat.Title)
	token := os.Getenv("token")
	projectId := os.Getenv("project_id")
	for now := range time.Tick(constants.TimeDelta) {
		resp, err := http.DefaultClient.Get(helpers.CreateGetMRsRequest(projectId, token, now))
		if err != nil {
			log.Printf("Err: %s\n", err.Error())
			continue
		} else if resp == nil {
			log.Printf("Body is nil")
			continue
		}
		jsonResp, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		var mrs []*datastruct.MR
		json.Unmarshal(jsonResp, &mrs)
		for _, mr := range mrs {
			_, err := b.Bot.Send(msg.Chat, helpers.CreateNewMRMessage(mr))
			if err != nil {
				log.Printf("Send err: %s\n", err.Error())
			}
		}
	}
}
