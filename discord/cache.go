package discord

import (
	"sync"

	"github.com/allegro/bigcache/v3"
	"github.com/bigopenworld/discord-bot/config"
	"github.com/bigopenworld/discord-bot/data"
	"github.com/bigopenworld/discord-bot/structure"
)

type BotCache struct {
	locked bool	
	sync.RWMutex
	GuildCache *bigcache.BigCache
	MemberCache *bigcache.BigCache
	ConfigCache *bigcache.BigCache
}
func NewBotCache() BotCache {
	return BotCache{}
}
func (botcache *BotCache) LockAllCache() bool {
	botcache.Lock()
	botcache.locked = true
	return true
}

func (botcache *BotCache) UnlockAllCache() bool {
	botcache.Unlock()
	botcache.locked = false
	return true
}


func (botcache *BotCache) init() bool {
	println("Init cache ... Locking cache struct")
	botcache.LockAllCache()

	// Guild cache init
	println("Init cache ... 1 of 3 : GuildCache")
	GuildCacheConfig := bigcache.Config {
		Shards: 1024,
		LifeWindow: config.GuildCacheExp,
		CleanWindow: config.ConfigCacheClean,
		MaxEntriesInWindow: 1000 * 10 * 60,
		MaxEntrySize: 500,
		Verbose: false,
		HardMaxCacheSize: config.MaxGuildMem,
		OnRemove: nil,
		OnRemoveWithReason: nil,
	}
	GuildCache, initErrGuildCache := bigcache.NewBigCache(GuildCacheConfig)
	if initErrGuildCache != nil {
		return false
	}
	botcache.GuildCache = GuildCache

	// Member cache init
	println("Init cache ... 2 of 3 : MemberCache")
	MemberCacheConfig := bigcache.Config {
		Shards: 1024,
		LifeWindow: config.GuildCacheExp,
		CleanWindow: config.ConfigCacheClean,
		MaxEntriesInWindow: 1000 * 10 * 60,
		MaxEntrySize: 500,
		Verbose: false,
		HardMaxCacheSize: config.MaxMemberMem,
		OnRemove: nil,
		OnRemoveWithReason: nil,
	}
	MemberCache, initErrMemberCache := bigcache.NewBigCache(MemberCacheConfig)
	if initErrMemberCache != nil {
		return false
	}
	botcache.MemberCache = MemberCache

	// Config cache init
	println("Init cache ... 3 of 3 : ConfigCache")
	ConfigCacheConfig := bigcache.Config {
		Shards: 1024,
		LifeWindow: config.GuildCacheExp,
		CleanWindow: config.ConfigCacheClean,
		MaxEntriesInWindow: 1000 * 10 * 60,
		MaxEntrySize: 500,
		Verbose: false,
		HardMaxCacheSize: config.MaxConfigMem,
		OnRemove: nil,
		OnRemoveWithReason: nil,
	}
	ConfigCache, initErrConfigCache := bigcache.NewBigCache(ConfigCacheConfig)
	if initErrConfigCache != nil {
		return false
	}
	botcache.ConfigCache = ConfigCache
	println("Init cache ... Unlocking cache struct")
	botcache.UnlockAllCache()
	println("All Cache init done !")
	return true
}




////////////////////////// Cache methods /////////////////////////

//// Write 

func (botcache *BotCache) SetGuild(guildobj structure.Guild) {
	bytes := data.EncodeToBytes(guildobj)
	if config.GuildMemCompression {
		bytes = data.Compress(bytes)
	}
	botcache.GuildCache.Set(guildobj.ID, bytes)
}
func (botcache *BotCache) SetMember(memberobj structure.Member) {
	bytes := data.EncodeToBytes(memberobj)
	if config.MemberMemCompression {
		bytes = data.Compress(bytes)
	}
	botcache.MemberCache.Set(memberobj.ID, bytes)
}
func (botcache *BotCache) SetConfig(configobj structure.Config) {
	bytes := data.EncodeToBytes(configobj)
	if config.ConfigMemCompression {
		bytes = data.Compress(bytes)
	}
	botcache.ConfigCache.Set(configobj.ID, bytes)
}

/// Read

func (BotCache *BotCache) GetGuild(id string) (bool, structure.Guild){
	bytes, err := BotCache.GuildCache.Get(id)
	if err != nil {
		return false, structure.Guild{}
	}
	if config.GuildMemCompression {
		bytes = data.Decompress(bytes)
	}
	return true, data.DecodeToGuild(bytes)
}
func (BotCache *BotCache) GetMember(id string) (bool, structure.Member){
	bytes, err := BotCache.MemberCache.Get(id)
	if err != nil {
		return false, structure.Member{}
	}
	if config.MemberMemCompression {
		bytes = data.Decompress(bytes)
	}
	return true, data.DecodeToMember(bytes)
}
func (BotCache *BotCache) GetConfig(id string) (bool, structure.Config){
	bytes, err := BotCache.ConfigCache.Get(id)
	if err != nil {
		return false, structure.Config{}
	}
	if config.ConfigMemCompression {
		bytes = data.Decompress(bytes)
	}
	return true, data.DecodeToConfig(bytes)
}