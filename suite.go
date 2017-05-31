//go:generate go build -v ./vendor/github.com/valyala/quicktemplate/qtc
//go:generate qtc -dir ./templates
package cyu

// Suite defines a group of features
type Suite struct {
	Config
	suites   []*Suite    // for sub-groups of tests
	features []*Feature  // features in the suite
	steps    interface{} //step map
}

// Attic Implementation
// func (su *Suite) Export() (err error) {
// 	if err = su.ReadStepMap(); err != nil {
// 		return
// 	}
// 	if err = su.ReadFeatureDir(); err != nil {
// 		return
// 	}
// 	if err = su.ExportFeatures(); err != nil {
// 		return
// 	}
// 	return
// }

// // ReadStepMap does stuff
// func (su *Suite) ReadStepMap() (err error) {
// 	return
// }

// // ReadFeatureDir does stuff
// func (su *Suite) ReadFeatureDir() (err error) {
// 	return filepath.Walk(
// 		su.FeaturePath,
// 		func(path string, info os.FileInfo, err error) error {
// 			fmt.Println(path)
// 			if strings.HasSuffix(path, ".feature") {
// 				if err := su.ReadFeature(path); err != nil {
// 					return err
// 				}
// 			}
// 			return nil
// 		},
// 	)
// }

// // ExportFeatures does stuff
// func (su *Suite) ExportFeatures() (err error) {
// 	return
// }

// // package features

// // import (
// // 	"fmt"
// // 	"os"
// // 	"testing"

// // 	gherkin "github.com/cucumber/gherkin-go"
// // 	"github.com/y0ssar1an/q"
// // )

// // ReadFeature reads a single feature file
// func (su *Suite) ReadFeature(featureFile string) (err error) {
// 	reader, err := os.Open(featureFile)
// 	defer reader.Close()
// 	if err != nil {
// 		err = fmt.Errorf("Unable to read feature file '%s': %s", featureFile, err)
// 		return
// 	}
// 	doc, err := gherkin.ParseGherkinDocument(reader)
// 	if err != nil {
// 		err = fmt.Errorf("Unable to parse feature file '%s': %s", featureFile, err)
// 		return
// 	}
// 	feature := &Feature{
// 		Path:    featureFile,
// 		Feature: doc.Feature,
// 	}
// 	features := append(su.features, feature)
// 	su.features = features
// 	return
// }

// // func ReadFeatureDir(featureDir string) (features []*Feature, err error) {
// // 	return
// // }

// // func AddFeatures(t *testing.T, files []string) {
// // 	for i, featureFile := range files {
// // 		feature, background, err := ReadFeatureFile(featureFile)
// // 		if err != nil {
// // 			t.Error(err.Error())
// // 		}
// // 		q.Q(feature)

// // 		j := 0
// // 		for _, scenario := range feature.Children {
// // 			switch s := scenario.(type) {
// // 			case *gherkin.Scenario:
// // 				// t.Run(fmt.Sprintf("f%d.%d", i, j), func(tScenario *testing.T) {
// // 				// 	TestScenario(tScenario, background, s)
// // 				// })
// // 				GenerateScenario(i, j, background, s)
// // 				j++
// // 			case *gherkin.ScenarioOutline:
// // 				// t.Run(fmt.Sprintf("f%d.%d", i, j), func(tScenario *testing.T) {
// // 				// 	TestScenarioOutline(tScenario, background, s)
// // 				// })
// // 				GenerateScenarioOutline(i, j, background, s)
// // 				j++
// // 			case *gherkin.Background:
// // 				// ignore
// // 			default:
// // 				t.Fatalf("Unknown feature child type in %s: %T", featureFile, scenario)
// // 			}
// // 		}
// // 	}
// // }

// func GenerateScenario(featureN, scenarioN int, background *gherkin.Background, scenario *gherkin.Scenario) {
// 	fmt.Printf("func TextF%dS%d(t *testing.T) {\n", featureN, scenarioN)
// 	fmt.Printf("  t.Log(%#v)\n", scenario.Name)
// 	if background != nil {
// 		fmt.Print("  callBackground(background)\n")
// 	}
// 	for _, step := range scenario.ScenarioDefinition.Steps {
// 		fmt.Printf("  callStep(%#v, %T)\n", step.Text, step.Argument)
// 	}
// 	fmt.Println("}")
// }

// func GenerateScenarioOutline(featureN, scenarioN int, background *gherkin.Background, scenario *gherkin.ScenarioOutline) {
// 	fmt.Printf("t.Log(%#v)\n", scenario.Name)
// }
