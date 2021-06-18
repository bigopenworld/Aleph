# Discord Bot (golang)

## Requirements : 

- Rethinkdb (2.4.1 or newer)
- golang (1.16 or newer)

## How to install the bot :

download the release from github, extract it

edit `./config/config.go` fill the token and other missing settings 

run `go build`

run `./discord-bot`

## Status debug :

exit code 0 => success (when shutdown command is done)

exit code 1 => Error occured

exit code 3 => Bot init error

exit code 4 => Bot connect error

exit code 5 => Bot cache init error

exit code 6 => Bot cache filling error