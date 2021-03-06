package main

import (
	"flag"
	"github.com/evilsocket/islazy/log"
)

var (
	debug     = false
	ver       = false
	address   = "127.0.0.1:8080"
	confFile  = "/etc/joe/joe.conf"
	usersPath = "/etc/joe/users"
	dataPath  = "/etc/joe/queries"

	newUser  = ""
	tokenTTL = 24

	docOutput = ""
	docFormat = "markdown"
)

func init() {
	flag.BoolVar(&debug, "debug", debug, "Enable debug logs.")
	flag.StringVar(&log.Output, "log", log.Output, "Log file path or empty for standard output.")
	flag.StringVar(&address, "address", address, "API address.")
	flag.StringVar(&confFile, "conf", confFile, "Configuration file.")
	flag.StringVar(&usersPath, "users", usersPath, "Path containing user credentials in YML.")
	flag.StringVar(&dataPath, "data", dataPath, "Data path.")

	flag.BoolVar(&ver, "version", ver, "Print version and exit.")

	flag.StringVar(&newUser, "new-user", newUser, "Create a new user with the provided username.")
	flag.IntVar(&tokenTTL, "token-ttl", tokenTTL, "How many hours a JWT token for this user is valid.")

	flag.StringVar(&docOutput, "doc", docOutput, "Generate the API documentation to this file.")
	flag.StringVar(&docFormat, "format", docFormat, "Format of the generated documentation.")
}
