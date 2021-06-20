package structure

import (
	"github.com/bigopenworld/discord-bot/config"
	"github.com/bigopenworld/discord-bot/database"
	"github.com/bwmarrin/discordgo"
	"gopkg.in/rethinkdb/rethinkdb-go.v6"
)


type Guild struct {
	ID string
	DiscGuild discordgo.Guild
	tags [5]string
}

func (guild *Guild) create() bool {
	err := rethinkdb.DB(config.DBdatabase).Table(config.DBusertable).Insert(map[string]interface{}{
		"id": guild.ID,
		"tags": guild.tags,
	}).Exec(database.DBsession)
	return err != nil
}