package data

import (
	"bytes"
	"encoding/gob"
	"log"

	"github.com/bigopenworld/discord-bot/structure"
)

func DecodeToGuild(s []byte) structure.Guild {

	t := structure.Guild {}
	dec := gob.NewDecoder(bytes.NewReader(s))
	err := dec.Decode(&t)
	if err != nil {
		log.Fatal(err)
	}
	return t
}

func DecodeToMember(s []byte) structure.Member {

	t := structure.Member{}
	dec := gob.NewDecoder(bytes.NewReader(s))
	err := dec.Decode(&t)
	if err != nil {
		log.Fatal(err)
	}
	return t
}

func DecodeToConfig(s []byte) structure.Config {

	t := structure.Config{}
	dec := gob.NewDecoder(bytes.NewReader(s))
	err := dec.Decode(&t)
	if err != nil {
		log.Fatal(err)
	}
	return t
}