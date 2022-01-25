package gomockx

import (
	"net/http"
	"strings"
)

type HttpRequestMatcher struct {
	expectedRequest *http.Request
}

func (matcher HttpRequestMatcher) Matches(arg interface{}) bool {
	switch r := arg.(type) {
	case *http.Request:
		if !strings.EqualFold(matcher.expectedRequest.Method, r.Method) {
			return false
		}

		return true
	default:
		return false
	}
}

func NewHttpRequestMatcher(expectedRequest *http.Request) HttpRequestMatcher {
	return HttpRequestMatcher{
		expectedRequest: expectedRequest,
	}
}
