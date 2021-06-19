package data

import (
	"io/ioutil"
	"log"
	"os"
)

func WriteToFile(s []byte, file string) {

	f, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}

	f.Write(s)
}

func ReadFromFile(path string) []byte {

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
