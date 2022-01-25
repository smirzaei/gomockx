package gomockx

import (
	"fmt"
	"io"
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

		// This is not the most efficient way but good enough for now.
		expectedBody, err := io.ReadAll(matcher.expectedRequest.Body)
		if err != nil {
			fmt.Printf("warning: failed to read expected request body, because: %+v\n", err)
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("warning: failed to read request body, because: %+v\n", err)
		}

		if !reflect.DeepEqual(expectedBody, body) {
			return false
		}

		return true
	default:
		return false
	}
}

func (matcher HttpRequestMatcher) String() string {
	return fmt.Sprintf("request is identical to %v", matcher.expectedRequest)
}

func NewHttpRequestMatcher(expectedRequest *http.Request) HttpRequestMatcher {
	return HttpRequestMatcher{
		expectedRequest: expectedRequest,
	}
}
