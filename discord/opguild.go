package discord

import (
	"errors"

	"github.com/bigopenworld/discord-bot/config"
	"github.com/bigopenworld/discord-bot/database"
	"github.com/bigopenworld/discord-bot/structure"
	"gopkg.in/rethinkdb/rethinkdb-go.v6"
)



func GuildCreate(guild structure.Guild) bool {
	res, errex := GuildExist(guild)
	if errex != nil {
		return false
	}
	if res {
		return true
	}
	err := rethinkdb.DB(config.DBdatabase).Table(config.DBguildtable).Insert(map[string]interface{}{
		"id": guild.ID,
		"tags": guild.Tags,
	}).Exec(database.DBsession)
	return err == nil
}

func GuildExist(guild structure.Guild) (bool, error) {
	res, err := rethinkdb.DB(config.DBdatabase).Table(config.DBguildtable).Count().Run(database.DBsession)


	if err != nil {
		return true, errors.New("Unknow DB error")
	}

	var doc int
	res.Next(&doc)
	res.Close()
	return (doc > 0), nil
}
func GuildGet(guild structure.Guild) (structure.Guild, error) {
	finded, guildata := BotVar.GetCache().GetGuild(guild.ID)
	if finded {
		return guildata, nil
	} else {
		res, err := rethinkdb.DB(config.DBdatabase).Table(config.DBguildtable).Get(guild.ID).Run(database.DBsession)
		if err != nil {
			return structure.Guild{}, errors.New("DataBase Error")
		}
		var doc structure.Guild
		res.Next(&doc)
		return doc, nil
	}
	
}