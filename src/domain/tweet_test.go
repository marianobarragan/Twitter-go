package domain_test

import (
	"github.com/marianobarragan/Twitter/src/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTextTweetPrintsUserAndText(t *testing.T) {

	// Initialization
	tweet := domain.NewTextTweet("grupoesfera", "This is my tweet")

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := "@grupoesfera: This is my tweet"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}

func TestCanGetAStringFromATweet(t *testing.T) {

	// Initialization
	tweet := domain.NewTextTweet("grupoesfera", "This is my tweet")

	// Operation
	text := tweet.String()

	// Validation
	expectedText := "@grupoesfera: This is my tweet"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}

func TestImageTweetPrintsUserTextAndImageURL(t *testing.T) {

	// Initialization
	tweet := domain.NewImageTweet("grupoesfera", "This is my image", "http://www.grupoesfera.com.ar/common/img/grupoesfera.png")
	// Operation
	text := tweet.PrintableTweet()
	// Validation
	expectedText := "@grupoesfera: This is my image - http://www.grupoesfera.com.ar/common/img/grupoesfera.png"
	assert.True(t, expectedText == text)

}

func TestQuoteTweetPrintsUserTextAndQuotedTweet(t *testing.T) {
	// Initialization
	quotedTweet := domain.NewTextTweet("grupoesfera", "This is my tweet")
	// Operation
	tweet := domain.NewQuoteTweet("nick", "Awesome", quotedTweet)
	text := tweet.PrintableTweet()
	// Validation
	expectedText := `@nick: Awesome - quotedTweet: @grupoesfera: This is my tweet`
	assert.True(t, expectedText == text)
}