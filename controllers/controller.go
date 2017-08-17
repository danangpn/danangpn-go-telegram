package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"github.com/danangpn/danangpn-go-telegram/models"
)

// var inp *models.CallInput

// ControllerResp struc
type ControllerResp struct{}

// NewController return pointer to UpdatesController struct
func NewController() *ControllerResp {
	return &ControllerResp{}
}

//InvokeTelegramAPI call API end point telegram
func (cr *ControllerResp) InvokeTelegramAPI(method string, apiURL string, apiURI string, query [][]string, target interface{}) error {
	u, _ := url.Parse(apiURL)
	u.Path = apiURI
	data := u.Query()
	for i := 0; i < len(query[0]); i++ {
		data.Add(query[0][i], query[1][i])
	}
	u.RawQuery = data.Encode()
	urlStr := u.String()

	client := &http.Client{}
	// r, rErr := http.NewRequest(method, urlStr, bytes.NewBufferString(data.Encode()))
	request, requestErr := http.NewRequest(method, urlStr, nil)
	if requestErr != nil {
		return requestErr
	}
	dumpRequest, _ := httputil.DumpRequestOut(request, true)
	fmt.Printf("Dump Request: %q\n", dumpRequest)
	response, responseErr := client.Do(request)
	if responseErr != nil {
		return responseErr
	}
	dumpResponse, _ := httputil.DumpResponse(response, true)
	fmt.Printf("Dump response: %q\n", dumpResponse)
	defer response.Body.Close()
	return json.NewDecoder(response.Body).Decode(target)
}

//InvokeTelegramAPIV2 vvv
func (cr *ControllerResp) InvokeTelegramAPIV2(httpPrm interface{}, queryString interface{}, target interface{}) error {
	var apiURL, apiURI, method, chatID, text, replyTo, offset string

	hp := httpPrm.(*models.HTTPParam)
	fmt.Printf("Dump in: %q\n", hp)
	apiURL = hp.APIURL
	apiURI = hp.APIURI
	method = hp.Method
	in := queryString.(*models.QueryParam).Query
	fmt.Printf("Dump in: %q\n", in)
	/*
		chatID = strconv.Itoa(in.QueryString.ChatID)
		text = in.QueryString.Text
		replyTo = strconv.Itoa(in.QueryString.ReplyToMessageID)
		offset = strconv.Itoa(in.Offset)
	//*/
	u, _ := url.Parse(apiURL)
	u.Path = apiURI
	data := u.Query()
	// data.Add(query[0][i], query[1][i])

	for i := 0; i < len(in[0]); i++ {
		data.Add(in[0][i], in[1][i])
	}

	data.Add("chat_id", chatID)
	data.Add("text", text)
	data.Add("reply_to_message_id", replyTo)
	data.Add("offset", offset)
	u.RawQuery = data.Encode()
	urlStr := u.String()

	client := &http.Client{}
	// r, rErr := http.NewRequest(method, urlStr, bytes.NewBufferString(data.Encode()))
	request, requestErr := http.NewRequest(method, urlStr, nil)
	if requestErr != nil {
		return requestErr
	}
	dumpRequest, _ := httputil.DumpRequestOut(request, true)
	fmt.Printf("Dump Request: %q\n", dumpRequest)
	response, responseErr := client.Do(request)
	if responseErr != nil {
		return responseErr
	}
	dumpResponse, _ := httputil.DumpResponse(response, true)
	fmt.Printf("Dump response: %q\n", dumpResponse)
	defer response.Body.Close()
	return json.NewDecoder(response.Body).Decode(target)
}
