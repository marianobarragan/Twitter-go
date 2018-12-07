package service_test

import (
	"github.com/marianobarragan/Twitter/src/domain"
	"github.com/marianobarragan/Twitter/src/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func isValidTweet(t * testing.T, publishedTweet domain.Tweet,user string, text string) bool{
	if publishedTweet.GetUser() != user || publishedTweet.GetText() != text {
		t.Errorf("Expected tweet is %s: %s \n but is %s: %s", user, text, publishedTweet.GetUser(), publishedTweet.GetText())
		return false
	}
	return true
}

func TestPublishedTweetIsSaved(t *testing.T) { // importo de testing el tipo T

	// Initialization
	service := service.NewTweetManager()
	var tweet * domain.TextTweet
	user := "grupooesfera"
	text := "This is my first tweet"
	tweet = domain.NewTextTweet(user, text)
	service.PublishTweet(tweet)

	// Operation
	service.PublishTweet(tweet)

	// Validation
	publishedTweet := service.GetTweet(0)
	assert.True(t, isValidTweet(t,publishedTweet,user,text), "Tweet is valid!")
	assert.NotNil(t, publishedTweet.GetDate())
}

func TestTweetWithoutUser(t *testing.T){
	// Initialization
	service := service.NewTweetManager()
	var tweet * domain.TextTweet
	var user string;
	text := "This is my first tweet"
	tweet = domain.NewTextTweet(user, text)
	service.PublishTweet(tweet)

	// Operation
	var err error
	_ , err = service.PublishTweet(tweet)

	// Validation
	assert.True(t,err != nil, "user is required")
}

func TestTweetWithoutTextIsNotPublished(t *testing.T){
	// Initialization
	service := service.NewTweetManager()
	var tweet * domain.TextTweet
	user := "grupooesfera"
	var text string
	tweet = domain.NewTextTweet(user, text)
	service.PublishTweet(tweet)

	// Operation
	var err error
	_ , err = service.PublishTweet(tweet)

	// Validation
	assert.True(t,err != nil, "text is required")
}

func TestTweetWithoExceeding140CharactersIsNotPublished(t *testing.T){
	// Initialization
	service := service.NewTweetManager()
	var tweet * domain.TextTweet
	user := "grupooesfera"
	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum"
	tweet = domain.NewTextTweet(user, text)
	service.PublishTweet(tweet)

	// Operation
	var err error
	_ , err = service.PublishTweet(tweet)

	// Validation
	assert.True(t,err != nil, "text is too long")
}

func TestImageTweetWithNoUrl(t *testing.T){
	// Initialization
	service := service.NewTweetManager()
	var tweet * domain.ImageTweet
	user := "grupooesfera"
	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum"
	var url string
	tweet = domain.NewImageTweet(user, text, url)
	service.PublishTweet(tweet)

	// Operation
	var err error
	_ , err = service.PublishTweet(tweet)

	// Validation
	assert.True(t,err != nil, "url is null")
}

func TestQuoteTweetWithNoOriginalTweet(t *testing.T){
	// Initialization
	service := service.NewTweetManager()
	var tweet * domain.QuoteTweet
	user := "grupooesfera"
	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum"
	originalTweet := new(domain.TextTweet)
	tweet = domain.NewQuoteTweet(user, text,originalTweet )
	service.PublishTweet(tweet)

	// Operation
	var err error
	_ , err = service.PublishTweet(tweet)

	// Validation
	assert.True(t,err != nil, "original tweet is null")
}

func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T){

	// Initialization
	service := service.NewTweetManager()
	user := "user"
	text := "text"
	var firstTweet, secondTweet *domain.TextTweet // Fill tweets with data

	firstTweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, text)

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
	service := service.NewTweetManager()

	var tweet *domain.TextTweet
	var id int

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTextTweet(user, text)
	id, _ = service.PublishTweet(tweet)
	publishedTweet := service.GetTweetById( id )
	assert.True(t,isValidTweet(t, publishedTweet, user,   text ),"tweet obtained by id is valid")
}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {
	// Initialization
	service := service.NewTweetManager()
	var tweet, secondTweet, thirdTweet *domain.TextTweet
	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)
	thirdTweet = domain.NewTextTweet(anotherUser, text)
	service.PublishTweet(tweet)
	service.PublishTweet(secondTweet)
	service.PublishTweet(thirdTweet)
	// Operation
	count := service.CountTweetsByUser(user)
	// Validation
	assert.True(t, count ==2, "Expected count is 2")
}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {
	// Initialization
	service := service.NewTweetManager()
	var firstTweet, secondTweet, thirdTweet *domain.TextTweet
	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	thirdText := "This is my third tweet"
	firstTweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)
	thirdTweet = domain.NewTextTweet(anotherUser, thirdText)
	// publish the 3 tweets
	service.PublishTweet(firstTweet)
	service.PublishTweet(secondTweet)
	service.PublishTweet(thirdTweet)
	// Operation
	tweets := service.GetTweetsByUser(user)

	// Validation
	assert.True(t, len(tweets) == 2, "must return 2 tweets")
	firstPublishedTweet := tweets[0]
	secondPublishedTweet := tweets[1]

	assert.True(t,firstPublishedTweet == firstTweet, "same tweets")
	assert.True(t,secondPublishedTweet == secondTweet, "same tweets")
	assert.True(t,isValidTweet(t, firstPublishedTweet, user,   text ),"tweet obtained by username is valid")
	assert.True(t,isValidTweet(t, secondPublishedTweet, user,   secondText ),"tweet obtained by username is valid")
}
