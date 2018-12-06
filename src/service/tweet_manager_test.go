package service_test

import (
	"github.com/marianobarragan/Twitter/src/domain"
	"github.com/marianobarragan/Twitter/src/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func isValidTweet(t * testing.T, publishedTweet *domain.Tweet,user string, text string) bool{
	if publishedTweet.User != user || publishedTweet.Text != text {
		t.Errorf("Expected tweet is %s: %s \n but is %s: %s", user, text, publishedTweet.User, publishedTweet.Text)
		return false
	}
	return true
}

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
	publishedTweet := service.GetTweet(0)
	assert.True(t, isValidTweet(t,publishedTweet,user,text), "Tweet is valid!")
	assert.NotNil(t, publishedTweet.Date)
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
	_ , err = service.PublishTweet(tweet)

	// Validation
	//if err != nil && err.Error() != "user is required" {
	//	t.Error("Expected error is required")
	//}
	assert.True(t,err != nil, "user is required")
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
	_ , err = service.PublishTweet(tweet)

	// Validation
	//if err != nil && err.Error() != "text is required" {
	//	t.Error("Expected error is required")
	//}
	assert.True(t,err != nil, "text is required")
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
	_ , err = service.PublishTweet(tweet)

	// Validation
	//if err != nil && err.Error() != "text is too long" {
	//	t.Error("Expected error is required")
	//}
	assert.True(t,err != nil, "text is too long")
}

func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T){

	// Initialization
	service.InitializeService()
	user := "user"
	text := "text"
	var firstTweet, secondTweet * domain.Tweet // Fill tweets with data

	firstTweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, text)

	// Operation
	service.PublishTweet(firstTweet)
	service.PublishTweet(secondTweet)

	// Validation
	publishedTweets := service.GetTweets()

	assert.True(t,len(publishedTweets)==2, "Expected size is 2")
	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]
	assert.True(t,isValidTweet(t, firstPublishedTweet, user, text),"first tweet is valid")
	assert.True(t,isValidTweet(t, secondPublishedTweet, user, text),"first tweet is valid")
}

func TestCanRetrieveTweetById(t *testing.T){
	service.InitializeService()

	var tweet *domain.Tweet
	var id int

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)
	id, _ = service.PublishTweet(tweet)
	publishedTweet := service.GetTweetById( id )
	assert.True(t,isValidTweet(t, publishedTweet, user,   text ),"tweet obtained by id is valid")
}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {
	// Initialization
	service.InitializeService()
	var tweet, secondTweet, thirdTweet *domain.Tweet
	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)
	thirdTweet = domain.NewTweet(anotherUser, text)
	service.PublishTweet(tweet)
	service.PublishTweet(secondTweet)
	service.PublishTweet(thirdTweet)
	// Operation
	count := service.CountTweetsByUser(user)
	// Validation
	assert.True(t, count ==2, "Expected count is 2")
}