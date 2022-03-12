package gomockx

import (
	"bytes"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	sampleBody = "foo=bar"
)

func createSampleGetRequestWithoutBody() (*http.Request, error) {
	return http.NewRequest(http.MethodGet, "https://httpbin.org", nil)
}

func createSampleRequestWithBody(reqBody string) (*http.Request, error) {
	bodyReader := bytes.NewReader([]byte(reqBody))
	return http.NewRequest(http.MethodPost, "https://httpbin.org", bodyReader)
}

func TestHttpRequestMatcherToReturnFalseIfArgIsOfADifferentType(t *testing.T) {
	// Arrange
	req, err := createSampleRequestWithBody(sampleBody)
	if err != nil {
		t.Fatal(err)
	}

	matcher := NewHttpRequestMatcher(req)

	// Act
	res := matcher.Matches("a different type")

	// Assert
	assert.False(t, res)
}

func TestHttpRequestMatcherToReturnFalseIfMethodsAreNotTheSame(t *testing.T) {
	// Arrange
	req1, err := createSampleRequestWithBody(sampleBody)
	if err != nil {
		t.Fatal(err)
	}

	req2, err := createSampleRequestWithBody(sampleBody)
	if err != nil {
		t.Fatal(err)
	}

	req2.Method = http.MethodPut

	matcher := NewHttpRequestMatcher(req1)

	// Act
	res := matcher.Matches(req2)

	// Assert
	assert.False(t, res)
}

func TestHttpRequestMatcherToReturnFalseIfURLsAreNotTheSame(t *testing.T) {
	// Arrange
	req1, err := createSampleRequestWithBody(sampleBody)
	if err != nil {
		t.Fatal(err)
	}

	req2, err := createSampleRequestWithBody(sampleBody)
	if err != nil {
		t.Fatal(err)
	}

	newUrl, err := url.Parse("https://go.dev")
	if err != nil {
		t.Fatal(err)
	}

	req2.URL = newUrl

	matcher := NewHttpRequestMatcher(req1)

	// Act
	res := matcher.Matches(req2)

	// Assert
	assert.False(t, res)
}

func TestHttpRequestMatcherToReturnFalseIfHeadersDoNotMatch(t *testing.T) {
	// Arrange
	req1, err := createSampleRequestWithBody(sampleBody)
	if err != nil {
		t.Fatal(err)
	}

	req2, err := createSampleRequestWithBody(sampleBody)
	if err != nil {
		t.Fatal(err)
	}

	req2.Header.Add("x-test", "foo")

	matcher := NewHttpRequestMatcher(req1)

	// Act
	res := matcher.Matches(req2)

	// Assert
	assert.False(t, res)
}

func TestHttpRequestMatcherToReturnFalseIfBodyDoNotMatch(t *testing.T) {
	// Arrange
	req1, err := createSampleRequestWithBody(sampleBody)
	if err != nil {
		t.Fatal(err)
	}

	req2, err := createSampleRequestWithBody("a different body")
	if err != nil {
		t.Fatal(err)
	}

	matcher := NewHttpRequestMatcher(req1)

	// Act
	res := matcher.Matches(req2)

	// Assert
	assert.False(t, res)
}

func TestHttpRequestMatcherToReturnFalseIfTheExpectedReqHasABodyButTheActualDoesNot(t *testing.T) {
	// Arrange
	req1, err := createSampleRequestWithBody(sampleBody)
	if err != nil {
		t.Fatal(err)
	}

	req2, err := createSampleGetRequestWithoutBody()
	if err != nil {
		t.Fatal(err)
	}

	req2.Method = req1.Method

	matcher := NewHttpRequestMatcher(req1)

	// Act
	res := matcher.Matches(req2)

	// Assert
	assert.False(t, res)
}

func TestHttpRequestMatcherToReturnFalseIfTheExpectedReqDoesNotHaveABodyButTheActualDoes(t *testing.T) {
	// Arrange
	req1, err := createSampleGetRequestWithoutBody()
	if err != nil {
		t.Fatal(err)
	}

	req2, err := createSampleRequestWithBody(sampleBody)
	if err != nil {
		t.Fatal(err)
	}

	req2.Method = req1.Method

	matcher := NewHttpRequestMatcher(req1)

	// Act
	res := matcher.Matches(req2)

	// Assert
	assert.False(t, res)
}

func TestHttpRequestMatcherToReturnTrueIfEverythingElseIsTheSame(t *testing.T) {
	// Arrange
	req1, err := createSampleRequestWithBody(sampleBody)
	if err != nil {
		t.Fatal(err)
	}

	req2, err := createSampleRequestWithBody(sampleBody)
	if err != nil {
		t.Fatal(err)
	}

	matcher := NewHttpRequestMatcher(req1)

	// Act
	res := matcher.Matches(req2)

	// Assert
	assert.True(t, res)
}

func TestHttpRequestMatcherToReturnTrueIfURLAndHeadersMatchAndBodyIsEmpty(t *testing.T) {
	// Arrange
	req1, err := createSampleGetRequestWithoutBody()
	if err != nil {
		t.Fatal(err)
	}

	req2, err := createSampleGetRequestWithoutBody()
	if err != nil {
		t.Fatal(err)
	}

	matcher := NewHttpRequestMatcher(req1)

	// Act
	res := matcher.Matches(req2)

	// Assert
	assert.True(t, res)
}
