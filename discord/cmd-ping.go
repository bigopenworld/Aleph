package discord

import (
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

type ping struct {
	minimumperm int
}
func (c *ping) checkperm(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	return m.Member.Permissions&int64(c.minimumperm) == int64(c.minimumperm) 
}
func (c *ping) run(s *discordgo.Session, m *discordgo.MessageCreate) {

	//println(status)
	msgtime, _ := m.Timestamp.Parse()
	septime := int((time.Now().UnixNano() / int64(time.Millisecond)) - (msgtime.UnixNano()/ int64(time.Millisecond)))
	s.ChannelMessageSend(m.ChannelID, "pong : " + strconv.Itoa(septime) + " ms")
}