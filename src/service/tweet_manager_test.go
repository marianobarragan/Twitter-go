package service_test

import (
	"github.com/marianobarragan/Twitter/src/service"
	"testing"
)

func TestPublishedTweetIsSaved(t *testing.T) { // importo de testing el tipo T
	var tweet string = "This is my tweet"
	service.PublishTweet(tweet)

	if service.Tweet != tweet {
		t.Error("Expected tweet is ", tweet)
	}
}