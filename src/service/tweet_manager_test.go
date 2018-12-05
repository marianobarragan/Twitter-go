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

func TestTweetWithoutUser(t *testing.T){
	// Initialization
	var tweet * domain.Tweet
	var user string;
	text := "This is my first tweet"
	tweet = domain.NewTweet(user, text)
	service.PublishTweet(tweet)

	// Operation
	var err error
	err = service.PublishTweet(tweet)

	// Validation
	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error is required")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T){
	// Initialization
	var tweet * domain.Tweet
	user := "grupooesfera"
	var text string
	tweet = domain.NewTweet(user, text)
	service.PublishTweet(tweet)

	// Operation
	var err error
	err = service.PublishTweet(tweet)

	// Validation
	if err != nil && err.Error() != "text is required" {
		t.Error("Expected error is required")
	}
}

func TestTweetWithoExceeding140CharactersIsNotPublished(t *testing.T){
	// Initialization
	var tweet * domain.Tweet
	user := "grupooesfera"
	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum"
	tweet = domain.NewTweet(user, text)
	service.PublishTweet(tweet)

	// Operation
	var err error
	err = service.PublishTweet(tweet)

	// Validation
	if err != nil && err.Error() != "text is too long" {
		t.Error("Expected error is required")
	}
}