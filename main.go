package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/hokaccha/go-prettyjson"
)

func main() {
	flag.Parse()
	args := flag.Args()

	numArgs := len(args)
	if numArgs != 1 {
		log.Fatal("invalid number of args")
	}

	p := args[0]
	bytes, err := ioutil.ReadFile(p)
	if err != nil {
		log.Fatal(err)
	}

	var deserialized map[string]interface{}
	err = json.Unmarshal(bytes, &deserialized)
	if err != nil {
		log.Fatal(err)
	}

	pretty, err := prettyjson.Marshal(deserialized)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(pretty))
}
