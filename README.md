# Discord Bot (golang)

## Requirements : 

- Rethinkdb (2.4.1 or newer)
- golang (1.16 or newer)

## How to install the bot :

download the release from github, extract it

`go build`

`./discord-bot`

## Status debug :

code 0 => success (when shutdown command is done)

code 1 => Error occured

code 3 => Bot init error

code 4 => Bot connect error

code 5 => Bot cache init error

code 6 => Bot cache filling error