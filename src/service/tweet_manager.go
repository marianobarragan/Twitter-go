package service

import "github.com/marianobarragan/Twitter/src/domain"
import "time"
import "errors"

var tweet *domain.Tweet;

func PublishTweet(tweet_p *domain.Tweet) error{
	if len(tweet_p.User) == 0 {
		return errors.New("user is required")
	}
	if len(tweet_p.Text) == 0 {
		return errors.New("text is required")
	}
	if len(tweet_p.Text) > 140 {
		return errors.New("text is too long")
	}
	tweet = tweet_p
	t := time.Now()
	tweet.Date = &t
	return nil
}

func GetTweet() *domain.Tweet{
	return tweet
}




