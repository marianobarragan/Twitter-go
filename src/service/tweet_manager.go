package service

import "github.com/marianobarragan/Twitter/src/domain"
import "time"

var tweet *domain.Tweet;

func PublishTweet(tweet_p *domain.Tweet){
	tweet = tweet_p;
	t := time.Now()
	tweet.Date = &t;
}

func GetTweet() *domain.Tweet{
	return tweet
}




