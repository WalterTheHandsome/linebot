package ccboybot

import (
	"fmt"
	"log"
	"strings"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

type ToUX struct {
	Range string
	URL   string
	Reg   *Reg
}

func (t *ToUX) GetReminder() string {
	ret := []string{
		ToUXReminderTemplate,
		"======= 輸出結果 =======",
		fmt.Sprintf(ToUXResultTemplate, "${範圍}", "${連結}"),
	}
	return strings.Join(ret, "\n")
}

func (t *ToUX) Output() string {
	return fmt.Sprintf(ToUXResultTemplate, t.Range, t.URL)
}

func (t *ToUX) Parse(from string) {
	if t.Reg.MatchString(from) {
		t.Range = strings.TrimSpace(t.Reg.GetSubMatchStringBySubName(from, "range"))
		t.URL = strings.TrimSpace(t.Reg.GetSubMatchStringBySubName(from, "url"))
		return
	}

	t.Range = "none"
	t.URL = "none"
}

type ToIT struct {
	Range   string
	Content string
	URL     string
	Reg     *Reg
}

func (t *ToIT) Output() string {
	return fmt.Sprintf(ToITResultTemplate, t.Range, t.Content, t.URL)
}

func (t *ToIT) GetReminder() string {
	ret := []string{
		ToITReminderTemplate,
		"======= 輸出結果 =======",
		fmt.Sprintf(ToITResultTemplate, "${範圍}", "${修改內容}", "${連結}"),
	}
	return strings.Join(ret, "\n")
}

func (t *ToIT) Parse(from string) {
	if t.Reg.MatchString(from) {
		t.Range = strings.TrimSpace(t.Reg.GetSubMatchStringBySubName(from, "range"))
		t.Content = strings.TrimSpace(t.Reg.GetSubMatchStringBySubName(from, "content"))
		t.URL = strings.TrimSpace(t.Reg.GetSubMatchStringBySubName(from, "url"))
		return
	}
	t.Content = "none"
	t.Range = "none"
	t.URL = "none"
}

var (
	toUX *ToUX
	toIT *ToIT
)

func ReplyLionCarousel(replyToken string) {
	contents, err := messaging_api.UnmarshalFlexContainer([]byte(LionMainMenuCarouselJsonString))
	if err != nil {
		log.Println("parse carousel err", err)
	}
	if _, err := bot.ReplyMessage(
		&messaging_api.ReplyMessageRequest{
			ReplyToken: replyToken,
			Messages: []messaging_api.MessageInterface{
				&messaging_api.FlexMessage{
					AltText:  "Flex message alt text",
					Contents: contents,
				},
			},
		},
	); err != nil {
		log.Println("reply message err", err)
	}
}

func reset() {
	log.Println("reset state and var")
	botState = STATE_NONE
	toUX = new(ToUX)
	toUX.Reg = new(Reg)
	toUX.Reg.Init(toUXRegString)
	toIT = new(ToIT)
	toIT.Reg = new(Reg)
	toIT.Reg.Init(toITRegString)
}

func HandleMenuLion(message webhook.TextMessageContent, event webhook.MessageEvent) {
	msg := message.Text

	switch botState {
	case STATE_LION_PENDING_FOR_CHOOSE:
		if IsCommand(msg) {
			switch msg {
			case MENU_LION_TO_UX:
				botState = STATE_LION_PENDING_FOR_TO_UX_INPUT
				ReplyTextMessage(toUX.GetReminder(), event.ReplyToken)
			case MENU_LION_TO_IT:
				botState = STATE_LION_PENDING_FOR_TO_IT_INPUT
				ReplyTextMessage(toIT.GetReminder(), event.ReplyToken)
			}
			return
		}
	case STATE_LION_PENDING_FOR_TO_UX_INPUT:
		if !IsCommand(msg) {
			toUX.Parse(msg)
			ReplyTextMessage(toUX.Output(), event.ReplyToken)
			reset()
			return
		}
	case STATE_LION_PENDING_FOR_TO_IT_INPUT:
		if !IsCommand(msg) {
			toIT.Parse(msg)
			ReplyTextMessage(toIT.Output(), event.ReplyToken)
			reset()
			return
		}
	default:
		reset()
	}
}

func init() {
	toUX = new(ToUX)
	toIT = new(ToIT)
	toUX.Reg = new(Reg)
	toIT.Reg = new(Reg)
	toUX.Reg.Init(toUXRegString)
	toIT.Reg.Init(toITRegString)
}
