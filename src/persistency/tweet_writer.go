package persistency

import "github.com/marianobarragan/Twitter/src/domain"

type TweetWriter interface {
	Write(tweet domain.Tweet)
	GetLastSavedTweet() domain.Tweet
	SetLastSavedTweet(tweet domain.Tweet)
}

type MemoryTweetWriter struct {

}

func (memoryTweetWriter *MemoryTweetWriter) Write(tweet domain.Tweet) {

}

func (memoryTweetWriter *MemoryTweetWriter) GetLastSavedTweet() domain.Tweet {
	return nil
}
func (memoryTweetWriter *MemoryTweetWriter) SetLastSavedTweet(tweet domain.Tweet) {}