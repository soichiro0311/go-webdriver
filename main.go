package main

import (
	"log"
	"os"

	"github.com/nlopes/slack"
	"github.com/sclevine/agouti"
)

var (
	// slackAPIToken(BOT)を設定
	slackAPIToken = "xoxb-xxxx"
	// ファイルアップロード先チャンネルを設定
	slackChannel = "@xxxx"
)

func main() {
	driver := agouti.PhantomJS()
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage(agouti.Browser("phantomjs"))
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}

	if err := page.Navigate("https://qiita.com/"); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}
	page.FindByClass("p-home_menuItem p-1 pl-2 pl-1@s mb-1").Click()

	page.Screenshot("/Users/seike0311/repos/go-webdriver/phantomjs_qiita.png")
	if err := postImage(slackAPIToken, slackChannel); err != nil {
		panic(err)
	}
}

func postImage(slackAPIToken, slackChannel string) error {

	file, err := os.Open("/Users/seike0311/repos/go-webdriver/phantomjs_qiita.png")
	if err != nil {
		return err
	}

	api := slack.New(slackAPIToken)

	_, err = api.UploadFile(
		slack.FileUploadParameters{
			Reader:   file,
			Filename: "upload.png",
			Channels: []string{slackChannel},
		})

	return err
}
