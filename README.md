# Discord Bot (golang)

## Requirements : 

- Rethinkdb (2.4.1 or newer) 
- golang (1.16 or newer)

> Note : You can disable the database option but this will disable many bot function 

## How to install the bot :

download the release from github, extract it

edit `./config/config.go` fill the token and other missing settings 

> Note : Please set MaxGuildMem & MaxMemberMem & MaxConfigMem according to your system memory

run `go build`

run `./discord-bot`

## Status debug :

exit code 0 => success (when shutdown command is done)

exit code 1 => Error occured

exit code 3 => Bot init error

exit code 4 => Bot connect error

exit code 5 => Bot cache init error

exit code 6 => Bot cache filling error

exit code 7 => DataBase connection error

## CMD that exist :

> Ping => Return the bot latency

---

[Discord] (https://discord.gg/)
