package main

import (
	"encoding/json"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"net/http"
	"strconv"
)

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

func helloFunc(bot *linebot.Client, event *linebot.Event) {
	reply := "こんにちは！！"
	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(reply)).Do(); err != nil {
		log.Print(err)
	}
}

func ledOnFunc(bot *linebot.Client, event *linebot.Event) {
	target := LandryUrl + "&data=on"
	_, err := http.Get(target)
	if err != nil {
		log.Print(err)
	}
	reply := "LEDをひからせました"
	if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(reply)).Do(); err != nil {
		log.Print(err)
	}
}

func ledOffFunc(bot *linebot.Client, event *linebot.Event) {
	target := LandryUrl + "&data=off"
	_, err := http.Get(target)
	if err != nil {
		log.Print(err)
	}
	reply := "LEDをけしました"
	if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(reply)).Do(); err != nil {
		log.Print(err)
	}
}

func askLEDFunc(bot *linebot.Client, event *linebot.Event) {
	leftBtn := linebot.NewMessageAction("つける", "LEDをつける！")
	rightBtn := linebot.NewMessageAction("けす", "LEDをけす！")

	template := linebot.NewConfirmTemplate("LEDをつけますか？", leftBtn, rightBtn)
	reply := linebot.NewTemplateMessage("Sorry :(, please update your app.", template)
	if _, err := bot.ReplyMessage(event.ReplyToken, reply).Do(); err != nil {
		log.Print(err)
	}
}

func askLandryFunc(bot *linebot.Client, event *linebot.Event) {
	leftBtn := linebot.NewMessageAction("取り込む", "洗濯物を取り込む！")
	rightBtn := linebot.NewMessageAction("取り込まない", "洗濯物を取り込まない！")

	template := linebot.NewConfirmTemplate("洗濯物を取り込みますか？", leftBtn, rightBtn)
	reply := linebot.NewTemplateMessage("Sorry :(, please update your app.", template)
	if _, err := bot.ReplyMessage(event.ReplyToken, reply).Do(); err != nil {
		log.Print(err)
	}
}
func getLandryFunc(bot *linebot.Client, event *linebot.Event) {
	target := LandryUrl + "&data=landry_on"
	_, err := http.Get(target)
	if err != nil {
		log.Print(err)
	}
	reply := "洗濯物を取り込みます"
	if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(reply)).Do(); err != nil {
		log.Print(err)
	}
}

func doNotingLandryFunc(bot *linebot.Client, event *linebot.Event) {
	reply := "洗濯物を取り込みません"
	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(reply)).Do(); err != nil {
		log.Print(err)
	}
}

func getWeatherFunc(bot *linebot.Client, event *linebot.Event, location string) {
	resp, err := http.Get(WeatherUrl + "&q=" + location)
	if err != nil {
		log.Print(err)
	}
	var d weatherData
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		log.Print(err)
	}
	reply := "Location: " + d.Name + "\n" + "celsius: " + strconv.FormatFloat(d.Main.Kelvin-273, 'f', 4, 64)
	if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(reply)).Do(); err != nil {
		log.Print(err)
	}
	resp.Body.Close()
}

func askReserveTimeFunc(bot *linebot.Client, event *linebot.Event) {
	reply := "時間指定してください"
	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(reply)).Do(); err != nil {
		log.Print(err)
	}
}

func testImageMapFunc(bot *linebot.Client, event *linebot.Event) {
	if _, err := bot.ReplyMessage(
		event.ReplyToken,
		linebot.NewImagemapMessage(
			"https://salty-garden-55595.herokuapp.com/image2/",
			"Imagemap alt text",
			linebot.ImagemapBaseSize{1040, 1040},
			linebot.NewMessageImagemapAction("1時間後に回収して", linebot.ImagemapArea{260 * 0, 0, 260, 260}),
			linebot.NewMessageImagemapAction("2時間後に回収して", linebot.ImagemapArea{260 * 1, 0, 260 * 2, 260}),
			linebot.NewMessageImagemapAction("3時間後に回収して", linebot.ImagemapArea{260 * 2, 0, 260 * 3, 260}),
			linebot.NewMessageImagemapAction("4時間後に回収して", linebot.ImagemapArea{260 * 3, 0, 260 * 4, 260}),
			linebot.NewMessageImagemapAction("5時間後に回収して", linebot.ImagemapArea{260 * 0, 260, 260, 260 * 2}),
			linebot.NewMessageImagemapAction("6時間後に回収して", linebot.ImagemapArea{260 * 1, 260, 260 * 2, 260 * 2}),
			linebot.NewMessageImagemapAction("7時間後に回収して", linebot.ImagemapArea{260 * 2, 260, 260 * 3, 260 * 2}),
			linebot.NewMessageImagemapAction("8時間後に回収して", linebot.ImagemapArea{260 * 3, 260, 260 * 4, 260 * 2}),
			linebot.NewMessageImagemapAction("9時間後に回収して", linebot.ImagemapArea{260 * 0, 260 * 2, 260, 260 * 3}),
			linebot.NewMessageImagemapAction("10時間後に回収して", linebot.ImagemapArea{260 * 1, 260 * 2, 260 * 2, 260 * 3}),
			linebot.NewMessageImagemapAction("11時間後に回収して", linebot.ImagemapArea{260 * 2, 260 * 2, 260 * 3, 260 * 3}),
			linebot.NewMessageImagemapAction("12時間後に回収して", linebot.ImagemapArea{260 * 3, 260 * 2, 260 * 4, 260 * 3}),
			linebot.NewMessageImagemapAction("13時間後に回収して", linebot.ImagemapArea{260 * 0, 260 * 3, 260, 260 * 4}),
			linebot.NewMessageImagemapAction("14時間後に回収して", linebot.ImagemapArea{260 * 1, 260 * 3, 260 * 2, 260 * 4}),
			linebot.NewMessageImagemapAction("15時間後に回収して", linebot.ImagemapArea{260 * 2, 260 * 3, 260 * 3, 260 * 4}),
			linebot.NewMessageImagemapAction("16時間後に回収して", linebot.ImagemapArea{260 * 3, 260 * 3, 260 * 4, 260 * 4}),
		),
	).Do(); err != nil {
		log.Print(err)
	}
}
