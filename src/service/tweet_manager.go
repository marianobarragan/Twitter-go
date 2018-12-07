package service

import (
	"github.com/marianobarragan/Twitter/src/domain"
)
import "time"
import "errors"

type TweetManager struct {
	tweets []domain.Tweet
	tweetsByUser map[string][]domain.Tweet
	idCounter int
}

func NewTweetManager() *TweetManager {
	var tweetManager = new (TweetManager)
	tweetManager.tweets = make([] domain.Tweet, 0) // Crea un slice vacío
	tweetManager.tweetsByUser = make(map [string][]domain.Tweet)
	return tweetManager
}

func (tweetManager * TweetManager) PublishTweet(tweet domain.Tweet) (id int, error error){

	if len(tweet.GetUser()) == 0 {
		error = errors.New("user is required")
		return
	}
	if len(tweet.GetText()) == 0 {
		error = errors.New("text is required")
		return
	}
	if len(tweet.GetText()) > 140 {
		error =  errors.New("text is too long")
		return
	}

	if imageTweet, ok := tweet.(*domain.ImageTweet); ok {
		if len(imageTweet.Url) == 0 {
			error = errors.New("url is required")
			return
		}
	}

	if quoteTweet, ok := tweet.(*domain.QuoteTweet); ok {
		if quoteTweet.OriginalTweet == nil {
			error = errors.New("quoted tweet is required")
			return
		}
	}

	t := time.Now()
	tweet.SetDate(&t)
	id = tweetManager.idCounter
	tweetManager.idCounter ++
	tweet.SetId(id)
	tweetManager.tweets = append(tweetManager.tweets, tweet)
	tweetManager.tweetsByUser[tweet.GetUser()] = append(tweetManager.tweetsByUser[tweet.GetUser()], tweet)
	return id, nil
}

func (tweetManager * TweetManager) GetTweet(position int) domain.Tweet{
	if position < len(tweetManager.tweets) {
		return tweetManager.tweets[position]
	}
	return nil
}

func (tweetManager * TweetManager) GetTweets() []domain.Tweet{
	return tweetManager.tweets
}

func (tweetManager * TweetManager) GetTweetById(id int) domain.Tweet{

	for _, tweet := range tweetManager.tweets {
		if tweet.GetId() == id {
			return tweet
		}
	}
	return nil
}

func (tweetManager * TweetManager) CountTweetsByUser(user string) (count int)  {

	for _, tweet := range tweetManager.tweets {
		if tweet.GetUser() == user {
			count ++
		}
	}
	return
}

func (tweetManager * TweetManager) GetTweetsByUser(username string) (userTweets []domain.Tweet){

	//for _, tweet := range tweets {
	//	if tweet.User == username {
	//		userTweets = append(userTweets, tweet)
	//	}
	//}
	userTweets = tweetManager.tweetsByUser[username]
	return
}
