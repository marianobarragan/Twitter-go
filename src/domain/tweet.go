package domain

import (
	"fmt"
	"time"
)

type Tweet interface {
	GetUser() string
	GetText() string
	GetDate() *time.Time
	GetId() int
	SetUser(string)
	SetText(string)
	SetDate(*time.Time)
	SetId(int)
	PrintableTweet() string
}

/************************************************************/

type TextTweet struct {
	User string
	Text string
	Date *time.Time
	Id int
}

func (tweet *TextTweet) GetUser() string {
	return tweet.User
}

func (tweet *TextTweet) GetText() string {
	return tweet.Text
}

func (tweet *TextTweet) GetDate() *time.Time {
	return tweet.Date
}

func (tweet *TextTweet) GetId() int {
	return tweet.Id
}

func (tweet *TextTweet) SetUser(user string){
	tweet.User = user
}

func (tweet *TextTweet) SetText(text string) {
	tweet.Text = text
}

func (tweet *TextTweet) SetDate(date *time.Time) {
	tweet.Date = date
}

func (tweet *TextTweet) SetId(id int) {
	tweet.Id = id
}

/************************************************************/

func (tweet *TextTweet) PrintableTweet() string {
	return "@" + tweet.GetUser() + ": " + tweet.GetText()
}

func (tweet *TextTweet) String() string {
	return tweet.PrintableTweet()
}

/************************************************************/

type ImageTweet struct {
	TextTweet
	Url string
}

func NewImageTweet(user string, text string, url string) *ImageTweet {

	return &ImageTweet{
		TextTweet: *NewTextTweet(user,text),
		Url:       url,
	}
}

func (imageTweet *ImageTweet) PrintableTweet() string {
	return "@" + imageTweet.GetUser() + ": " + imageTweet.GetText() + " - " + imageTweet.getImage()
}

func (imageTweet *ImageTweet) getImage() string {
	return imageTweet.Url
}

func (imageTweet *ImageTweet) setImage(url string) {
	imageTweet.Url = url
}

/************************************************************/

type QuoteTweet struct {
	TextTweet
	OriginalTweet Tweet
}

func NewQuoteTweet(user string, text string, quotedTweet Tweet) *QuoteTweet{
	return &QuoteTweet{
		TextTweet:       *NewTextTweet(user,text),
		OriginalTweet: quotedTweet,
	}
}

func (quoteTweet * QuoteTweet) PrintableTweet() string {
	return fmt.Sprintf(`@%s: %s - quotedTweet: %s`,quoteTweet.User,quoteTweet.Text,quoteTweet.OriginalTweet.PrintableTweet()) // "@" + quoteTweet.GetUser() + ": " + imageTweet.GetText() + " - " + imageTweet.getImage()
}

func (quoteTweet * QuoteTweet) setIdOriginalTweet(id int){
	quoteTweet.Id = id
}

func (quoteTweet * QuoteTweet) getIdOriginalTweet() int {
	return quoteTweet.Id
}

/************************************************************/

type Stringer interface {
	String() string
}

func NewTextTweet(user string, text string ) * TextTweet {
	return &TextTweet{user,text, nil, 0	}
}

