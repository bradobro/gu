package gu_test

import (
	"errors"
	"testing"

	"github.com/bradobro/gu"
)

func TestReporter(t *testing.T) {
	rpt := &gu.Reporter{
		Verbosity: gu.VerbosityInsane,
		MaxDepth:  3,
	}
	rpt.Log(t, "Log from the reporter says hello.")
	rpt.Logf(t, "Logf from the reporter says, %q.", "Hello")
}

const (
	shortMsg   = "Brief message"
	longMsg    = "Extra message with\nseveral lines."
	detailsMsg = "Here are some \ndetails"
	metaMsg    = "Here are a bunch of technical details about test workings \nfor when you doubt the assertion or the runner."
)

var (
	failShort     error = errors.New(shortMsg)
	failLong      error = errors.New(gu.ShortSeparator + longMsg)
	failDetails   error = errors.New(gu.SectionSeparator + detailsMsg)
	failMeta      error = errors.New(gu.SectionSeparator + metaMsg)
	failShortLong error = errors.New(shortMsg + gu.ShortSeparator + longMsg)
	failAll       error = errors.New(shortMsg + gu.ShortSeparator + longMsg +
		gu.SectionSeparator + detailsMsg + gu.SectionSeparator + metaMsg)
	failShortMeta error = errors.New(shortMsg + gu.ShortSeparator + gu.SectionSeparator + metaMsg)
)

func TestFailureMessageParsing(t *testing.T) {
	// test blank case
	rpt := &gu.Reporter{}
	s, l, d, m := rpt.Parse(nil)
	assertEquals(t, s, "")
	assertEquals(t, l, "")
	assertEquals(t, d, "")
	assertEquals(t, m, "")
	// test normal order
	testMessageParse(t, nil, "", "", "", "")
	testMessageParse(t, nil, shortMsg, "", "", "")
	testMessageParse(t, nil, shortMsg, longMsg, "", "")
	testMessageParse(t, nil, shortMsg, longMsg, detailsMsg, metaMsg)

	testMessageParse(t, errors.New("something"), "something", "", "", "")
	testMessageParse(t, errors.New("something\n"+longMsg), "something", longMsg, "", "")
	testMessageParse(t, errors.New(" \n \n \n "+longMsg), "Extra message with", "several lines.", "", "")
}

func testMessageParse(t *testing.T, err error, short, long, details, meta string) {
	rpt := &gu.Reporter{}
	if err == nil {
		err = gu.Failure(short, long, details, meta)
	}
	s, l, d, m := rpt.Parse(err)
	assertEquals(t, s, short)
	assertEquals(t, l, long)
	assertEquals(t, d, details)
	assertEquals(t, m, meta)
}

func TestReportLevels(t *testing.T) {
	rpt := &gu.Reporter{MaxDepth: 5}

	ct, buf := newTestT(t)
	rpt.Verbosity = gu.VerbositySilent
	rpt.Report(ct, 0, failAll, rpt, "extra 1", "extra 2")
	assertEquals(t, buf.String(), "")

	ct, buf = newTestT(t)
	rpt.Verbosity = gu.VerbosityShort
	rpt.Report(ct, 0, failAll, rpt, "extra 1", "extra 2")
	assertStringContains(t, buf.String(), "Brief message")

	ct, buf = newTestT(t)
	rpt.Verbosity = gu.VerbosityLong
	rpt.Report(ct, 0, failAll, rpt, "extra 1", "extra 2")
	assertStringContains(t, buf.String(), "Brief message\nEXTRA INFO: Extra message with\nseveral lines.\n")

	ct, buf = newTestT(t)
	rpt.Verbosity = gu.VerbosityActuals
	rpt.Report(ct, 0, failAll, rpt, "extra 1", "extra 2")
	assertStringContains(t, buf.String(), "gu.Reporter{Verbosity")

	ct, buf = newTestT(t)
	rpt.Verbosity = gu.VerbosityExpecteds
	rpt.Report(ct, 0, failAll, rpt, "extra 1", "extra 2")
	assertStringContains(t, buf.String(), `"extra 1", "extra 2"`)

	ct, buf = newTestT(t)
	rpt.Verbosity = gu.VerbosityDebug
	rpt.Report(ct, 1, failAll, rpt, "extra 1", "extra 2")
	assertStringContains(t, buf.String(), "Here are some \ndetails\n")
	assertStringContains(t, buf.String(), "gu/reporter_test.go")

	ct, buf = newTestT(t)
	rpt.Verbosity = gu.VerbosityInsane
	rpt.Report(ct, 0, failAll, rpt, "extra 1", "extra 2")
	assertStringContains(t, buf.String(), "Here are a bunch of technical details")
}

func TestStackReporting(t *testing.T) {

	// Frame count too high should not cause a panic
	f := gu.Frames(0, 1000)
	t.Log(f)
}
