package domain

import "time"

type Tweet struct {
	User string
	Text string
	Date *time.Time
	Id int
}

func NewTweet(user string, text string ) * Tweet {
	return &Tweet{user,text, nil, 0	}
}