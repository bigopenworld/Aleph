package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bigopenworld/discord-bot/discord"
)

func main() {
	discord.CreateAndStart()

	
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	//dg.Close()
}