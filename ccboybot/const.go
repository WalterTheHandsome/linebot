package ccboybot

import "fmt"

const (
	ROUTE_PATH = "/ccboy"
)

const (
	ENV_BOT_CHANNEL_SECRET       = "BOT_CHANNEL_SECRET"
	ENV_BOT_CHANNEL_ACCESS_TOKEN = "BOT_CHANNEL_ACCESS_TOKEN"
	ENV_BOT_USER_ID              = "BOT_CHANNEL_USER_ID"
)

const (
	MENU_LION            = "/雄獅"
	MENU_LION_TO_UX      = "/toUX"
	MENU_LION_TO_IT      = "/toIT"
	MENU_LION_ONLINE_1   = "/online1"
	MENU_LION_ONLINE_2_3 = "/online2And3"
)

const (
	EXAMPLE_RANGE         = "範圍"
	EXAMPLE_URL           = "連結"
	EXAMPLE_CONTENT       = "修改內容"
	EXAMPLE_ONLINE_DATE   = "日期"
	EXAMPLE_ONLINE_TIME   = "上線時間"
	EXAMPLE_TICKET_NUMBER = "單號"
	EXAMPLE_UPDATE_ITEM   = "更新項目"
	EXAMPLE_PR_LINK       = "PR網址"

	FIELD_RANGE       = "range"
	FIELD_URL         = "url"
	FIELD_CONTENT     = "content"
	FIELD_ONLINE_DATE = "onlineDate"
	FIELD_ONLINE_TIME = "onlineTime"
	FIELD_TICKET_NO   = "ticketNo"
	FIELD_UPDATE_ITEM = "updateItem"
	FIELD_PR_LINK     = "prLink"
)

const (
	STATE_NONE = iota
	STATE_LION_PENDING_FOR_CHOOSE
	STATE_LION_PENDING_FOR_TO_UX_INPUT
	STATE_LION_PENDING_FOR_TO_IT_INPUT
	STATE_LION_PENDING_FOR_TO_ONLINE_1
	STATE_LION_PENDING_FOR_TO_ONLINE_2_AND_3
)

func GenExampleField(from string) string {
	return fmt.Sprintf("${%s}", from)
}

// const regex
const (
	toUXRegString = `1\.(?P<` + FIELD_RANGE + `>.+)\n2\.(?P<` + FIELD_URL + `>.+)`
	toITRegString = `1\.(?P<` + FIELD_RANGE + `>.+)\n2\.(?P<` + FIELD_CONTENT + `>.+)\n3\.(?P<` + FIELD_URL + `>.+)`
)

var (
	ToUXReminderTemplate        = "請輸入下列欄位:\n1.${範圍}\n2.${連結}\n"
	ToITReminderTemplate        = "請輸入下列欄位:\n1.${範圍}\n2.${修改內容}\n3.${連結}\n"
	OnlineStep1ReminderTemplate = "請輸入下列欄位:\n1.${日期}\n2.${連結}\n"
)

// const lion
const (
	ToUXResultTemplate        = "*%s* 的切版更新到demo機了,再麻煩你有空的時候幫忙驗收,感謝\n%s"
	ToITResultTemplate        = "下方為 *%s* 的切版檔, *%s* 已更新上測試機,再麻煩了,謝謝~\n%s"
	OnlineStep1ResultTemplate = "*%s* 預計上線PB\n%s\n若沒問題將開始進行上版流程"

	LionMainMenuCarouselJsonString = `{
		"type": "carousel",
		"contents": [
			{
				"type": "bubble",
				"size": "mega",
				"header": {
					"type": "box",
					"layout": "vertical",
					"contents": [
						{
							"type": "text",
							"text": "切版檔UX確認文案",
							"color": "#ffffff",
							"size": "24px"
						},
						{
							"type": "text",
							"text": "123",
							"size": "20px",
							"color": "#007bff"
						},
						{
							"type": "text",
							"text": "123",
							"size": "20px",
							"color": "#007bff"
						}
					]
				},
				"body": {
					"type": "box",
					"layout": "vertical",
					"contents": [
						{
							"type": "text",
							"text": "${範圍}的切版更新到demo機了,再麻煩你有空的時候幫忙驗收,感謝~ ${連結}",
							"color": "#888888",
							"style": "normal",
							"offsetTop": "10px",
							"lineSpacing": "2px",
							"size": "18px",
							"weight": "regular",
							"decoration": "none",
							"position": "relative",
							"align": "start",
							"wrap": true
						}
					]
				},
				"footer": {
					"type": "box",
					"layout": "vertical",
					"contents": [
						{
							"type": "button",
							"action": {
								"type": "message",
								"label": "開始",
								"text": "/toUX"
							}
						}
					]
				},
				"styles": {
					"header": {
						"backgroundColor": "#007bff"
					}
				}
			},
			{
				"type": "bubble",
				"header": {
					"type": "box",
					"layout": "vertical",
					"contents": [
						{
							"type": "text",
							"color": "#ffffff",
							"size": "24px",
							"text": "切版檔交付IT文案"
						},
						{
							"type": "text",
							"text": "123",
							"size": "20px",
							"color": "#007bff"
						},
						{
							"type": "text",
							"text": "123",
							"size": "20px",
							"color": "#007bff"
						}
					]
				},
				"body": {
					"type": "box",
					"layout": "vertical",
					"contents": [
						{
							"type": "text",
							"text": "下方為 ${範圍} 的切版檔,${內容} 已更新上測試機,再麻煩了,謝謝~ ${連結}",
							"wrap": true,
							"size": "18px",
							"lineSpacing": "2px",
							"weight": "regular",
							"style": "normal",
							"position": "relative",
							"align": "start",
							"offsetTop": "10px",
							"color": "#888888"
						}
					]
				},
				"footer": {
					"type": "box",
					"layout": "vertical",
					"contents": [
						{
							"type": "button",
							"action": {
								"type": "message",
								"label": "開始",
								"text": "/toIT"
							}
						}
					],
					"action": {
						"type": "message",
						"label": "開始",
						"text": "/toIT"
					}
				},
				"styles": {
					"header": {
						"backgroundColor": "#007bff"
					}
				}
			},
			{
				"type": "bubble",
				"header": {
					"type": "box",
					"layout": "vertical",
					"contents": [
						{
							"type": "text",
							"text": "上線正式機流程",
							"size": "24px",
							"color": "#ffffff"
						},
						{
							"type": "text",
							"text": "step1. 確認本次上線PB",
							"size": "20px",
							"color": "#ffffff"
						},
						{
							"type": "text",
							"text": "123",
							"size": "20px",
							"color": "#007bff"
						}
					]
				},
				"body": {
					"type": "box",
					"layout": "vertical",
					"contents": [
						{
							"type": "text",
							"text": "${日期} 預計上線PB ${連結} 若沒問題將開始進行上版流程",
							"size": "18px",
							"lineSpacing": "2px",
							"color": "#888888",
							"weight": "regular",
							"style": "normal",
							"position": "relative",
							"align": "start",
							"wrap": true,
							"offsetTop": "10px"
						}
					]
				},
				"footer": {
					"type": "box",
					"layout": "vertical",
					"contents": [
						{
							"type": "button",
							"action": {
								"type": "message",
								"label": "開始",
								"text": "/online1"
							}
						}
					]
				},
				"styles": {
					"header": {
						"backgroundColor": "#007bff"
					}
				}
			},
			{
				"type": "bubble",
				"header": {
					"type": "box",
					"layout": "vertical",
					"contents": [
						{
							"type": "text",
							"text": "上線正式機流程",
							"size": "24px",
							"color": "#ffffff"
						},
						{
							"type": "text",
							"text": "step2. 確認預計上線時間",
							"size": "20px",
							"color": "#ffffff"
						},
						{
							"type": "text",
							"text": "step3. 通知房it組長上線時間、pr網址",
							"size": "20px",
							"color": "#ffffff"
						}
					]
				},
				"body": {
					"type": "box",
					"layout": "vertical",
					"contents": [
						{
							"type": "text",
							"text": "Step2",
							"size": "18px",
							"lineSpacing": "2px",
							"color": "#888888",
							"weight": "regular",
							"style": "normal",
							"position": "relative",
							"align": "start",
							"wrap": true,
							"offsetTop": "10px"
						},
						{
							"type": "text",
							"text": "${日期} 訂房夜間過版 單號: ${單號} 更新項目: ${更新項目}  以上為 ${日期} 晚間上線的內容, 請問預計安排 ${上線時間} 上線可以嗎? PR release後續更新",
							"size": "18px",
							"lineSpacing": "2px",
							"color": "#888888",
							"weight": "regular",
							"style": "normal",
							"position": "relative",
							"align": "start",
							"wrap": true,
							"offsetTop": "10px"
						},
						{
							"type": "text",
							"text": "Step3",
							"size": "18px",
							"lineSpacing": "2px",
							"color": "#888888",
							"weight": "regular",
							"style": "normal",
							"position": "relative",
							"align": "start",
							"wrap": true,
							"offsetTop": "15px"
						},
						{
							"type": "text",
							"text": "${日期} 訂房夜間過版(含PR) ${上線時間} 上線 單號: ${單號} PR網址: ${PR網址} 更新項目: ${更新項目}",
							"size": "18px",
							"lineSpacing": "2px",
							"color": "#888888",
							"weight": "regular",
							"style": "normal",
							"position": "relative",
							"align": "start",
							"wrap": true,
							"offsetTop": "10px"
						}
					]
				},
				"footer": {
					"type": "box",
					"layout": "vertical",
					"contents": [
						{
							"type": "button",
							"action": {
								"type": "message",
								"label": "開始",
								"text": "/online1"
							}
						}
					]
				},
				"styles": {
					"header": {
						"backgroundColor": "#007bff"
					}
				}
			}
		]
	}`
)
