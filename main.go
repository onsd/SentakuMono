package main

import (
	"bytes"
	"fmt"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/nfnt/resize"
	"image/jpeg"
	"log"
	"net/http"
	"os"
	"strconv"
)

const (
	LandryUrl  = "http://cloud.nefry.studio:1880/nefrysetting/setdata?user=wag_sasa&key=c81f2e5c73bf68f39789ac96b6b79a28f31f0a265e251316173b776475d6398a"
	WeatherUrl = "http://api.openweathermap.org/data/2.5/weather?APPID=fa86d88fb6e35f48efc6c87fbb35c611"
)

func main() {

	bot, err := linebot.New(
		os.Getenv("LINE_CHANNEL_SECRET"),
		os.Getenv("LINE_CHANNEL_TOKEN"),
	)

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/image2/", imageHandler)
	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
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
					num := string([]rune(message.Text)[:1])

					if isValidDigit(num) {
						if num == "1" {
							alpha := string([]rune(message.Text)[1:2])
							if isValidDigit(alpha) {
								num = num + alpha
								reserveFunc(num, bot, event)
								return
							}
						}
						reserveFunc(num, bot, event)
						return
					}
					switch message.Text {
					case "こんにちは":
						helloFunc(bot, event)
					case "LEDをつける！":
						ledOnFunc(bot, event)
					case "LEDをけす！":
						ledOffFunc(bot, event)
					case "LED":
						askLEDFunc(bot, event)
					case "洗濯物":
						askLandryFunc(bot, event)
					case "洗濯物を取り込む！":
						getLandryFunc(bot, event)
					case "洗濯物を取り込まない！":
						doNotingLandryFunc(bot, event)
					case "天気":
						loc := "London,uk"
						getWeatherFunc(bot, event, loc)
					case "予約":
						askReserveTimeFunc(bot, event)
					case "id":
						reply := event.Source.UserID
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(reply)).Do(); err != nil {
							log.Print(err)
						}
					case "○時間後に予約します！":
						testImageMapFunc(bot, event)
					default:
						reply := "ぼくには、よくわからんどり〜"
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(reply)).Do(); err != nil {
							log.Print(err)
						}

					}
				}
			}
		}
	})
	fmt.Println("Serve at " + os.Getenv("PORT"))
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		fmt.Println(err)
	}

}

func HttpPost(uuid, date string) error {
	jsonStr := `{"uuid":"` + uuid + `","date":"` + date + `"}`

	req, err := http.NewRequest(
		"POST",
		"https://script.google.com/macros/s/AKfycbwVg-rMmqVQdQ60GX8o6Ski7ZNxS2wwSQE2qZwRut1Fo_I6jv8/exec",
		bytes.NewBuffer([]byte(jsonStr)),
	)

	if err != nil {
		return err
	}

	// Content-Type 設定
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	s := r.URL.Path[len("/image2/"):]

	size, err := strconv.Atoi(s)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid request. %s is not ID.\n", s)
		return
	}
	file, err := os.Open("2.jpg")
	if err != nil {
		log.Fatal(err)
	}
	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(uint(size), 0, img, resize.Lanczos3)
	// write new image to file
	jpeg.Encode(w, m, nil)
}

func isValidDigit(s string) bool {
	if i, err := strconv.Atoi(s); err == nil {
		if 1 <= i && i <= 16 {
			return true
		}
	}
	return false
}

func reserveFunc(num string, bot *linebot.Client, event *linebot.Event) {
	reply := num + "時間後に予約しました。"
	err := HttpPost(event.Source.UserID, num)
	if err != nil {
		log.Print(err)
	}
	if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(reply)).Do(); err != nil {
		log.Print(err)
	}
	return
}
