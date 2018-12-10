package persistency_test

import (
	"github.com/marianobarragan/Twitter/src/domain"
	"github.com/marianobarragan/Twitter/src/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"strings"
	"testing"
)

type FileTweetWriterMocked struct {
	mock.Mock
}

func (fileTweetWriterMocked *FileTweetWriterMocked) Write(tweet domain.Tweet) {
	 _ = fileTweetWriterMocked.Called(tweet)
}

func (fileTweetWriterMocked *FileTweetWriterMocked) GetLastSavedTweet() domain.Tweet {
	_ = fileTweetWriterMocked.Called()
	return new(domain.TextTweet)
}

func (fileTweetWriterMocked *FileTweetWriterMocked) SetLastSavedTweet(tweet domain.Tweet) {
	_ = fileTweetWriterMocked.Called()
}

func TestPublishedTweetIsSavedToExternalResource(t *testing.T) {

	// Initialization
	var tweet *domain.TextTweet
	user := "grupooesfera"
	text := "This is my first tweet"
	tweet = domain.NewTextTweet(user, text)
	// create an instance of our test object
	tweetWriter := new(FileTweetWriterMocked)
	tweetWriter.On("Write", tweet).Return(nil, nil)
	tweetWriter.On("GetLastSavedTweet").Return(nil, nil)
	// tweetWriter = NewMemoryTweetWriter() // Mock implementation
	tweetManager := service.NewTweetManager(tweetWriter)
	tweetManager.PublishTweet(tweet)


	// Operation
	id, _ := tweetManager.PublishTweet(tweet)

	// Validation
	publishedTweet, _ := tweetManager.GetTweetById(id)
	//memoryWriter := (tweetWriter).(*persistency.MemoryTweetWriter)
	savedTweet := tweetWriter.GetLastSavedTweet()

	tweetWriter.AssertExpectations(t)
	assert.True(t, savedTweet != publishedTweet)
}

func TestCanSearchForTweetContainingText(t *testing.T) {
	// Initialization
	// var tweetWriter = new(FileTweetWriterMocked)
	tweetWriter := service.NewMemoryTweetWriter()
	tweetManager := service.NewTweetManager(tweetWriter)
	// Create and publish a tweet
	var tweet * domain.TextTweet
	user := "grupooesfera"
	text := "This is my first tweet"
	tweet = domain.NewTextTweet(user, text)
	tweetManager.PublishTweet(tweet)
	// Operation
	searchResult := make(chan domain.Tweet)
	query := "first"
	tweetManager.SearchTweetsContaining(query, searchResult)

	// Validation
	foundTweet := <-searchResult

	assert.False(t, foundTweet == nil)
	assert.True(t, strings.Contains(foundTweet.GetText(), query))
}