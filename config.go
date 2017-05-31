package cyu

import (
	"log"
	"path/filepath"
)

// Config defines configuration settings for Cyu.
type Config struct {
	BaseDir       string // directory prepended to non-absolute directories
	FeatureDir    string // Gherkin .feature files are found under this directory
	TestDir       string // Test output files are found under this directory
	StepMapDir    string // step map files are found under here
	TemplateDir   string // the particular templates are found here
	TemplateName  string // Use this base template
	TestsTypeName string // prepended to test files
	WorldName     string // the world (test context) class name, used to run the same steps in a different context
}

var DefaultConfig = Config{
	BaseDir:       ".",           // cwd absolutized
	FeatureDir:    "features",    //
	TestDir:       "tests",       //
	StepMapDir:    ".",           // base dir
	TemplateDir:   "",            // if blank, use templates from cyu package
	TemplateName:  "go",          // usually used for language selection
	TestsTypeName: "Feature",     // used in comments
	WorldName:     "TestContext", // class used to instantiate test world
}

func (c *Config) Configure(other *Config) {
	*c = *other
	c.CheckConfig()
}

func (c *Config) CheckConfig() {
	var err error

	fixup := func(s *string) {
		if !filepath.IsAbs(*s) {
			*s = filepath.Join(c.BaseDir, *s)
		}
	}

	if c.BaseDir, err = filepath.Abs(c.BaseDir); err != nil {
		log.Fatalf("Problem absolutizing base path '%s': %s", c.BaseDir, err)
	}

	fixup(&(c.FeatureDir))
	fixup(&(c.TestDir))
	fixup(&(c.StepMapDir))
	fixup(&(c.TemplateDir)) // this maybe should be based off the executable dir?

}
