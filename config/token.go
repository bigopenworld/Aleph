package config

import "os"

func get_token(tokenconfig string) string {
	var return_token string = tokenconfig
	token, ok  := os.LookupEnv("TOKEN")
	if ok {
		return_token = token
	}
	return return_token
}