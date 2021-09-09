package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bigopenworld/discord-bot/cmd"
	"github.com/bigopenworld/discord-bot/discord"
)

func main() {
	discord.CreateAndStart()
	fmt.Println(cmd.NewFlag(cmd.INFO),"Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	discord.BotVar.Shutdown()
	// Cleanly close down the Discord session.
	//dg.Close()
}