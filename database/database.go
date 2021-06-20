package database

import (
	"github.com/bigopenworld/discord-bot/config"
	"gopkg.in/rethinkdb/rethinkdb-go.v6"
)

var DBsession *rethinkdb.Session


func Connect() bool {
	var opt rethinkdb.ConnectOpts
	switch config.DBmultihosts {
	case true : 
		opt = rethinkdb.ConnectOpts{
			Addresses: config.DBhosts,
			InitialCap: config.DBInitialCap,
			MaxOpen: config.DBMaxOpen,
			NumRetries: config.DBMaxRetry,
			DiscoverHosts: config.DBDiscoverHosts,
			Username: config.DBuser,
			Password: config.DBpassword,
		}
	case false : 
		opt = rethinkdb.ConnectOpts{
			Address: config.DBhost,
			InitialCap: config.DBInitialCap,
			MaxOpen: config.DBMaxOpen,
			NumRetries: config.DBMaxRetry,
			DiscoverHosts: config.DBDiscoverHosts,
			Username: config.DBuser,
			Password: config.DBpassword,
		}
	}
	session, err := rethinkdb.Connect(opt)
	DBsession = session
	return err == nil
}
func Test() bool {
	_, err := rethinkdb.Expr("Hello World").Run(DBsession)
	return err == nil
}