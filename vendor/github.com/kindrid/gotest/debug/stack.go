package debug

import (
	"fmt"
	"path"
	"runtime"
	"strings"
)

// // CallerInfo gives info about the current call stack.
// func CallerInfo(depth int) (msg, fileName string, fileLine int) {
// 	programCoun, fileName, fileLine, ok := runtime.Caller(depth)
// 	if ok {
// 		msg = fmt.Sprintf("%s:%d", fileName, fileLine)
// 	}
// 	return
// }
//
// // CallerSimple gives a single string with condensed information about the current call stack.
// func CallerSimple(depth int) string {
// 	msg, _, _ := CallerInfo(depth)
// 	return msg
// }

const minSkip = 1

// formatFrames applies a function to each level of call frame.
func formatFrames(
	startDepth,
	maxDepth int,
	formatter func(fullpath string, line int) string,
) (result []string) {
	if startDepth < minSkip {
		startDepth = minSkip // at least ignore this code and the caller
	}
	result = make([]string, maxDepth)
	for i := startDepth; i < maxDepth+startDepth; i++ {
		_, fullpath, line, ok := runtime.Caller(i)
		if !ok {
			continue
		}
		result = append(result, formatter(fullpath, line))
	}
	return
}

// // CallStack gives a short array of the call stack descriptions.
// func LongCallStack(startDepth, maxDepth int) (result []string) {
// 	if startDepth < 3 {
// 		startDepth = 3 // at least ignore this code and the caller
// 	}
// 	result = make([]string, maxDepth)
// 	for i := startDepth; i < maxDepth+startDepth; i++ {
// 		msg, _, _ := CallerInfo(i)
// 		if msg == "" {
// 			break
// 		}
// 		result = append(result, msg)
// 	}
// 	return
// }

// ShortStack tries to give one line of info that can help locate a bug.
func ShortStack(startDepth, maxDepth int) string {
	formatter := func(fullpath string, line int) string {
		file := path.Base(fullpath)
		return fmt.Sprintf("%s:%d", file, line)
	}
	result := strings.Join(formatFrames(startDepth, maxDepth, formatter), ",")
	return strings.Trim(result, " \n\r\t,")
}

// FormattedCallStack returns the call stack printout as lines.
func FormattedCallStack(startDepth, maxDepth int) string {
	formatter := func(fullpath string, line int) string {
		return fmt.Sprintf("%s:%d", fullpath, line)
	}
	result := strings.Join(formatFrames(startDepth, maxDepth, formatter), "\n")
	return strings.Trim(result, " \n\r\t")
}
