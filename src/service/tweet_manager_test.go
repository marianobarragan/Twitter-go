package service_test

import (
	"github.com/marianobarragan/Twitter/src/domain"
	"github.com/marianobarragan/Twitter/src/service"
	"testing"
)

func TestPublishedTweetIsSaved(t *testing.T) { // importo de testing el tipo T

	// Initialization
	var tweet * domain.Tweet
	user := "grupooesfera"
	text := "This is my first tweet"
	tweet = domain.NewTweet(user, text)
	service.PublishTweet(tweet)

	// Operation
	service.PublishTweet(tweet)

	// Validation
	publishedTweet := service.GetTweet()
	if publishedTweet.User != user || publishedTweet.Text != text {
		t.Errorf("Expected tweet is %s: %s \n but is %s: %s", user, tweet, publishedTweet.User, publishedTweet.Text)
	}

	if publishedTweet.Date == nil {
		t.Error("Expected date can't be nil")
	}
}