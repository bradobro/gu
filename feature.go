package cyu

import gherkin "github.com/cucumber/gherkin-go"

type Feature struct {
	Path    string           // file holding the feature source
	Feature *gherkin.Feature // AST of the feature
}
