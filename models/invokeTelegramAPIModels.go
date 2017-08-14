package models

// NewCallInput return pinter to User struc
func NewCallInput() *CallInput {
	return &CallInput{}
}

// SetMethod callInput
func (c *HTTPParam) SetMethod(val string) {
	c.Method = val
}

// SetAPIURL callInput
func (c *HTTPParam) SetAPIURL(val string) {
	c.APIURL = val
}

// SetAPIURI callInput
func (c *HTTPParam) SetAPIURI(val string) {
	c.APIURI = val
}

// SetChatID SendMessageInput
func (c *SendMessageInput) SetChatID(val int) {
	c.ChatID = val
}

// SetReplyToMessageID SendMessageInput
func (c *SendMessageInput) SetReplyToMessageID(val int) {
	c.ReplyToMessageID = val
}

// SetDisableNotification SendMessageInput
func (c *SendMessageInput) SetDisableNotification(val bool) {
	c.DisableNotification = val
}

// SetText SendMessageInput
func (c *SendMessageInput) SetText(val string) {
	c.Text = val
}

// SetQueryParam SendMessageInput
func (c *QueryParam) SetQueryParam(val [][]string) {
	c.Query = val
}

//CallInput struct
type CallInput struct {
	Method      string
	APIURL      string
	APIURI      string
	Offset      int
	QueryString *SendMessageInput
}

//HTTPParam struct
type HTTPParam struct {
	Method string
	APIURL string
	APIURI string
}

//QueryParam struct
type QueryParam struct {
	Query [][]string
}

//SendMessageInput struct
type SendMessageInput struct {
	ChatID                int
	Text                  string
	ParseMode             string
	DisableWebPagePreview bool
	DisableNotification   bool
	ReplyToMessageID      int
}
