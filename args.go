package main

import (
	"flag"
	"os"
)

type Args struct {
	verboseLevel string
	args         []string
}

func getEnvOrDefault(name, def string) string {
	if value, ok := os.LookupEnv(name); ok {
		return value
	}
	return def
}

func New() *Args {
	return new(Args).Parse()
}

func (a *Args) Parse() *Args {
	flag.StringVar(&a.verboseLevel, "verbose", getEnvOrDefault("VERBOSE", "info"), "Set verbosity level")
	flag.Parse()
	a.args = flag.Args()
	return a
}
