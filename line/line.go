package line

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/shunta27/honyakun/translate"
)

func botSetup() *linebot.Client {
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}
	return bot
}

// BotCallback function
func BotCallback(c *gin.Context) {
	bot := botSetup()

	events, err := bot.ParseRequest(c.Request)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			log.Print(err)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				replyText := translate.Translate(message.Text)
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyText)).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
}
