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

type ToIT struct {
	Range   string
	Content string
	URL     string
	Reg     *Reg
}

type OnlineStep1 struct {
	Date string
	URL  string
	Reg  *Reg
}

type OnlineStep2AndStep3 struct {
	Date        string
	TicketNo    string
	URL         string
	UpdateItem  string
	OnlineTime  string
	ProjectName string
	PRLinks     []string
	Reg         *Reg
}

func (o OnlineStep2AndStep3) GetReminder() string {
	ret := []string{
		OnlineStep2AndStep3ReminderTemplate,
		OutputSeperatorTemplate,
		"Step2",
		fmt.Sprintf(OnlineStep2ResultTemplate,
			GenExampleField(EXAMPLE_ONLINE_DATE),
			GenExampleField(EXAMPLE_TICKET_NUMBER),
			GenExampleField(EXAMPLE_UPDATE_ITEM),
			GenExampleField(EXAMPLE_ONLINE_DATE),
			GenExampleField(EXAMPLE_ONLINE_TIME)),
		"Step3",
		// OnlineStep3ResultTemplate,
	}
	return strings.Join(ret, "\n")
}

func (o *OnlineStep2AndStep3) Output() string {
	ret := []string{
		fmt.Sprintf(OnlineStep2ResultTemplate,
			o.Date,
			o.TicketNo,
			o.UpdateItem,
			o.Date,
			o.OnlineTime),
	}
	return strings.Join(ret, "\n")
}

func (o *OnlineStep2AndStep3) Parse(from string) {
	fmt.Println("from ", from)
	if o.Reg.MatchString(from) {
		o.Date = strings.TrimSpace(o.Reg.GetSubMatchStringBySubName(from, EXAMPLE_ONLINE_DATE))
		o.OnlineTime = strings.TrimSpace(o.Reg.GetSubMatchStringBySubName(from, EXAMPLE_ONLINE_TIME))

		return
	}
	o.Date = FIELD_NONE_INDICATOR
	o.OnlineTime = FIELD_NONE_INDICATOR
	o.PRLinks = []string{}
	o.ProjectName = FIELD_NONE_INDICATOR
	o.TicketNo = FIELD_NONE_INDICATOR
	o.URL = FIELD_NONE_INDICATOR
	o.UpdateItem = FIELD_NONE_INDICATOR
}

func (o *OnlineStep2AndStep3) Reset() {
	o.Date = ""
	o.OnlineTime = ""
	o.PRLinks = []string{}
	o.ProjectName = ""
	o.TicketNo = ""
	o.URL = ""
	o.UpdateItem = ""
}

func (o *OnlineStep1) GetReminder() string {
	ret := []string{
		OnlineStep1ReminderTemplate,
		OutputSeperatorTemplate,
		fmt.Sprintf(OnlineStep1ResultTemplate, GenExampleField(EXAMPLE_ONLINE_DATE), GenExampleField(EXAMPLE_URL)),
	}
	return strings.Join(ret, "\n")
}

func (o *OnlineStep1) Output() string {
	return fmt.Sprintf(OnlineStep1ResultTemplate, o.Date, o.URL)
}

func (o *OnlineStep1) Parse(from string) {
	fmt.Println("from ", from)
	if o.Reg.MatchString(from) {
		o.Date = strings.TrimSpace(o.Reg.GetSubMatchStringBySubName(from, FIELD_ONLINE_DATE))
		o.URL = strings.TrimSpace(o.Reg.GetSubMatchStringBySubName(from, FIELD_URL))
		return
	}
	o.Date = "none"
	o.URL = "none"
}

func (o *OnlineStep1) Reset() {
	o.Date = ""
	o.URL = ""
}

func (t *ToUX) GetReminder() string {
	ret := []string{
		ToUXReminderTemplate,
		OutputSeperatorTemplate,
		fmt.Sprintf(ToUXResultTemplate, GenExampleField(EXAMPLE_RANGE), GenExampleField(EXAMPLE_URL)),
	}
	return strings.Join(ret, "\n")
}

func (t *ToUX) Output() string {
	return fmt.Sprintf(ToUXResultTemplate, t.Range, t.URL)
}

func (t *ToUX) Parse(from string) {
	if t.Reg.MatchString(from) {
		t.Range = strings.TrimSpace(t.Reg.GetSubMatchStringBySubName(from, FIELD_RANGE))
		t.URL = strings.TrimSpace(t.Reg.GetSubMatchStringBySubName(from, FIELD_URL))
		return
	}
	t.Range = "none"
	t.URL = "none"
}

func (t *ToUX) Reset() {
	t.Range = ""
	t.URL = ""
}

func (t *ToIT) Output() string {
	return fmt.Sprintf(ToITResultTemplate, t.Range, t.Content, t.URL)
}

func (t *ToIT) GetReminder() string {
	ret := []string{
		ToITReminderTemplate,
		OutputSeperatorTemplate,
		fmt.Sprintf(ToITResultTemplate, GenExampleField(EXAMPLE_RANGE), GenExampleField(EXAMPLE_CONTENT), GenExampleField(EXAMPLE_URL)),
	}
	return strings.Join(ret, "\n")
}

func (t *ToIT) Parse(from string) {
	if t.Reg.MatchString(from) {
		t.Range = strings.TrimSpace(t.Reg.GetSubMatchStringBySubName(from, FIELD_RANGE))
		t.Content = strings.TrimSpace(t.Reg.GetSubMatchStringBySubName(from, FIELD_CONTENT))
		t.URL = strings.TrimSpace(t.Reg.GetSubMatchStringBySubName(from, FIELD_URL))
		return
	}
	t.Content = "none"
	t.Range = "none"
	t.URL = "none"
}

func (t *ToIT) Reset() {
	t.Content = ""
	t.Range = ""
	t.URL = ""
}

var (
	toUX                *ToUX
	toIT                *ToIT
	onlineStep1         *OnlineStep1
	onlineStep2AndStep3 *OnlineStep2AndStep3
)

func GenExampleField(field string) string {
	return fmt.Sprintf("${%s}", field)
}

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
	toUX.Reset()
	toIT.Reset()
	onlineStep1.Reset()
	onlineStep2AndStep3.Reset()

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
			case MENU_LION_ONLINE_1:
				botState = STATE_LION_PENDING_FOR_TO_ONLINE_1
				ReplyTextMessage(onlineStep1.GetReminder(), event.ReplyToken)
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
	case STATE_LION_PENDING_FOR_TO_ONLINE_1:
		if !IsCommand(msg) {
			onlineStep1.Parse(msg)
			ReplyTextMessage(onlineStep1.Output(), event.ReplyToken)
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
	onlineStep1 = new(OnlineStep1)
	onlineStep2AndStep3 = new(OnlineStep2AndStep3)

	toUX.Reg = new(Reg)
	toIT.Reg = new(Reg)
	onlineStep1.Reg = new(Reg)
	onlineStep2AndStep3.Reg = new(Reg)

	toUX.Reg.Init(toUXRegString)
	toIT.Reg.Init(toITRegString)
	onlineStep1.Reg.Init(onlineStep1RegString)
	onlineStep2AndStep3.Reg.Init(onlineStep2AndStep3RegString)
}
