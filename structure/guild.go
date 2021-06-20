package structure

import (
	"github.com/bwmarrin/discordgo"
)

type Guild struct {
	ID string
	DiscGuild discordgo.Guild
	Tags []string
	desc string
}
func NewGuild() *Guild {
	return new(Guild)
}
func (guild *Guild) addtag(tag string) {
	guild.Tags = append(guild.Tags, tag)
}
func (guild *Guild) updatedesc(desc string) {
	guild.desc = desc
}