package main

import (
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
}

type urlShortener struct {
	urlToCode map[string]string
	codeToUrl map[string]string
}

func newUrlShortener() *urlShortener {
	var u urlShortener
	u.urlToCode = make(map[string]string)
	u.codeToUrl = make(map[string]string)

	return &u
}

func (u *urlShortener) shorten(url string) string {
	if code, ok := u.urlToCode[url]; ok {
		return code
	}

	code := u.generateCode()
	u.urlToCode[url] = code
	u.codeToUrl[code] = url

	return code
}

func (u *urlShortener) restore(code string) string {
	if url, ok := u.codeToUrl[code]; ok {
		return url
	}

	return ""
}

func (u *urlShortener) generateCode() string {
	var code strings.Builder

	for i := 0; i < 6; i++ {
		switch rand.Intn(3) {
		case 0:
			code.WriteByte(byte('a' + rand.Intn(26)))
		case 1:
			code.WriteByte(byte('A' + rand.Intn(26)))
		case 2:
			code.WriteByte(byte('0' + rand.Intn(10)))
		}
	}

	return code.String()
}
