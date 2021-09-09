package config

import (
	"time"
)

// DataBase config

const (
	DBenabled       = true                      // change to false if you dont want to store data => changing to false will disable some bot fuction
	DBhost          = ""                        // leave blank for localhost
	DBuser          = ""                        // leave blank for none
	DBpassword      = ""                        // leave blank for none
	DBdatabase      = "bigopenworld-discordbot" // leave for default (bigopenworld-discordbot)
	DBusertable     = "users"                   // where to store user
	DBguildtable    = "guild"                   // where to store guild
	DBcooldowntable = "cooldown"
)

// DataBase advanced config

var (
	DBmultihosts    = false                      // to use multiple hosts instead of one (disable "DBhost" the setting for one host only)
	DBhosts         = []string{"host1", "host2"} // replace host1 and host2 with the correct value
	DBInitialCap    = 10                         // how many connection should be created
	DBMaxOpen       = 20                         // how many connection is allowed
	DBMaxRetry      = 3                          // how many retry sould be attempt when a query fail
	DBDiscoverHosts = false                      // should the database automatically connect new servers when they are added to the cluster
)

// Discord config

var Token = get_token("") // enter the bot token like get_token("your-token")

// Bot settings

const (
	Prefix       = ">"              // the default prefix of the bot
	Cache        = true             // true = cache enabled / false = cache disabled
	MaxRestart   = 5                // how many restart are authorized before killing the process
	RestartReset = true             // If the bot have sucessufly started do we resset restart cont to 0
	RestartWait  = 10 * time.Second // How many time we wait before a new restart attempt
)

// Bot cache settings (disabled if cache disabled)

const (
	GuildCacheExp  = 10 * time.Minute // dont modify unless you know what you are doing
	MemberCacheExp = 10 * time.Minute // dont modify unless you know what you are doing

	GuildCacheClean  = 10 * time.Minute // dont modify unless you know what you are doing
	MemberCacheClean = 10 * time.Minute // dont modify unless you know what you are doing
)

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

const (
	MaxGuildMem  = 1024 // Value in MB => 1024 = 1 GB of Ram
	MaxMemberMem = 1024 // Value in MB => 1024 = 1 GB of Ram
	LCooldownMem = 1024 // Value in MB => 1024 = 1 GB of Ram  => This set the low Cooldown memory cache (for no-spam cooldown)
	HCooldownMem = 1024 // Value in MB => 1024 = 1 GB of Ram  => This set the high Cooldown memory cache (for cooldown like RP cmd cooldown)

	GuildMemCompression  = true // chose to enable or disable cache compression
	MemberMemCompression = true // chose to enable or disable cache compression
	LCooldownCompression = true // chose to enable or disable cache compression
	HCooldownCompression = true // chose to enable or disable cache compression
)

/*
Cooldown settings

*/

const (
	LCooldownDB = true
	HCooldownDB = true

	LCooldownCache = true
	HCooldownCache = true

	LCacheExp = 10 * time.Minute // dont modify unless you know what you are doing
	HCacheExp = 10 * time.Minute // dont modify unless you know what you are doing

	LCacheClean = 10 * time.Minute // dont modify unless you know what you are doing
	HCacheClean = 10 * time.Minute // dont modify unless you know what you are doing
)
