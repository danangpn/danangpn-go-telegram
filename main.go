package main

import (
	"fmt"
	"net/http"
	"telegram/controllers"
	"telegram/models"
	"time"
)

var upd *models.Updates
var msg *models.SendMessageResp
var inp *models.CallInput
var smi *models.SendMessageInput
var ctrl *controllers.ControllerInvokeTelegramAPI

func main() {
	/*http.HandleFunc("/", run)
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)

	}//*/
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)

}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
func run(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello World by Chee Yeo 2015")
	var Offset int
	URL := "https://api.telegram.org/"
	Token := "bot364404824:AAHzvBLmkQqvSiZBsyo5eTbEk6mvoH6sa8w"
	Method := "POST"
	Offset = 0
	for {
		ctrl.GetUpdate(URL, Method, Token, Offset, &upd)
		fmt.Printf("\nResult getUpdate: %+v\n", upd)
		if len(upd.Result) != 0 {
			for i := 0; i < len(upd.Result); i++ {
				ChatID := upd.Result[i].Message.Chat.ID
				ReplyToMsgID := upd.Result[i].Message.MessageID
				Text := "Hi, " + upd.Result[i].Message.From.FirstName + "\nTerimakasih telah menghubungi kami, permintaan Anda sedang kami proses.\n(Msg: " + upd.Result[i].Message.Text + ")"
				ctrl.SendMessage(URL, Method, Token, ChatID, Text, ReplyToMsgID, &msg)
				Offset = upd.Result[i].UpdateID + 1
			}
		}
		time.Sleep(5 * time.Second)
	}
}
