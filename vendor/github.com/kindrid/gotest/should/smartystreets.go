package should

import smarty "github.com/smartystreets/assertions"

var (
	// AlmostEqual is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	AlmostEqual = smarty.ShouldAlmostEqual
	// BeBetween is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	BeBetween = smarty.ShouldBeBetween
	// BeBetweenOrEqual is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	BeBetweenOrEqual = smarty.ShouldBeBetweenOrEqual
	// BeBlank is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	BeBlank = smarty.ShouldBeBlank
	// BeChronological is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	BeChronological = smarty.ShouldBeChronological
	// BeEmpty is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	BeEmpty = smarty.ShouldBeEmpty
	// BeFalse is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	BeFalse = smarty.ShouldBeFalse
	// BeGreaterThan is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	BeGreaterThan = smarty.ShouldBeGreaterThan
	// BeGreaterThanOrEqualTo is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	BeGreaterThanOrEqualTo = smarty.ShouldBeGreaterThanOrEqualTo
	// BeIn is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	BeIn = smarty.ShouldBeIn
	// BeLessThan is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	BeLessThan = smarty.ShouldBeLessThan
	// BeLessThanOrEqualTo is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	BeLessThanOrEqualTo = smarty.ShouldBeLessThanOrEqualTo
	// BeNil is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	BeNil = smarty.ShouldBeNil
	// BeTrue is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	BeTrue = smarty.ShouldBeTrue
	// BeZeroValue is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	BeZeroValue = smarty.ShouldBeZeroValue
	// Contain is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	Contain = smarty.ShouldContain
	// ContainKey is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	ContainKey = smarty.ShouldContainKey
	// ContainSubstring is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	ContainSubstring = smarty.ShouldContainSubstring
	// EndWith is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	EndWith = smarty.ShouldEndWith
	// Equal is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	Equal = smarty.ShouldEqual
	// EqualTrimSpace is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	EqualTrimSpace = smarty.ShouldEqualTrimSpace
	// EqualWithout is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	EqualWithout = smarty.ShouldEqualWithout
	// HappenAfter is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	HappenAfter = smarty.ShouldHappenAfter
	// HappenBefore is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	HappenBefore = smarty.ShouldHappenBefore
	// HappenBetween is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	HappenBetween = smarty.ShouldHappenBetween
	// HappenOnOrAfter is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	HappenOnOrAfter = smarty.ShouldHappenOnOrAfter
	// HappenOnOrBefore is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	HappenOnOrBefore = smarty.ShouldHappenOnOrBefore
	// HappenOnOrBetween is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	HappenOnOrBetween = smarty.ShouldHappenOnOrBetween
	// HappenWithin is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	HappenWithin = smarty.ShouldHappenWithin
	// HaveLength is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	HaveLength = smarty.ShouldHaveLength
	// HaveSameTypeAs is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	HaveSameTypeAs = smarty.ShouldHaveSameTypeAs
	// Implement is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	Implement = smarty.ShouldImplement
	// NotAlmostEqual is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	NotAlmostEqual = smarty.ShouldNotAlmostEqual
	// NotBeBetween is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	NotBeBetween = smarty.ShouldNotBeBetween
	// NotBeBetweenOrEqual is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	NotBeBetweenOrEqual = smarty.ShouldNotBeBetweenOrEqual
	// NotBeBlank is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	NotBeBlank = smarty.ShouldNotBeBlank
	// NotBeEmpty is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	NotBeEmpty = smarty.ShouldNotBeEmpty
	// NotBeIn is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	NotBeIn = smarty.ShouldNotBeIn
	// NotBeNil is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	NotBeNil = smarty.ShouldNotBeNil
	// NotContain is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	NotContain = smarty.ShouldNotContain
	// NotContainKey is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	NotContainKey = smarty.ShouldNotContainKey
	// NotContainSubstring is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	NotContainSubstring = smarty.ShouldNotContainSubstring
	// NotEndWith is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	NotEndWith = smarty.ShouldNotEndWith
	// NotEqual is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	NotEqual = smarty.ShouldNotEqual
	// NotHappenOnOrBetween is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	NotHappenOnOrBetween = smarty.ShouldNotHappenOnOrBetween
	// NotHappenWithin is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	NotHappenWithin = smarty.ShouldNotHappenWithin
	// NotHaveSameTypeAs is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	NotHaveSameTypeAs = smarty.ShouldNotHaveSameTypeAs
	// NotImplement is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	NotImplement = smarty.ShouldNotImplement
	// NotPanic is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	NotPanic = smarty.ShouldNotPanic
	// NotPanicWith is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	NotPanicWith = smarty.ShouldNotPanicWith
	// NotPointTo is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	NotPointTo = smarty.ShouldNotPointTo
	// NotResemble is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	NotResemble = smarty.ShouldNotResemble
	// NotStartWith is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	NotStartWith = smarty.ShouldNotStartWith
	// Panic is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	Panic = smarty.ShouldPanic
	// PanicWith is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	PanicWith = smarty.ShouldPanicWith
	// PointTo is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	PointTo = smarty.ShouldPointTo
	// Resemble is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	Resemble = smarty.ShouldResemble
	// StartWith is imported from smartystreets/assertions. See https://gowalker.org/github.com/smartystreets/assertions
	StartWith = smarty.ShouldStartWith
)
