package main

import (
	"flag"

	"github.com/bradobro/cyu"
	"github.com/y0ssar1an/q"
)

var usage = `%s - generate or update tests for multiple languages from a set of feature files

`

func readFlags() *cyu.Config {
	opts := cyu.Config{}
	flag.StringVar(&(opts.BaseDir), "base", cyu.DefaultConfig.BaseDir, "directory prepended to non-absolute directories")
	flag.StringVar(&(opts.FeatureDir), "features", cyu.DefaultConfig.FeatureDir, "Gherkin .feature files are found under this directory")
	flag.StringVar(&(opts.TestDir), "tests", cyu.DefaultConfig.TestDir, "Test output files are found under this directory")
	flag.StringVar(&(opts.StepMapsDir), "stepmaps", cyu.DefaultConfig.StepMapsDir, "step map files are found under here")
	flag.StringVar(&(opts.TemplatesDir), "templates", cyu.DefaultConfig.TemplatesDir, "the particular templates are found here")
	flag.StringVar(&(opts.TemplateName), "template", cyu.DefaultConfig.TemplateName, "Use this base template")
	flag.StringVar(&(opts.TestsTypeName), "testtype", cyu.DefaultConfig.TestsTypeName, "prepended to test files")
	flag.StringVar(&(opts.WorldName), "world", cyu.DefaultConfig.WorldName, "the world (test context) class name, used to run the same steps in a different context")
	flag.Parse()
	return &opts
}

func main() {
	suite := cyu.Suite{}
	suite.Configure(readFlags())
	// suite = config.(cyu.Suite) // embed the config into the suite
	q.Q(suite)
	// err := suite.Export()
	// if err != nil {
	// 	log.Fatalf("Error processing gherkin: %s", err)
}
