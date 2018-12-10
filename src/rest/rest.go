package rest

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/marianobarragan/Twitter/src/domain"
	"github.com/marianobarragan/Twitter/src/service"
	"net/http"
	"strconv"
)

type GinServer struct {
	GinEngine *gin.Engine
	service   *service.TweetManager
}

func StartGinServer(service *service.TweetManager) {

	server := new(GinServer)
	server.service = service

	router := gin.Default()
	api := router.Group("api")
	{
		v1 := api.Group("v1")
		{
			v1.GET("tweets", server.getAllTweets)
			// v1.GET("tweets/position/:position", server.getTweetByPosition)
			v1.GET("tweets/id/:id", server.getTweetById)
			v1.PUT("tweets/text", server.publishTextTweet)
			v1.PUT("tweets/image", server.publishImageTweet)
			v1.PUT("tweets/quoted", server.publishQuotedTweet)
			v1.GET("tweets/user/:user", server.getTweetsByUser)
		}

	}

	router.Run(":8080")

}

func (ginServer *GinServer) getAllTweets(c *gin.Context) {
	c.JSON(200, gin.H{
		"tweets": ginServer.service.GetTweets(),
	})
}

func (ginServer *GinServer) getTweetByPosition(c *gin.Context) {

	var value int64
	var tweet domain.Tweet
	var err error

	if value, err = strconv.ParseInt(c.Param("position"), 10, 64); err != nil {

		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if tweet, err = ginServer.service.GetTweet(int(value)); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"tweet": tweet,
	})

}

func (ginServer *GinServer) getTweetById(c *gin.Context) {

	var value int64
	var tweet domain.Tweet
	var err error

	if value, err = strconv.ParseInt(c.Param("id"), 10, 64); err != nil {
		err = errors.New("invalid id")
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if tweet, err = ginServer.service.GetTweetById(int(value)); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"tweet": tweet,
	})

}

func (ginServer *GinServer) publishTextTweet(c *gin.Context){
	var request NewTweetRequest
	var err error
	if err = c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(request.User) == 0 || len(request.Text) == 0 {
		err = errors.New("missing required parameters")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	textTweet := domain.NewTextTweet(request.User,request.Text)
	ginServer.service.PublishTweet(textTweet)
	c.JSON(200, "OK")
}

func (ginServer *GinServer) publishImageTweet(c *gin.Context){
	var request NewTweetRequest
	var err error
	if err = c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(request.User) == 0 || len(request.Text) == 0 || len(request.Url) == 0 {
		err = errors.New("missing required parameters")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	imageTweet := domain.NewImageTweet(request.User,request.Text, request.Url)
	ginServer.service.PublishTweet(imageTweet)
	c.JSON(200, "OK")
}

func (ginServer *GinServer) publishQuotedTweet (c *gin.Context){
	var request NewTweetRequest
	var quotedTweet domain.Tweet
	var err error
	if err = c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if quotedTweet, err = ginServer.service.GetTweetById(request.QuotedTweetId); len(request.User) == 0 || len(request.Text) == 0 || err != nil {
		if err == nil { err = errors.New("missing required parameters") }
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	textTweet := domain.NewQuoteTweet(request.User,request.Text,quotedTweet)
	ginServer.service.PublishTweet(textTweet)
	c.JSON(200, "OK")
}

type NewTweetRequest struct {
	User string `json:"user"`
	Text string `json:"text"`
	Url string `json:"url"`
	QuotedTweetId int `json:"quotedTweetId"`
}

func (ginServer *GinServer) getTweetsByUser(c *gin.Context) {


	var tweets []domain.Tweet
	var err error



	if tweets = ginServer.service.GetTweetsByUser( c.Param("user")); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"tweet": tweets,
	})

}
