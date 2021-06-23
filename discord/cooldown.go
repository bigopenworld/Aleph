package discord

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/bigopenworld/discord-bot/cmd"
	"github.com/bigopenworld/discord-bot/config"
	"github.com/bigopenworld/discord-bot/database"
	"github.com/bigopenworld/discord-bot/structure"
	"gopkg.in/rethinkdb/rethinkdb-go.v6"
)
type Test struct {
	ping string
}
type Cooldown struct {
	id string
	data map[string]string
}

// Exp is in sec => true cooldown active / false cooldown expired
func LCooldownGet(user structure.Member, command string, exp time.Duration) (bool, error){
	valid, res := BotVar.GetCache().GetLcooldown(user.ID, command)
	if valid {
		t := res.Add(exp) 
		//println(t, "  ", time.Now().Unix())
		if ( t.UTC().Unix() > time.Now().UTC().Unix() ) {
			return true, nil
		}
		return false, nil
	}
	if config.DBenabled && config.LCooldownDB {
		resdb, err := rethinkdb.DB(config.DBdatabase).Table(config.DBcooldowntable).Get(user.ID).Run(database.DBsession)
		if err != nil {
			return false, errors.New("DataBase Error")
		}
		var doc map[string]string
		resdb.Next(&doc)
		value, exist := doc["data."+command]
		if !exist {
			return false, nil
		}
		t, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return false, errors.New("Error in string conv")
		}
		timeparsed := time.Unix(t, 0)
		if ( timeparsed.Add(exp).Unix()  > time.Now().Unix() ) {
			BotVar.GetCache().SetLcool(user, command, timeparsed)
			return true, nil
		}
		return false, nil

	}
	return false, nil
	
} 

func HCooldownGet(user structure.Member, cmd string) bool{
	return true 
} 

func LCooldownset(user structure.Member, command string){
	t := time.Now()
	BotVar.GetCache().SetLcool(user, command, t)
	if config.DBenabled && config.LCooldownDB {
		
		err := rethinkdb.DB(config.DBdatabase).Table(config.DBcooldowntable).Insert(map[string]string{
			"id": user.ID,
			"data."+command: strconv.FormatInt(t.UTC().Unix(), 10),
		}, rethinkdb.InsertOpts{
			Conflict: "update",
		}).Exec(database.DBsession)
		if err != nil {
			fmt.Println(cmd.NewFlag(1),"Error occured when writing cooldown to DB")
		}

	}
}
/*
"data": map[string]string {
				"ping": strconv.Itoa(int(t.UTC().Unix())),
			},
*/