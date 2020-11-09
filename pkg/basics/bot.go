package basics

import (
	"context"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"mr-bot/pkg/constants"
	"mr-bot/pkg/helpers"
	"net/http"
	"os"
	"strconv"
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
	token := os.Getenv("token")
	projectId := os.Getenv("project_id")
	log.Printf("Start with user %s in chat %s for project %s\n", msg.Sender.Username, msg.Chat.Title, projectId)
	for range time.Tick(constants.TimeDelta) {
		mrs := helpers.GetMRsFromResponse(http.DefaultClient.Get(helpers.CreateGetMRsRequest(projectId, token)))
		log.Printf("Have %d mrs\n", len(mrs))
		for _, mr := range mrs {
			if _, ok := constants.GitlabToTg[mr.Author.Username]; !ok {
				log.Printf("Author %s is not in the accepted list\n", mr.Author.Username)
				continue
			}

			comments := helpers.GetCommentsFromResponse(
				http.DefaultClient.Get(helpers.CreateGetCommentsRequest(projectId, token, strconv.Itoa(mr.Iid))),
			)
			log.Printf("Have %d comments\n", len(comments))

			if !helpers.IsAssigned(comments) {
				duty := helpers.GetDuty(constants.GitlabToTg[mr.Author.Username])
				log.Println("Duty: " + duty)

				_, err := http.DefaultClient.Post(
					helpers.CreateAddCommentRequest(projectId, token, strconv.Itoa(mr.Iid), constants.TgToGitlab[duty]),
					"", nil,
				)
				if err != nil {
					log.Printf("Can't create comment: %s\n", err.Error())
					continue
				}

				_, err = b.Bot.Send(msg.Chat, helpers.CreateNewMRMessage(mr, duty))
				if err != nil {
					log.Printf("Send err: %s\n", err.Error())
				}
			}
		}
	}
}
