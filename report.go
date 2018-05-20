package cyu

import (
	"fmt"
	"strings"
	// "github.com/kindrid/gotest/debug"
)

// Verbosity Constants: these are conventions only. Assertions and test
// functions can interpret these however they want.
const (
	// VerbositySilent shows no output except panics
	VerbositySilent = iota - 1
	// VerbosityShort, by convention, shows the first line of the failure string.
	VerbosityShort
	// VerbosityLong adds the first paragraph of failure string (up to SectionSeparator)
	VerbosityLong
	// VerbosityActuals adds the entire failure string and a (possibly shortened) representation of the actual value
	VerbosityActuals
	// VerbosityExpecteds adds (possibly shortened) representation(s) of expected values.
	VerbosityExpecteds
	// VerbosityDebug adds granular information to successes as well as failures.
	VerbosityDebug
	// VerbosityInsane Adds information to test meta concerns, such as logic within assertions.
	VerbosityInsane
)

// Message convention constants
const (
	// ShortSeparator ends the failure message short portion (and begins the long
	// portion)
	ShortSeparator = "\n"

	// ShortLength is an arbitrary length used to shorten failure messages that
	// don't contain ShortSeparator
	ShortLength = 80

	// SectionSeparator separates the long, details, and internals sections.
	SectionSeparator = "\n~~~~~~~~~~\n"
)

func trim(s string) string {
	return strings.Trim(s, " \n\t\r")
}

func splitShortLong(s string) (short, long string) {
	sl := strings.SplitN(trim(s), ShortSeparator, 2)
	if len(sl) > 1 {
		short = trim(sl[0])
		long = trim(sl[1])
	} else {
		short = trim(s)
	}
	// if len(short) > ShortLength {
	// 	short = short[:ShortLength]
	// }
	return
}

// Failure creates a failure message from its components
func Failure(short, long, details, meta string) (result string) {
	result = short + ShortSeparator + long + SectionSeparator + details + SectionSeparator + meta
	return
}

type Reporter struct {
	T         T
	Verbosity int
	MaxDepth  int
}

func (r *Reporter) Logf(format string, args ...interface{}) {
	r.T.Logf(format, args...)
}
func (r *Reporter) Log(message string) {
	r.T.Logf("%s", message)
}

// Parse divides a failure message into parts that may be muted depending on verbosity levels
func (r *Reporter) Parse(msg string) (short, long, details, meta string) {
	if msg == "" {
		return
	}
	secs := strings.Split(msg, SectionSeparator)
	short, long = splitShortLong(secs[0])
	if len(secs) > 1 {
		details = trim(secs[1])
	}
	if len(secs) > 2 {
		meta = trim(secs[2])
	}
	return
}

func (r *Reporter) Report(skip int, fail string, params ...interface{}) {
	var msg string
	terseMsg, extraMsg, detailsMsg, metaMsg := r.Parse(fail)
	// if StackDepth > 0 {
	// msg += fmt.Sprintf("\nTest Failure Stack Trace: %s\n\n", debug.FormattedCallStack(StackLevel, StackDepth))
	// }
	msg += r.Sprintv(VerbosityShort, "%s", "FAILED")
	if namer, ok := r.T.(Namer); ok { // this should allow use in Go < 1.8
		msg += r.Sprintv(VerbosityShort, " %s:", namer.Name())
	}
	msg += r.Sprintv(VerbosityShort, "%s", terseMsg)
	msg += r.Sprintv(VerbosityLong, "\nEXTRA INFO: %s\n", extraMsg)
	if len(params) > 0 {
		msg += r.Inspectv(VerbosityActuals, "\nACTUAL", params[0])
	}
	if len(params) > 1 {
		msg += r.Inspectv(VerbosityExpecteds, "\nEXPECTED", params[1:])
	}
	if detailsMsg != "" {
		msg += r.Sprintv(VerbosityDebug, "\nDETAILS: %s\n", detailsMsg)
	}
	if metaMsg != "" {
		msg += r.Sprintv(VerbosityInsane, "\nINTERNALS (FOR DEBUGGING ASSERTIONS): %s\n", metaMsg)
	}
	r.Log(msg)
}

// Sprintv formats a string if Verbosity >= minLevel, otherwise returns ""
func (r *Reporter) Sprintv(minLevel int, format string, args ...interface{}) string {
	if r.Verbosity < minLevel {
		return ""
	}
	return fmt.Sprintf(format, args...)
}

// Inspectv returns a detailed introspection of objects if Verbosity >= minLevel.
func (r *Reporter) Inspectv(minLevel int, label string, inspected ...interface{}) (result string) {
	if r.Verbosity < minLevel {
		return
	}
	if label != "" {
		result = fmt.Sprintf("%s: \n", label)
	}
	for _, x := range inspected {
		result += fmt.Sprintf("%#v\n", x)
	}
	return
}
