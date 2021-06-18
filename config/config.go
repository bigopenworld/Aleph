package config

import (
	"time"
)

// DataBase config

var DBhost = "" // leave blank for localhost
var DBuser = "" // leave blank for none
var DBpassword = "" // leave blank for none
var DBdatabase = "" // leave blank for default (bigopenworld-discordbot)

// Discord config

var Token = "" // enter the bot token

// Bot settings 

var Cache = true // true = cache enabled / false = cache disabled

// Bot cache settings (disabled if cache disabled) dont modify unless you know what you are doing 

/*
Cron settings 

Seconds Minutes Hours Day-of-Month Month Day-of-week 

ex : "0 15 * * * *" run every 15 min 
*/

var GuildCacheForceClean = "0 20 * * * *" // dont modify unless you know what you are doing 
var UserCacheForceClean = "0 20 * * * *" // dont modify unless you know what you are doing 
var ConfigCacheForceClean = "0 10 * * * *" // dont modify unless you know what you are doing 

var GuildCacheExp = 10*time.Minute
var MemberCacheExp = 10*time.Minute
var ConfigCacheExp = 5*time.Minute

var GuildCacheClean = 10*time.Minute
var MemberCacheClean = 10*time.Minute
var ConfigCacheClean = 5*time.Minute
