package data

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"

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

func DecodeToCooldown(s []byte) time.Time {

	t := time.Time{}
	dec := gob.NewDecoder(bytes.NewReader(s))
	err := dec.Decode(&t)
	if err != nil {
		log.Fatal(err)
	}
	return t
}
