package main
 
import (
    "log"
 
    "github.com/slack-go/slack"
)
 
func main() {
 
    OAUTH_TOKEN := "xoxb-2573512501568-2562533291841-dTPwnClKPg7QURL6Mw3lbNUM" 
    CHANNEL_ID := "C02GVF2FAG0" 
 
    api := slack.New(OAUTH_TOKEN)
    attachment := slack.Attachment{
    	Color:         "",
    	Fallback:      "",
    	CallbackID:    "",
    	ID:            0,
    	AuthorID:      "",
    	AuthorName:    "",
    	AuthorSubname: "",
    	AuthorLink:    "",
    	AuthorIcon:    "",
    	Title:         "",
    	TitleLink:     "",
    	Pretext:       "Pretext",
    	Text:          "Hello User i am ur new bot!",
    	ImageURL:      "",
    	ThumbURL:      "",
    	ServiceName:   "",
    	ServiceIcon:   "",
    	FromURL:       "",
    	OriginalURL:   "",
    	Fields:        []slack.AttachmentField{},
    	Actions:       []slack.AttachmentAction{},
    	MarkdownIn:    []string{},
    	Blocks:        slack.Blocks{},
    	Footer:        "",
    	FooterIcon:    "",
    	Ts:            "",
    }
 
    channelId, timestamp, err := api.PostMessage(
        CHANNEL_ID,
        slack.MsgOptionText("This is the main message", false),
        slack.MsgOptionAttachments(attachment),
        slack.MsgOptionAsUser(true),
    )
 
    if err != nil {
        log.Fatalf("%s\n", err)
    }
 
    log.Printf("Message successfully sent to Channel %s at %s\n", channelId, timestamp)
}