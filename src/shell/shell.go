package shell

import (
	"github.com/abiosoft/ishell"
	"github.com/marianobarragan/Twitter/src/domain"
	"github.com/marianobarragan/Twitter/src/service"
)

func StartShell(service *service.TweetManager) {
	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")
	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Username: ")

			username := c.ReadLine()

			c.Print("Write your tweet: ")

			tweet := c.ReadLine()

			newTweet := domain.NewTextTweet(username, tweet);
			_, err := service.PublishTweet(newTweet)

			if err != nil {
				c.Print(err)
				return
			}
			c.Print("Tweet sent\n")
			return

		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Shows a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweet, err := service.GetTweet(0)

			if err == nil {
				c.Println(tweet.PrintableTweet())
			} else {
				c.Println(err)
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showAllTweets",
		Help: "Shows all tweets",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweets := service.GetTweets()

			for _, tweet := range tweets {
				c.Printf(tweet.PrintableTweet())
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "countUserTweets",
		Help: "Shows all tweets",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Username: ")

			username := c.ReadLine()

			count := service.CountTweetsByUser(username)

			c.Println(count)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showUserTweets",
		Help: "Shows all tweets",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Username: ")

			username := c.ReadLine()

			tweets := service.GetTweetsByUser(username)

			for _, tweet := range tweets {
				c.Printf(tweet.PrintableTweet())
			}

			return
		},
	})

	shell.Run()
}
