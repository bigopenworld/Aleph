package discord

import (
	"log"

	"github.com/bigopenworld/discord-bot/config"
	"github.com/bigopenworld/discord-bot/lang"
	"github.com/bigopenworld/discord-bot/structure"
	"github.com/bwmarrin/discordgo"
)

type Cmd interface {
	props(s *discordgo.Session, m *discordgo.MessageCreate)
	checkperm(s *discordgo.Session, m *discordgo.MessageCreate) bool
	run(s *discordgo.Session, m *discordgo.MessageCreate)
}

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	guildobj := structure.NewGuild(m.GuildID)
	status := GuildCreate(*guildobj)
	if !status {
		log.Printf("Error : Guild OP Failled")
		// add log
	}
	if m.Author.ID == s.State.User.ID {
		return
	}
	prefix := config.Prefix
	var perm bool = false
	switch m.Content {
	case prefix + "ping" : {
		var cmd Cmd = &ping{}
		cmd.props(s,m)
		if !cmd.checkperm(s, m) {
			perm = true 
			break
		}
		cmd.run(s, m)
	}
	}
	if perm {
		s.ChannelMessageSend(m.ChannelID, lang.EN_PermMissing)
	}
	
} 