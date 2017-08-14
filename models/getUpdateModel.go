package models

// NewUpdates return pinter to User struc
func NewUpdates() *Updates {
	return &Updates{}
}

//Updates for update All Object telegram API
type Updates struct {
	Status      bool      `json:"ok"`
	Result      []*Update `json:"result"`
	ErrorCode   int       `json:"error_code"`
	Description string    `json:"description"`
}

//SendMessageResp for update All Object telegram API
type SendMessageResp struct {
	Status      bool     `json:"ok"`
	Result      *Message `json:"result"`
	ErrorCode   int      `json:"error_code"`
	Description string   `json:"description"`
}

//Update for update Object telegram API
type Update struct {
	UpdateID          int          `json:"update_id"`
	Message           *Message     `json:"message"`
	EditedMessage     *Message     `json:"edited_message"`
	ChannelPost       *Message     `json:"channel_post"`
	EditedChannelPost *Message     `json:"edited_channel_post"`
	InlineQuery       *InlineQuery `json:"inline_query"`
}

//InlineQuery for update Object telegram API
type InlineQuery struct {
	ID       string    `json:"id"`
	From     *User     `json:"from"`
	Location *Location `json:"location"`
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
}

//Location for update Object telegram API
type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

//Message for update Object telegram API
type Message struct {
	MessageID             int                `json:"message_id"`
	From                  *User              `json:"from"`
	Date                  int                `json:"date"`
	Chat                  *Chat              `json:"chat"`
	ForwardFrom           *User              `json:"forward_from"`
	ForwardFromChat       *Chat              `json:"forward_from_chat"`
	ForwardFromMessageID  int                `json:"forward_from_message_id"`
	ForwardDate           int                `json:"forward_date"`
	ReplyToMessage        *Message           `json:"reply_to_message"`
	EditDate              int                `json:"edit_date"`
	Text                  string             `json:"text"`
	Entities              []*MessageEntity   `json:"entities"`
	Audio                 *Audio             `json:"audio"`
	Document              *Document          `json:"document"`
	Game                  *Game              `json:"game"`
	Photo                 []*PhotoSize       `json:"photo"`
	Sticker               *Sticker           `json:"sticker"`
	Video                 *Video             `json:"video"`
	Voice                 *Voice             `json:"voice"`
	VideoNote             *VideoNote         `json:"video_note"`
	NewChatMembers        []*User            `json:"new_chat_members"`
	Caption               string             `json:"caption"`
	Contact               *Contact           `json:"contact"`
	Location              *Location          `json:"location"`
	Venue                 *Venue             `json:"venue"`
	NewChatMember         *User              `json:"new_chat_member"`
	LeftChatMember        *User              `json:"left_chat_member"`
	NewChatTitle          string             `json:"new_chat_title"`
	NewChatPhoto          []*PhotoSize       `json:"new_chat_photo"`
	DeleteChatPhoto       bool               `json:"delete_chat_photo"`
	GroupChatCreated      bool               `json:"group_chat_created"`
	SupergroupChatCreated bool               `json:"supergroup_chat_created"`
	ChannelChatCreated    bool               `json:"channel_chat_created"`
	MigrateToChatID       int                `json:"migrate_to_chat_id"`
	MigrateFromChatID     int                `json:"migrate_from_chat_id"`
	PinnedMessage         *Message           `json:"pinned_message"`
	Invoice               *Invoice           `json:"invoice"`
	SuccessfulPayment     *SuccessfulPayment `json:"successful_payment"`
}

//SuccessfulPayment telegram API
type SuccessfulPayment struct {
	Currency                string     `json:"currency"`
	TotalAmount             int        `json:"total_amount"`
	InvoicePayload          string     `json:"invoice_payload"`
	ShippingOptionID        string     `json:"shipping_option_id"`
	OrderInfo               *OrderInfo `json:"order_info"`
	TelegramPaymentChargeID string     `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeID string     `json:"provider_payment_charge_id"`
}

//OrderInfo telegram API
type OrderInfo struct {
	Name            string           `json:"name"`
	PhoneNumber     string           `json:"phone_number"`
	Email           string           `json:"email"`
	ShippingAddress *ShippingAddress `json:"shipping_address"`
}

//ShippingAddress telegram API
type ShippingAddress struct {
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	City        string `json:"city"`
	StreetLine1 string `json:"street_line1"`
	StreetLine2 string `json:"street_line2"`
	PostCode    string `json:"post_code"`
}

//Invoice telegram API
type Invoice struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	StartParameter string `json:"start_parameter"`
	Currency       string `json:"currency"`
	TotalAmount    int    `json:"total_amount"`
}

//MessageEntity telegram API
type MessageEntity struct {
	Type   string `json:"type"`
	Offset int16  `json:"offset"`
	Length int16  `json:"length"`
	URL    string `json:"url"`
	User   *User  `json:"user"`
}

//Venue telegram API
type Venue struct {
	Location     *Location `json:"location"`
	Title        string    `json:"title"`
	Address      string    `json:"address"`
	FoursquareID string    `json:"foursquare_id"`
}

//Contact telegram API
type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	UserID      int    `json:"user_id"`
}

//User telegram API
type User struct {
	ID           int    `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}

//Audio telegram API
type Audio struct {
	FileID    string `json:"file_id"`
	Duration  int    `json:"duration"`
	Performer string `json:"performer"`
	Title     string `json:"title"`
	MimeType  string `json:"mime_type"`
	FileSize  int    `json:"file_size"`
}

//Game telegram API
type Game struct {
	Title        string         `json:"title"`
	Description  string         `json:"description"`
	Photo        *PhotoSize     `json:"photo"`
	Text         string         `json:"text"`
	TextEntities *MessageEntity `json:"text_entities"`
	Animation    *Animation     `json:"animationtext"`
}

//Voice telegram API
type Voice struct {
	FileID   string `json:"file_id"`
	Duration int    `json:"duration"`
	MimeType string `json:"mime_type"`
	FileSize int    `json:"file_size"`
}

//Document telegram API
type Document struct {
	FileID   string     `json:"file_id"`
	Thumb    *PhotoSize `json:"thumb"`
	FileName string     `json:"file_name"`
	MimeType string     `json:"mime_type"`
	FileSize int        `json:"file_size"`
}

//PhotoSize telegram API
type PhotoSize struct {
	FileID   string `json:"file_id"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	FileSize int    `json:"file_size"`
}

//Video telegram API
type Video struct {
	FileID   string     `json:"file_id"`
	Width    int        `json:"width"`
	Height   int        `json:"height"`
	Duration int        `json:"duration"`
	Thumb    *PhotoSize `json:"thumb"`
	MimeType string     `json:"mime_type"`
	FileSize int        `json:"file_size"`
}

//VideoNote telegram API
type VideoNote struct {
	FileID   string     `json:"file_id"`
	Length   int        `json:"length"`
	Duration int        `json:"duration"`
	Thumb    *PhotoSize `json:"thumb"`
	FileSize int        `json:"file_size"`
}

//Sticker telegram API
type Sticker struct {
	FileID       string        `json:"file_id"`
	Width        int           `json:"width"`
	Height       int           `json:"height"`
	Thumb        *PhotoSize    `json:"thumb"`
	Emoji        string        `json:"emoji"`
	SetName      string        `json:"set_name"`
	MaskPosition *MaskPosition `json:"mask_position"`
	FileSize     int           `json:"file_size"`
}

//MaskPosition telegram API
type MaskPosition struct {
	Point  string  `json:"point"`
	XShift float64 `json:"x_shift"`
	YShift float64 `json:"y_shift"`
	Scale  float64 `json:"scale"`
}

//Animation telegram API
type Animation struct {
	FileID   string     `json:"file_id"`
	FileName string     `json:"file_name"`
	Thumb    *PhotoSize `json:"thumb"`
	MimeType string     `json:"mime_type"`
	FileSize int        `json:"file_size"`
}

//Chat telegram API
type Chat struct {
	ID                          int        `json:"id"`
	FirstName                   string     `json:"First_name"`
	LastName                    string     `json:"Last_name"`
	Username                    string     `json:"username"`
	Type                        string     `json:"Type"`
	Title                       string     `json:"title"`
	AllMembersAreAdministrators bool       `json:"all_members_are_administrators"`
	Photo                       *ChatPhoto `json:"photo"`
	Description                 string     `json:"description"`
	InviteLink                  string     `json:"invite_link"`
}

//ChatPhoto Telegram API
type ChatPhoto struct {
	SmallFileID string `json:"small_file_id"`
	BigFileID   string `json:"big_file_id"`
}
