package gomockx

import (
	"net/http"
	"reflect"
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

		if matcher.expectedRequest.URL.String() != r.URL.String() {
			return false
		}

		if !reflect.DeepEqual(matcher.expectedRequest.Header, r.Header) {
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
