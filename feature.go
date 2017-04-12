package gocu

import gherkin "github.com/cucumber/gherkin-go"

type Feature struct {
	Path    string
	Feature *gherkin.Feature
}
