package main

import (
	"github.com/marianobarragan/Twitter/src/rest"
	"github.com/marianobarragan/Twitter/src/service"
	"github.com/marianobarragan/Twitter/src/shell"
)

func main() {

	service := service.NewTweetManager(nil)

	go rest.StartGinServer(service);

	shell.StartShell(service);
}