package persistency

import (
	"github.com/marianobarragan/Twitter/src/domain"
	"io/ioutil"
	"os"
)

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

// FileTweetWriter
type FileTweetWriter struct {
	file *os.File
}

var SAVEFILE = "."
func NewFileTweetWriter() *FileTweetWriter {
	file, _ := os.OpenFile(
		SAVEFILE,
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)

	writer := new(FileTweetWriter)
	writer.file = file

	return writer
}

func (writer *FileTweetWriter) WriteTweet(tweet domain.Tweet) {

	go func() {
		if writer.file != nil {
			byteSlice := []byte(tweet.PrintableTweet() + "\n")
			writer.file.Write(byteSlice)
		}
	}()
}


func (fileTweetWriter FileTweetWriter) GetSavedTweets() string{
	dat, err := ioutil.ReadFile(SAVEFILE)
	if err != nil {
		panic(err)
	}
	tmpstr := string(dat) // TODO
	return tmpstr
}