package structure

import "github.com/bwmarrin/discordgo"

type Member struct {
	ID string
	DiscMember discordgo.Member
}
type Guild struct {
	ID string
	DiscGuild discordgo.Guild
}
type Config struct {
	ID string
}