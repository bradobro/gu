package main

import (
	"flag"
	"log"

	"github.com/bradobro/gocu"
	"github.com/y0ssar1an/q"
)

var usage = `%s - generate cucumber tests in go from a set of feature files

`

var defaults = gocu.Suite{
	FeaturePath: "features",
	Config:      "gocu.json",
	Register:    "RegisterSteps",
}

func readFlags() *gocu.Suite {
	opts := gocu.Suite{}
	flag.StringVar(&(opts.FeaturePath), "features", defaults.FeaturePath, "base directory for feature files")
	flag.StringVar(&(opts.Config), "config", defaults.Config, "config file name")
	// flag.StringVar(&(opts.stepMap), "map", defaults.stepMap, "step map file name")
	flag.Parse()
	q.Q(opts)
	return &opts
}

func main() {
	suite := readFlags()
	q.Q(suite)
	err := suite.Export()
	if err != nil {
		log.Fatalf("Error processing gherkin: %s", err)
	}
}
