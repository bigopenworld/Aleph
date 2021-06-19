package discord

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/bigopenworld/discord-bot/config"
	"github.com/bigopenworld/discord-bot/database"
	"github.com/bigopenworld/discord-bot/structure"
	"github.com/bwmarrin/discordgo"
)

var BotVar *BotStruct

type BotStruct struct {
	sync.RWMutex
	restarted int
	session   *discordgo.Session
	cache     BotCache
}

// constructor
func CreateAndStart() bool {
	BotVar = new(BotStruct)
	BotVar.start()
	return true
}

// start func
func (bot *BotStruct) start() bool {
	fmt.Println("Bot Starting ... Locking Bot struct")
	bot.Lock()
	fmt.Println("Bot Starting ... 1 of 3 : Bot init and connection")
	err := bot.connect()
	if err != 0 {
		fmt.Println("Bot Starting failed ... Unlocking Bot struct")
		bot.Unlock()
		bot.tryrestartorkill(err, false)
		return true
	}
	fmt.Println("Bot Starting ... 2 of 3 : Bot cache starting")
	if config.Cache {
		bot.cache = NewBotCache()
		cacheresult := bot.cache.init()
		if !cacheresult {
			fmt.Println("Bot Starting failed ... Unlocking Bot struct")
			bot.Unlock()	
			bot.tryrestartorkill(5, false)
			return true
		}
	} else {
		fmt.Println("Cache disabled ... skiping")
	}

	fmt.Println("Bot Starting ... 3 of 3 : DataBase connection init")
	dbstatus := database.Connect()
	if !dbstatus {
		fmt.Println("Bot Starting failed ... Unlocking Bot struct")
		bot.Unlock()
		bot.tryrestartorkill(7, false)
		return true
	}
	fmt.Println("Bot Starting ... Unlocking Bot struct")
	bot.Unlock()
	fmt.Println("All done !")
	bot.fillcache()
	return true
}

func (bot *BotStruct) Shutdown() {
	bot.tryrestartorkill(0, true)
}

// restarter or killer
func (bot *BotStruct) tryrestartorkill(code int, kill bool) {
	if kill {
		bot.session.Close()
		fmt.Println("Bot killed nicely & peacefully")
		os.Exit(code)
	}
	bot.restarted++
	fmt.Println("Bot Error Occured ... Restart ", bot.restarted, " of ", config.MaxRestart)
	if bot.restarted < config.MaxRestart {
		bot.session.Close()
		time.Sleep(config.RestartWait)
		bot.start()
		return
	} else {
		os.Exit(code)
	}
}

// connect
func (bot *BotStruct) connect() int {
	discord, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		return 3
	}
	bot.session = discord
	bot.session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
	err = bot.session.Open()
	if err != nil {
		return 4
	}
	return 0
}


/* 
Cache Part 

*/

// Get the cache struct
func (bot *BotStruct) GetCache() *BotCache {
	return &bot.cache
}

// fill the bot cache to maximum => Begin with Guilds / Config / Users
func (bot *BotStruct) fillcache() bool {
	
	bot.Lock()
	fmt.Println("Init 1 of 1 ... Filling cache")

	listoftest := splitdata(bot.session.State.Guilds)
	var wg sync.WaitGroup
	for pid, list := range listoftest {
		wg.Add(1)
		go fillcacheprocess(pid, list, bot, &wg)
	}
	wg.Wait()
	fmt.Println("Init Done ... All done")
	bot.Unlock()

	//client.session.State.RUnlock()
	return true
}

// TODO : Fill the cache + cache memory management

//
func fillcacheprocess(pid int, list []*discordgo.Guild, client *BotStruct, wg *sync.WaitGroup) {
	defer wg.Done()
	var total = len(list)
	var done int
	for _, guild := range list {
		done++
		//fmt.Println(guild.ID)
		//guild_fetch, err := client.session.Guild(guild.ID)
		BotVar.cache.SetGuild(structure.Guild{
			ID : guild.ID,
			DiscGuild: *guild,
		})
		/*if err != nil {
			log.Panic("Error when filling cache")
		}*/
		//fmt.Println(guild.ID)
		//client.cache.GuildCache.Set(guild.ID, guild.ApproximateMemberCount, cache.DefaultExpiration )
		//client.cache.guildmembersonline.Set(guild.ID, guild.ApproximateMemberCount, cache.DefaultExpiration )
		//log.Println("Process",pid ," : guild ", done, " / ",total)

	}
	fmt.Println("Process", pid, " : guild ", done, " / ", total)
}

// DataSpliter
func splitdata(list []*discordgo.Guild) [config.Chunk + 1][]*discordgo.Guild {
	var listlen = len(list)
	var part int = listlen / config.Chunk
	var res [config.Chunk + 1][]*discordgo.Guild
	for i, elem := range list {
		res[i/part] = append(res[i/part], elem)
	}
	return res
}
