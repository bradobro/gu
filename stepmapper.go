package cyu

// StepMapper is an interface providing a a mapping between different representations of steps.
type StepMapper interface {
	Match(prose string) *Step // first step with a pattern
}

// StepMap is our implementation of a StepMapper
type StepMap struct {
	Steps []*Step
}
