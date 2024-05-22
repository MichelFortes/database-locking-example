package config

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	User string
	Pass string
	DB   string
}

func (c *Config) validate() {

	if len(c.User) == 0 {
		panic(errors.New("user is not set"))
	}
	if len(c.Pass) == 0 {
		panic(errors.New("pass is not set"))
	}
	if len(c.DB) == 0 {
		panic(errors.New("db is not set"))
	}
}

func FromArgs() (Config, error) {
	var confPath string
	flag.StringVar(&confPath, "c", "", "-c [config file]")
	flag.Parse()
	return parse(confPath)
}

func parse(file string) (Config, error) {

	conf := Config{}

	if len(file) == 0 {
		return conf, errors.New("file is not set")
	}

	b, e := os.ReadFile(file)
	if e != nil {
		return conf, fmt.Errorf("could not read config file %s", file)
	}

	for i, line := range strings.Split(string(b), "\n") {
		if !strings.Contains(line, "=") {
			return conf, fmt.Errorf("linvalid format at line %d ", i)
		}
		parts := strings.Split(line, "=")
		switch parts[0] {
		case "POSTGRES_USER":
			conf.User = parts[1]
		case "POSTGRES_PASSWORD":
			conf.Pass = parts[1]
		case "POSTGRES_DB":
			conf.DB = parts[1]
		}
	}

	conf.validate()

	return conf, nil
}
