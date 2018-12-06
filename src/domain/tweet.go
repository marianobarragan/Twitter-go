package domain

import "time"

type Tweet struct {
	User string
	Text string
	Date *time.Time
	Id int
}

type Stringer interface {
	String() string
}

func (tweet *Tweet) PrintableTweet() string{
	return "@" + tweet.User + ": " + tweet.Text
}

func (tweet *Tweet) String() string {
	return tweet.PrintableTweet()
}

func NewTweet(user string, text string ) * Tweet {
	return &Tweet{user,text, nil, 0	}
}