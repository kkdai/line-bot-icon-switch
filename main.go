// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/kkdai/line-bot-sdk-go/linebot"
)

//For LINE Friends images, please refer to https://developers.line.biz/media/messaging-api/sticker_list.pdf
const (
	BrownImage string = "https://stickershop.line-scdn.net/stickershop/v1/sticker/52002734/iPhone/sticker_key@2x.png"
	ConyImage  string = "https://stickershop.line-scdn.net/stickershop/v1/sticker/52002735/iPhone/sticker_key@2x.png"
	SallyImage string = "https://stickershop.line-scdn.net/stickershop/v1/sticker/52002736/iPhone/sticker_key@2x.png"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				var sendr *linebot.Sender

				//If user already select the sender feedback, prepare user nick name and icon here.
				switch {
				case strings.EqualFold(message.Text, "Brown"):
					sendr = linebot.NewSender("Brown", BrownImage)
				case strings.EqualFold(message.Text, "Cony"):
					sendr = linebot.NewSender("Cony", ConyImage)
				case strings.EqualFold(message.Text, "Sally"):
					sendr = linebot.NewSender("Sally", SallyImage)
				default:
					//User input other than our provide range, notify user by quick reply.
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("Please select following LINE Friends to reply you: Brown, Cony and Sally.")).Do(); err != nil {
						log.Print(err)
					}
				}
				if sendr != nil {
					//Send message with switched sender.
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("Hi, this is "+message.Text+", Nice to meet you.").WithSender(sendr)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	}
}
