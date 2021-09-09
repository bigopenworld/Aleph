package structure

import "github.com/bwmarrin/discordgo"


type Member struct {
	ID string
	DiscMember discordgo.User
}