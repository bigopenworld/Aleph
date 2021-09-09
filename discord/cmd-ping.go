package discord

import (
	"strconv"
	"time"

	"github.com/bigopenworld/discord-bot/lang"
	"github.com/bigopenworld/discord-bot/structure"
	"github.com/bwmarrin/discordgo"
)

type ping struct {
	name string
	minimumperm int
	cooldown time.Duration
	allow bool
}
func (c *ping) props(s *discordgo.Session, m *discordgo.MessageCreate) {
	c.name = "ping"
	c.allow = true
	c.minimumperm = discordgo.PermissionKickMembers
	c.cooldown = 5 * time.Second
	member := structure.Member{
		ID: m.Author.ID,
		DiscMember: *m.Author,
	}
	if c.cooldown != 0 {
		active, err := LCooldownGet(member, c.name, c.cooldown)
		if err != nil {
			s.ChannelMessageSend(m.Message.ChannelID, lang.EN_Cooldown_Error)
		}
		if active {
			s.ChannelMessageSend(m.Message.ChannelID, lang.EN_Cooldown)
			c.allow = false
			return
		}
		LCooldownset(member,c.name)
	}


}
func (c *ping) checkperm(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	p, err := s.UserChannelPermissions(m.Author.ID, m.ChannelID)
	if err != nil {
		return false
	}
	return p&int64(c.minimumperm) == int64(c.minimumperm) 
}
func (c *ping) run(s *discordgo.Session, m *discordgo.MessageCreate) {
	if !c.allow {
		return
	}

	//println(status)
	msgtime, _ := m.Timestamp.Parse()
	septime := int((time.Now().UnixNano() / int64(time.Millisecond)) - (msgtime.UnixNano()/ int64(time.Millisecond)))
	s.ChannelMessageSend(m.ChannelID, "pong : " + strconv.Itoa(septime) + " ms")
}