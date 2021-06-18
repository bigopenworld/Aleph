package discord

import (
	"sync"

	"github.com/bigopenworld/discord-bot/config"
	"github.com/patrickmn/go-cache"
)

type BotCache struct {
	locked bool	
	sync.RWMutex
	GuildCache *cache.Cache
	MemberCache *cache.Cache
	ConfigCache *cache.Cache
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
	botcache.LockAllCache()
	botcache.GuildCache = cache.New(config.GuildCacheExp, config.GuildCacheClean)
	botcache.MemberCache = cache.New(config.MemberCacheExp, config.MemberCacheClean)
	botcache.ConfigCache = cache.New(config.ConfigCacheExp, config.ConfigCacheClean)
	botcache.UnlockAllCache()
	return true
}