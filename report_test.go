package cyu_test

import (
	"testing"

	"github.com/bradobro/cyu"
)

func TestReporter(t *testing.T) {
	rpt := &cyu.Reporter{
		T:         t,
		Verbosity: cyu.VerbosityInsane,
		MaxDepth:  3,
	}
	rpt.Log("Log from the reporter says hello.")
	rpt.Logf("Logf from the reporter says, %q.", "Hello")

}

const (
	shortMsg      = "Brief message"
	longMsg       = "Extra message with\nseveral lines."
	detailsMsg    = "Here are some \ndetails"
	metaMsg       = "Here are a bunch of technical details about test workings \nfor when you doubt the assertion or the runner."
	failShort     = shortMsg
	failLong      = cyu.ShortSeparator + longMsg
	failDetails   = cyu.SectionSeparator + detailsMsg
	failMeta      = cyu.SectionSeparator + metaMsg
	failShortLong = failShort + failLong
	failAll       = failShort + failLong + failDetails + failMeta

	failShortMeta = failShort + cyu.ShortSeparator + cyu.SectionSeparator + failMeta
)

func TestFailureMessageParsing(t *testing.T) {
	// test blank case
	rpt := &cyu.Reporter{}
	s, l, d, m := rpt.Parse("")
	testStringEqual(t, s, "")
	testStringEqual(t, l, "")
	testStringEqual(t, d, "")
	testStringEqual(t, m, "")
	// test normal order
	testMessageParse(t, "", "", "", "", "")
	testMessageParse(t, "", shortMsg, "", "", "")
	testMessageParse(t, "", shortMsg, longMsg, "", "")
	testMessageParse(t, "", shortMsg, longMsg, detailsMsg, metaMsg)

	testMessageParse(t, "something", "something", "", "", "")
	testMessageParse(t, "something\n"+longMsg, "something", longMsg, "", "")
	testMessageParse(t, " \n \n \n "+longMsg, "Extra message with", "several lines.", "", "")
}

func testStringEqual(t *testing.T, actual, expected string) {
	if actual != expected {
		t.Errorf("Actual %#v\nExpected %#v", actual, expected)
	}
}

func testMessageParse(t *testing.T, msg, short, long, details, meta string) {
	rpt := &cyu.Reporter{}
	if msg == "" {
		msg = cyu.Failure(short, long, details, meta)
	}
	t.Log("")
	s, l, d, m := rpt.Parse(msg)
	testStringEqual(t, s, short)
	testStringEqual(t, l, long)
	testStringEqual(t, d, details)
	testStringEqual(t, m, meta)
}

func TestReportLevels(t *testing.T) {
	// IWBNI test output could be redirected to a buffer, but the only
	//
	rpt := &cyu.Reporter{T: t}
	t.Log("Expect increasingly verbose logs:")
	for i := cyu.VerbositySilent; i <= cyu.VerbosityInsane; i++ {
		rpt.Verbosity = i
		t.Logf("Verbosity = %d", i)
		rpt.Report(cyu.VerbosityInsane, failAll, rpt, "extra 1", "extra 2")
	}
}
