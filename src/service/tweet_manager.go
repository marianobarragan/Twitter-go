package service

import (
	"github.com/marianobarragan/Twitter/src/domain"
)
import "time"
import "errors"

type TweetManager struct {
	tweets []*domain.Tweet
	tweetsByUser map[string][]*domain.Tweet
	idCounter int
}

func NewTweetManager() *TweetManager {
	var tweetManager = new (TweetManager)
	tweetManager.tweets = make([]* domain.Tweet, 0) // Crea un slice vacío
	tweetManager.tweetsByUser = make(map [string][]*domain.Tweet)
	return tweetManager
}

func (tweetManager * TweetManager) PublishTweet(tweet_p *domain.Tweet) (id int, error error){

	if len(tweet_p.User) == 0 {
		error = errors.New("user is required")
		return
	}
	if len(tweet_p.Text) == 0 {
		error = errors.New("text is required")
		return
	}
	if len(tweet_p.Text) > 140 {
		error =  errors.New("text is too long")
		return
	}

	t := time.Now()
	tweet_p.Date = &t
	id = tweetManager.idCounter
	tweetManager.idCounter ++
	tweet_p.Id = id
	tweetManager.tweets = append(tweetManager.tweets, tweet_p)
	//if _, exists := tweetsByUser[tweet_p.User]; !exists {
	//	tweetsByUser[tweet_p.User] = make([]* domain.Tweet, 0)
	//	tweetsByUser[tweet_p.User] =
	//} else {
	//	tweetsByUser[tweet_p.User] = append(tweetsByUser[tweet_p.User], tweet_p)
	//}
	tweetManager.tweetsByUser[tweet_p.User] = append(tweetManager.tweetsByUser[tweet_p.User], tweet_p)
	return id, nil
}

func (tweetManager * TweetManager) GetTweet(position int) *domain.Tweet{
	if position < len(tweetManager.tweets) {
		return tweetManager.tweets[position]
	}
	return nil
}

func (tweetManager * TweetManager) GetTweets() []*domain.Tweet{
	return tweetManager.tweets
}

func (tweetManager * TweetManager) GetTweetById(id int) *domain.Tweet{

	for _, tweet := range tweetManager.tweets {
		if tweet.Id == id {
			return tweet
		}
	}
	return nil
}

func (tweetManager * TweetManager) CountTweetsByUser(user string) (count int)  {

	for _, tweet := range tweetManager.tweets {
		if tweet.User == user {
			count ++
		}
	}
	return
}

func (tweetManager * TweetManager) GetTweetsByUser(username string) (userTweets []*domain.Tweet){

	//for _, tweet := range tweets {
	//	if tweet.User == username {
	//		userTweets = append(userTweets, tweet)
	//	}
	//}
	userTweets = tweetManager.tweetsByUser[username];
	return
}