package should

import "testing"

const (
	shortMsg      = "Brief message"
	longMsg       = "Extra message with\nseveral lines."
	detailsMsg    = "Here are some \ndetails"
	metaMsg       = "Here are a bunch of technical details about test workings \nfor when you doubt the assertion or the runner."
	failShort     = shortMsg
	failLong      = ShortSeparator + longMsg
	failDetails   = SectionSeparator + detailsMsg
	failMeta      = SectionSeparator + metaMsg
	failShortLong = failShort + failLong
	failAll       = failShort + failLong + failDetails + failMeta

	failShortMeta = failShort + ShortSeparator + SectionSeparator + failMeta
)

func testStringEqual(t *testing.T, actual, expected string) {
	if actual != expected {
		t.Errorf("Actual %#v\nExpected %#v", actual, expected)
	}
}

func testMessageParse(t *testing.T, msg, short, long, details, meta string) {
	if msg == "" {
		msg = FormatFailure(short, long, details, meta)
	}
	s, l, d, m := ParseFailure(msg)
	testStringEqual(t, s, short)
	testStringEqual(t, l, long)
	testStringEqual(t, d, details)
	testStringEqual(t, m, meta)
}

func skipTestFailureMessageCreation(t *testing.T) {
	// Test everything and nothing
	testStringEqual(t, FormatFailure(shortMsg, longMsg, detailsMsg, metaMsg), failAll)
	testStringEqual(t, FormatFailure("", "", "", ""), ShortSeparator+SectionSeparator+SectionSeparator)

	// Test combinations
	testStringEqual(t, FormatFailure(shortMsg, "", "", ""), failShort+ShortSeparator+SectionSeparator+SectionSeparator)
	testStringEqual(t, FormatFailure(shortMsg, "", detailsMsg, ""), failShort+ShortSeparator+failDetails+SectionSeparator)
	testStringEqual(t, FormatFailure(shortMsg, "", "", metaMsg), failShortMeta)
}

func TestFailureMessageParsing(t *testing.T) {
	// test normal order
	testMessageParse(t, "", "", "", "", "")
	testMessageParse(t, "", shortMsg, "", "", "")
	testMessageParse(t, "", shortMsg, longMsg, "", "")
	testMessageParse(t, "", shortMsg, longMsg, detailsMsg, metaMsg)

	// // short grabs everything if the msg seems non-compliant
	testMessageParse(t, "", shortMsg, "", "", "")
	testMessageParse(t, "something\n"+longMsg, "something", longMsg, "", "")
	testMessageParse(t, " \n \n \n "+longMsg, "Extra message with", "several lines.", "", "")
}
