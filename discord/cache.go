package discord

import (
	"fmt"
	"sync"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/bigopenworld/discord-bot/cmd"
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
	Lcooldown *bigcache.BigCache
	Hcooldown *bigcache.BigCache
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
	println(cmd.NewFlag(cmd.OK),"Init cache ... Locking cache struct")
	botcache.LockAllCache()

	// Guild cache init
	println(cmd.NewFlag(cmd.OK),"Init cache ... 1 of 2 : GuildCache")
	GuildCacheConfig := bigcache.Config {
		Shards: 1024,
		LifeWindow: config.GuildCacheExp,
		CleanWindow: config.GuildCacheClean,
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
	println(cmd.NewFlag(cmd.OK),"Init cache ... 2 of 2 : MemberCache")
	MemberCacheConfig := bigcache.Config {
		Shards: 1024,
		LifeWindow: config.MemberCacheExp,
		CleanWindow: config.MemberCacheClean,
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

	
	println(cmd.NewFlag(cmd.OK),"Init cache ... Unlocking cache struct")
	botcache.UnlockAllCache()
	println(cmd.NewFlag(cmd.SUCCESS),"All Cache init done !")
	return true
}
func (botcache *BotCache) initcooldown() bool {
	println(cmd.NewFlag(cmd.OK),"Init cooldown cache ... Locking cache struct")
	botcache.LockAllCache()

	// Lcooldown cache init
	println(cmd.NewFlag(cmd.OK),"Init cache ... 1 of 2 : Lcooldown")
	if config.LCooldownCache {
		Lcooldownconfig := bigcache.Config {
			Shards: 1024,
			LifeWindow: config.LCacheExp,
			CleanWindow: config.LCacheClean,
			MaxEntriesInWindow: 1000 * 10 * 60,
			MaxEntrySize: 500,
			Verbose: false,
			HardMaxCacheSize: config.LCooldownMem,
			OnRemove: nil,
			OnRemoveWithReason: nil,
		}
		LcooldownCache, initErrGuildCache := bigcache.NewBigCache(Lcooldownconfig)
		if initErrGuildCache != nil {
			return false
		}
		botcache.Lcooldown = LcooldownCache
	} else {
		fmt.Println(cmd.NewFlag(cmd.WARNING),"Lcooldown Cache disabled ... skiping")
	}


	// Hcooldown cache init
	println(cmd.NewFlag(cmd.OK),"Init cache ... 2 of 2 : Hcooldown")
	if config.HCooldownCache {
		Hcooldownconfig := bigcache.Config {
			Shards: 1024,
			LifeWindow: config.MemberCacheExp,
			CleanWindow: config.MemberCacheClean,
			MaxEntriesInWindow: 1000 * 10 * 60,
			MaxEntrySize: 500,
			Verbose: false,
			HardMaxCacheSize: config.HCooldownMem,
			OnRemove: nil,
			OnRemoveWithReason: nil,
		}
		HcooldownCache, initErrMemberCache := bigcache.NewBigCache(Hcooldownconfig)
		if initErrMemberCache != nil {
			return false
		}
		botcache.Hcooldown = HcooldownCache
	} else {
		fmt.Println(cmd.NewFlag(cmd.WARNING),"Hcooldown Cache disabled ... skiping")
	}
	
	println(cmd.NewFlag(cmd.OK),"Init cache ... Unlocking cache struct")
	botcache.UnlockAllCache()
	println(cmd.NewFlag(cmd.SUCCESS),"All Cache init done !")
	return true
}





////////////////////////// Cache methods /////////////////////////

//// Write 

func (botcache *BotCache) SetGuild(guildobj structure.Guild) {
	if config.Cache {
		bytes := data.EncodeToBytes(guildobj)
		if config.GuildMemCompression {
			bytes = data.Compress(bytes)
		}
		botcache.GuildCache.Set(guildobj.ID, bytes)
	}
}
func (botcache *BotCache) SetMember(memberobj structure.Member) {
	if config.Cache {
		bytes := data.EncodeToBytes(memberobj)
		if config.MemberMemCompression {
			bytes = data.Compress(bytes)
		}
		botcache.MemberCache.Set(memberobj.ID, bytes)
	}

}
func (botcache *BotCache) SetHcool(memberobj structure.Member,cooldowncmd string, t time.Time) {
	if config.HCooldownCache {
		bytes := data.EncodeToBytes(t)
		if config.HCooldownCompression {
			bytes = data.Compress(bytes)
		}
		botcache.MemberCache.Set(memberobj.ID+cooldowncmd, bytes)
	}
}
func (botcache *BotCache) SetLcool(memberobj structure.Member,cooldowncmd string, t time.Time) {
	if config.LCooldownCache {
		bytes := data.EncodeToBytes(t)
		if config.LCooldownCompression {
			bytes = data.Compress(bytes)
		}
		err := botcache.Lcooldown.Set(memberobj.ID+cooldowncmd, bytes)
		if err != nil {
			println("error when writing cache")
		}
	}
}


/// Read

func (BotCache *BotCache) GetGuild(id string) (bool, structure.Guild){
	if config.Cache {
		bytes, err := BotCache.GuildCache.Get(id)
		if err != nil {
			return false, structure.Guild{}
		}
		if config.GuildMemCompression {
			bytes = data.Decompress(bytes)
		}
		return true, data.DecodeToGuild(bytes)
	} else {
		return false, structure.Guild{}
	}
}
func (BotCache *BotCache) GetMember(id string) (bool, structure.Member){
	if config.Cache {
		bytes, err := BotCache.MemberCache.Get(id)
		if err != nil {
			return false, structure.Member{}
		}
		if config.MemberMemCompression {
			bytes = data.Decompress(bytes)
		}
		return true, data.DecodeToMember(bytes)
	} else {
		return false, structure.Member{}
	}
}
func (BotCache *BotCache) GetLcooldown(id string, cooldowncmd string) (bool, time.Time){
	if config.LCooldownCache {
		bytes, err := BotCache.Lcooldown.Get(id+cooldowncmd)
		if err != nil {
			println(err.Error())
			return false, time.Time{}
		}
		if config.MemberMemCompression {
			bytes = data.Decompress(bytes)
		}
		return true, data.DecodeToCooldown(bytes)
	} else {
		return false, time.Time{}
	}
}
func (BotCache *BotCache) GetHcooldown(id string, cooldowncmd string) (bool, time.Time){
	if config.HCooldownCache {
		bytes, err := BotCache.Hcooldown.Get(id+cooldowncmd)
		if err != nil {
			return false, time.Time{}
		}
		if config.MemberMemCompression {
			bytes = data.Decompress(bytes)
		}
		return true, data.DecodeToCooldown(bytes)
	} else {
		return false, time.Time{}
	}
}
