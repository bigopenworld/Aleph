package discord

import (
	"time"

	"github.com/bigopenworld/discord-bot/lang"
	"github.com/bigopenworld/discord-bot/structure"
	"github.com/bwmarrin/discordgo"
)

type help struct {
	name string
	minimumperm int
	cooldown time.Duration
	allow bool
}
func (c *help) props(s *discordgo.Session, m *discordgo.MessageCreate) {
	c.name = "help"
	c.allow = true
	c.minimumperm = discordgo.PermissionKickMembers
	c.cooldown = 2 * time.Second
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
func (c *help) checkperm(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	p, err := s.UserChannelPermissions(m.Author.ID, m.ChannelID)
	if err != nil {
		return false
	}
	return p&int64(c.minimumperm) == int64(c.minimumperm) 
}
func (c *help) run(s *discordgo.Session, m *discordgo.MessageCreate) {
	if !c.allow {
		return
	}

	//println(status)
	s.ChannelMessageSend(m.ChannelID, lang.EN_HelpMsg)
}