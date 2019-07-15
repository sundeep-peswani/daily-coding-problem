package main

import (
	"math/rand"
	"regexp"
	"testing"
	"time"
)

func TestUrlShortener(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	u := newUrlShortener()
	baseURL := "https://example.com"

	exampleComCode := u.shorten(baseURL)
	exampleComCodeAgain := u.shorten(baseURL)
	if exampleComCode != exampleComCodeAgain {
		t.Errorf("urlShortener.shorten(): duplicate ignored, expected %s, got %s\n", exampleComCode, exampleComCodeAgain)
	}

	differentExampleComCode := u.shorten(baseURL + "/")
	if exampleComCode == differentExampleComCode || exampleComCodeAgain == differentExampleComCode {
		t.Errorf("urlShortener.shorten(): non-duplicated marked as dupe. did not want %s or %s, got %s\n", exampleComCode, exampleComCodeAgain, differentExampleComCode)
	}

	if u.restore(exampleComCode) != baseURL {
		t.Errorf("urlShortener.restore(): unable to find %s, expected %s, got %s\n", exampleComCode, baseURL, u.restore(exampleComCode))
	}

	if u.restore(exampleComCodeAgain) != baseURL {
		t.Errorf("urlShortener.restore(): unable to find %s, expected %s, got %s\n", exampleComCodeAgain, baseURL, u.restore(exampleComCode))
	}

	if u.restore(differentExampleComCode) != baseURL+"/" {
		t.Errorf("urlShortener.restore(): unable to find %s, expected %s, got %s\n", differentExampleComCode, baseURL, u.restore(exampleComCode))
	}

	if u.restore("faked.") != "" {
		t.Errorf("urlShortener.restore(): able to restore fake code, got %s\n", u.restore("faked."))
	}
}

func TestGenerateCode(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	u := newUrlShortener()
	re := regexp.MustCompile("[^a-zA-Z0-9]")

	for i := 0; i < 10000; i++ {
		actual := u.generateCode()
		if len(re.FindString(actual)) > 0 {
			t.Errorf("urlShortener.generateCode(): found non-alphanumeric character: %s\n", actual)
		}
	}
}
