package controllers

import (
	"strconv"
	"github.com/danangpn/danangpn-go-telegram/models"
)

// ControllerInvokeTelegramAPI struc
type ControllerInvokeTelegramAPI struct{}

// //CallInput extract CallInput models
// type CallInput struct {
// 	*models.CallInput
// }

// //SendMessageInput extract CallInput models
// type SendMessageInput struct {
// 	*models.SendMessageInput
// }

var upd *models.Updates
var msg *models.SendMessageResp
var ctrl *ControllerResp

//SendMessage function
// func (cita *ControllerInvokeTelegramAPI) SendMessage(apiURL string, botToken string, chatID int, text string, replyToMsgID int) error {
func (cita *ControllerInvokeTelegramAPI) SendMessage(apiURL string, method string, botToken string, chatID int, text string, replyToMsgID int, target interface{}) error {
	apiURI := botToken + "/sendMessage"
	hp := new(models.HTTPParam)
	query := [][]string{{"chat_id", "reply_to_message_id", "text"}, {strconv.Itoa(chatID), strconv.Itoa(replyToMsgID), text}}
	inParam := new(models.QueryParam)
	inParam.SetQueryParam(query)

	hp.SetAPIURL(apiURL)
	hp.SetAPIURI(apiURI)
	if method != "" {
		hp.SetMethod(method)
	} else {
		hp.SetMethod("POST")
	}
	// resl := ctrl.InvokeTelegramAPIV2(hp, inParam, &msg)

	return ctrl.InvokeTelegramAPIV2(hp, inParam, target)

}

//GetUpdate function
func (cita *ControllerInvokeTelegramAPI) GetUpdate(apiURL string, method string, botToken string, offset int, target interface{}) error {
	// query := [][]string{{"offset"}, {offset}}
	hp := new(models.HTTPParam)
	apiURI := botToken + "/getUpdates"
	hp.SetAPIURL(apiURL)
	hp.SetAPIURI(apiURI)
	queryUpd := [][]string{{"offset"}, {strconv.Itoa(offset)}}
	inParam := new(models.QueryParam)
	inParam.SetQueryParam(queryUpd)
	if method != "" {
		hp.SetMethod(method)
	} else {
		hp.SetMethod("POST")
	}
	return ctrl.InvokeTelegramAPIV2(hp, inParam, target)
} //*/
