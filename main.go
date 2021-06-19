package main

import (
	"bytes"
	"encoding/gob"
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
	discord.BotVar.Shutdown()
	// Cleanly close down the Discord session.
	//dg.Close()
}
func getRealSizeOf(v interface{}) (int, error) {
    b := new(bytes.Buffer)
    if err := gob.NewEncoder(b).Encode(v); err != nil {
        return 0, err
    }
    return b.Len(), nil
}