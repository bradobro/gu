package should

import (
	"fmt"
	"net/http"
	"strings"
)

// DescribeResponse gives a short-ish report of a response
func DescribeResponse(rsp *http.Response) (desc string) {
	template := `%d %s
%s
-----body----
%s`
	var headers []string
	for k, v := range rsp.Header {
		headers = append(headers, fmt.Sprintf("%15s: %s", k, v))
	}
	body := rsp.Body
	desc = fmt.Sprintf(template, rsp.StatusCode, rsp.Status, headers, body)
	return
}

// DescribeRequest gives a short-ish report of a request.
func DescribeRequest(req *http.Request) (desc string) {
	var headers []string
	for k, v := range req.Header {
		headers = append(headers, fmt.Sprintf("%15s: %s", k, v))
	}
	body := req.Body
	template := `%s %s
%s
-----body----
%s`
	desc = fmt.Sprintf(template, req.Method, req.URL, strings.Join(headers, "\n"), body)
	return
}

// MatchHTTPStatusCode asserts that the documented and actual HTTP status codes match
func MatchHTTPStatusCode(actual interface{}, expected ...interface{}) (fail string) {
	if msg := exactly(1, expected); msg != Ok {
		return msg
	}
	eStatus, ok := expected[0].(int)
	if !ok {
		return fmt.Sprintf("Expected value should be an int, not a %T.", expected[0])
	}
	aRsp, ok := actual.(*http.Response)
	if !ok {
		return "Actual value should be a *http.Response"
	}
	aStatus := aRsp.StatusCode
	if Equal(aStatus, eStatus) == "" {
		return ""
	}

	short := fmt.Sprintf(
		"HTTP Status Expected: %d %s. Got: %d %s.",
		eStatus, http.StatusText(int(eStatus)),
		aStatus, http.StatusText(aStatus),
	)

	// details formatted request

	return FormatFailure(short, DescribeRequest(aRsp.Request), "", "")
}
