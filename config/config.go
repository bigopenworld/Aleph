package config

import (
	"time"
)

// DataBase config

const DBenabled = true // change to false if you dont want to store data => changing to false will disable some bot fuction
const DBhost = "" // leave blank for localhost
const DBuser = "" // leave blank for none
const DBpassword = "" // leave blank for none
const DBdatabase = "bigopenworld-discordbot" // leave for default (bigopenworld-discordbot)
const DBusertable = "users"// where to store user
const DBguildtable = "guild" // where to store guild

// DataBase advanced config 

const DBmultihosts = false // to use multiple hosts instead of one (disable "DBhost" the setting for one host only)
var DBhosts = []string{"host1", "host2"} // replace host1 and host2 with the correct value
const DBInitialCap = 10 // how many connection should be created
const DBMaxOpen = 20 // how many connection is allowed
const DBMaxRetry = 3 // how many retry sould be attempt when a query fail
const DBDiscoverHosts = false // should the database automatically connect new servers when they are added to the cluster

// Discord config

const Token = "" // enter the bot token

// Bot settings 

const Prefix = ">" // the default prefix of the bot
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

/* this settings does nothing now maybe will be use for a next version
const GuildCacheForceClean = "0 20 * * * *" // dont modify unless you know what you are doing 
const UserCacheForceClean = "0 20 * * * *" // dont modify unless you know what you are doing 
*/

const GuildCacheExp = 10*time.Minute // dont modify unless you know what you are doing 
const MemberCacheExp = 10*time.Minute // dont modify unless you know what you are doing 

const GuildCacheClean = 10*time.Minute // dont modify unless you know what you are doing 
const MemberCacheClean = 10*time.Minute // dont modify unless you know what you are doing 

/* 
By default this set is 1, increment if your bot take a long time to load data to the cache

this setting define how many process + 1 are spawned to fill the cache 

*/
const Chunk = 1

/* 
Memory management settings 

Warning : setting this value too high can crash your system
Warning : setting this to 0 disable memory management => Use with caution
*/

const MaxGuildMem = 1024 // Value in MB => 1024 = 1 GB of Ram 
const MaxMemberMem = 1024 // Value in MB => 1024 = 1 GB of Ram 

const GuildMemCompression = true // chose to enable or disable cache compression
const MemberMemCompression = true // chose to enable or disable cache compression