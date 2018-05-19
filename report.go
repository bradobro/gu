package cyu

import (
	"fmt"
	"strings"

	"github.com/kindrid/gotest/debug"
	"github.com/kindrid/gotest/should"
)

// Verbosity Constants: these are conventions only. Assertions and test
// functions can interpret these however they want.
const (
	// Silent shows no output except panics
	Silent = iota - 1
	// Short, by convention, shows the first line of the failure string.
	Short
	// Long adds the first paragraph of failure string (up to SectionSeparator)
	Long
	// Actuals adds the entire failure string and a (possibly shortened) representation of the actual value
	Actuals
	// Expecteds adds  (possibly shortened) representation(s) of expected values.
	Expecteds
	// Debug Adds granular information to successes as well as failures.
	Debug
	// Insane Adds information to test meta concerns, such as logic within assertions.
	Insane
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
		return trim(sl[0])[:ShortLength], trim(sl[1])
	}
	if len(s) > ShortLength { // message too long, so return rest as long portion
		return trim(s[:ShortLength]), trim(s[ShortLength:])
	}
	return trim(s), ""
}

// ParseFailure divides a failure message into parts that may be muted depending on verbosity levels
func ParseFailure(msg string) (short, long, details, meta string) {
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

// FormatFailure creates a failure message from its components
func FormatFailure(short, long, details, meta string) (result string) {
	result = short + ShortSeparator + long + SectionSeparator + details + SectionSeparator + meta
	return
}

type Reporter struct {
	T         T
	Verbosity int
	MaxDepth  int
}

func (r *Reporter) Writef(format string, args ...interface{}) {
	r.T.Logf(format, args...)
}

func (r *Reporter) Report(skip int, fail string, params ...interface{}) {
	var msg string
	terseMsg, extraMsg, detailsMsg, metaMsg := should.ParseFailure(fail)
	// if StackDepth > 0 {
	// 	msg += fmt.Sprintf("\nTest Failure Stack Trace: %s\n\n", debug.FormattedCallStack(StackLevel, StackDepth))
	// }
	if namer, ok := r.T.(Namer); ok {
		msg += r.Sprintv(Short, "FAILED %s: %s", namer.Name(), terseMsg)
	} else {
		msg += r.Sprintv(Short, "FAILED: %s", terseMsg)

	}
	msg += r.Sprintv(Long, "\nEXTRA INFO: %s\n", extraMsg+"\nCalls:"+debug.ShortStack(3, 10))
	if len(params) > 0 {
		msg += r.Inspectv(Actuals, "\nACTUAL", params[0])
	}
	if len(params) > 1 {
		msg += r.Inspectv(Expecteds, "\nEXPECTED", expected)
	}
	if detailsMsg != "" {
		msg += r.Sprintv(Debug, "\nDETAILS: %s\n", detailsMsg)
	}
	if metaMsg != "" {
		msg += r.Sprintv(Insane, "\nINTERNALS (FOR DEBUGGING ASSERTIONS): %s\n", metaMsg)
	}
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
