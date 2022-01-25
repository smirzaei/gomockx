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

func createSampleRequest(reqBody string) (*http.Request, error) {
	bodyReader := bytes.NewReader([]byte(reqBody))
	return http.NewRequest(http.MethodPost, "https://httpbin.org/get", bodyReader)
}

func TestHttpRequestMatcherToReturnFalseIfArgIsOfADifferentType(t *testing.T) {
	// Arrange
	req, err := createSampleRequest(sampleBody)
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
	req1, err := createSampleRequest(sampleBody)
	if err != nil {
		t.Fatal(err)
	}

	req2, err := createSampleRequest(sampleBody)
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
	req1, err := createSampleRequest(sampleBody)
	if err != nil {
		t.Fatal(err)
	}

	req2, err := createSampleRequest(sampleBody)
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
	req1, err := createSampleRequest(sampleBody)
	if err != nil {
		t.Fatal(err)
	}

	req2, err := createSampleRequest(sampleBody)
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
	req1, err := createSampleRequest(sampleBody)
	if err != nil {
		t.Fatal(err)
	}

	req2, err := createSampleRequest("a different body")
	if err != nil {
		t.Fatal(err)
	}

	matcher := NewHttpRequestMatcher(req1)

	// Act
	res := matcher.Matches(req2)

	// Assert
	assert.False(t, res)
}

func TestHttpRequestMatcherToReturnTrueIfEverythingElseIsTheSame(t *testing.T) {
	// Arrange
	req1, err := createSampleRequest(sampleBody)
	if err != nil {
		t.Fatal(err)
	}

	req2, err := createSampleRequest(sampleBody)
	if err != nil {
		t.Fatal(err)
	}

	matcher := NewHttpRequestMatcher(req1)

	// Act
	res := matcher.Matches(req2)

	// Assert
	assert.True(t, res)
}
