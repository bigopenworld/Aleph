package structure

import (
	"github.com/bwmarrin/discordgo"
)

type Guild struct {
	ID string
	DiscGuild discordgo.Guild
	fetch bool
	Tags []string
	desc string
}
func NewGuild(id string) *Guild {
	g := new(Guild)
	g.ID = id
	g.fetch = false
	return g
}
func (guild *Guild) Addtag(tag string) {
	guild.Tags = append(guild.Tags, tag)
}
func (guild *Guild) Updatedesc(desc string) {
	guild.desc = desc
}
func (guild *Guild) IsFetch() bool {
	return guild.fetch
}
func (guild *Guild) Setdisc(disc *discordgo.Guild) {
	if !guild.fetch {
		guild.DiscGuild = *disc
	}
}