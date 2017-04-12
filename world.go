package gocu

import "testing"

// World defines a test context
type World interface {
	BeforeSuite(t *testing.T) error
	AfterSuite(t *testing.T) error
	BeforeFeature(t *testing.T) error
	AfterFeature(t *testing.T) error
	BeforeScenario(t *testing.T) error
	AfterScenario(t *testing.T) error

	// Internal Key Value Stores
	Clear()
	SetString(name, value string)
	GetString(name string) string
	SetInt(name string, value int)
	GetInt(name string) int
	SetValue(name string, value interface{})
	GetValue(name string) interface{}
	SetTable(name string, value Table)
	GetTable(name string) Table
}
