package config

import (
	"time"
)

// DataBase config

const DBhost = "" // leave blank for localhost
const DBuser = "" // leave blank for none
const DBpassword = "" // leave blank for none
const DBdatabase = "bigopenworld-discordbot" // leave blank for default (bigopenworld-discordbot)

// Discord config

const Token = "" // enter the bot token

// Bot settings 

const Cache = true // true = cache enabled / false = cache disabled
const MaxRestart = 5 // how many restart are authorized before killing the process
const RestartReset = true // If the bot have sucessufly started do we resset restart cont to 0
const RestartWait = 10*time.Second // How many time we wait before a new restart attempt

// Bot cache settings (disabled if cache disabled)  

/*
Cron settings 

Seconds Minutes Hours Day-of-Month Month Day-of-week 

ex : "0 15 * * * *" run every 15 min 
*/

const GuildCacheForceClean = "0 20 * * * *" // dont modify unless you know what you are doing 
const UserCacheForceClean = "0 20 * * * *" // dont modify unless you know what you are doing 
const ConfigCacheForceClean = "0 10 * * * *" // dont modify unless you know what you are doing 

const GuildCacheExp = 10*time.Minute // dont modify unless you know what you are doing 
const MemberCacheExp = 10*time.Minute // dont modify unless you know what you are doing 
const ConfigCacheExp = 5*time.Minute // dont modify unless you know what you are doing 

const GuildCacheClean = 10*time.Minute // dont modify unless you know what you are doing 
const MemberCacheClean = 10*time.Minute // dont modify unless you know what you are doing 
const ConfigCacheClean = 5*time.Minute // dont modify unless you know what you are doing 

/* 
By default this set is 1, increment if your bot take a long time to load data to the cache

this setting define how many process are spawned to fill the cache
*/
const Chunk = 1 