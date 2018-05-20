package cyu_test

import (
	"testing"

	"github.com/bradobro/cyu"
)

func TestReporter(t *testing.T) {
	rpt := &cyu.Reporter{
		Verbosity: cyu.VerbosityInsane,
		MaxDepth:  3,
	}
	rpt.Log(t, "Log from the reporter says hello.")
	rpt.Logf(t, "Logf from the reporter says, %q.", "Hello")
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
	assertEquals(t, s, "")
	assertEquals(t, l, "")
	assertEquals(t, d, "")
	assertEquals(t, m, "")
	// test normal order
	testMessageParse(t, "", "", "", "", "")
	testMessageParse(t, "", shortMsg, "", "", "")
	testMessageParse(t, "", shortMsg, longMsg, "", "")
	testMessageParse(t, "", shortMsg, longMsg, detailsMsg, metaMsg)

	testMessageParse(t, "something", "something", "", "", "")
	testMessageParse(t, "something\n"+longMsg, "something", longMsg, "", "")
	testMessageParse(t, " \n \n \n "+longMsg, "Extra message with", "several lines.", "", "")
}

func testMessageParse(t *testing.T, msg, short, long, details, meta string) {
	rpt := &cyu.Reporter{}
	if msg == "" {
		msg = cyu.Failure(short, long, details, meta)
	}
	s, l, d, m := rpt.Parse(msg)
	assertEquals(t, s, short)
	assertEquals(t, l, long)
	assertEquals(t, d, details)
	assertEquals(t, m, meta)
}

func TestReportLevels(t *testing.T) {
	rpt := &cyu.Reporter{MaxDepth: 5}

	ct, buf := newTestT(t)
	rpt.Verbosity = cyu.VerbositySilent
	rpt.Report(ct, 0, failAll, rpt, "extra 1", "extra 2")
	assertEquals(t, buf.String(), "")

	ct, buf = newTestT(t)
	rpt.Verbosity = cyu.VerbosityShort
	rpt.Report(ct, 0, failAll, rpt, "extra 1", "extra 2")
	testStringContains(t, buf.String(), "Brief message")

	ct, buf = newTestT(t)
	rpt.Verbosity = cyu.VerbosityLong
	rpt.Report(ct, 0, failAll, rpt, "extra 1", "extra 2")
	testStringContains(t, buf.String(), "Brief message\nEXTRA INFO: Extra message with\nseveral lines.\n")

	ct, buf = newTestT(t)
	rpt.Verbosity = cyu.VerbosityActuals
	rpt.Report(ct, 0, failAll, rpt, "extra 1", "extra 2")
	testStringContains(t, buf.String(), "cyu.Reporter{Verbosity")

	ct, buf = newTestT(t)
	rpt.Verbosity = cyu.VerbosityExpecteds
	rpt.Report(ct, 0, failAll, rpt, "extra 1", "extra 2")
	testStringContains(t, buf.String(), `"extra 1", "extra 2"`)

	ct, buf = newTestT(t)
	rpt.Verbosity = cyu.VerbosityDebug
	rpt.Report(ct, 1, failAll, rpt, "extra 1", "extra 2")
	testStringContains(t, buf.String(), "Here are some \ndetails\n")
	testStringContains(t, buf.String(), "cyu/report_test.go")

	ct, buf = newTestT(t)
	rpt.Verbosity = cyu.VerbosityInsane
	rpt.Report(ct, 0, failAll, rpt, "extra 1", "extra 2")
	testStringContains(t, buf.String(), "Here are a bunch of technical details")
}
