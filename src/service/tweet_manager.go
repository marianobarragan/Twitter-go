package service

import (
	"fmt"
)

var Tweet string;

func PublishTweet(tweet string){
	Tweet = tweet
}

func main() {
	fmt.Println("Hola a todos")
}
