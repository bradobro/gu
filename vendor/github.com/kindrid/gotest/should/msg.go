package should

import (
	"strings"
)

// tools for working with failure messages

const (
	// ShortSeparator ends the failure message short portion (and begins the long
	// portion)
	ShortSeparator = "\n"

	// ShortLength is an arbitrary length used to shorten failure messages that
	// don't contain ShortSeparator
	ShortLength = 80

	// LongSeparator ends the failure message explanation section (and begins the
	// details)
	LongSeparator = "\n# DETAILS:"

	// DetailsSeparator ends the failure debugging portion (and begins details for
	// debugging asserts and test runner.)
	DetailsSeparator = "\n# INTERNALS:"

	// SectionSeparator separates the long, details, and internals sections.
	SectionSeparator = "\n~~~~~~~~~~\n"
)

func trim(s string) string {
	return strings.Trim(s, " \n\t\r")
}

func splitShortLong(s string) (short, long string) {
	sl := strings.SplitN(trim(s), ShortSeparator, 2)
	if len(sl) > 1 {
		return trim(sl[0]), trim(sl[1])
	}
	if len(s) > ShortLength {
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
