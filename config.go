package main

import (
	"github.com/BurntSushi/toml"
	"log"
)

type JobDesc struct {
	Dependencies []string
        Command string
}

type Configuration struct {
	Jobdescs map[string]JobDesc
}

func LoadConfig(filename string) Configuration {
	var conf Configuration
	if _, err := toml.DecodeFile(filename, &conf); err != nil {
                log.Fatal(err)
	}
        return conf
}
