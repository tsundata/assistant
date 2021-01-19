package utils

import (
	"math/rand"
	"net"
	"regexp"
	"strings"
)

const (
	UrlRegex = `https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`
)

func IsUrl(text string) bool {
	re := regexp.MustCompile(UrlRegex)
	return re.MatchString(text)
}

func IsIPv4(host string) bool {
	return net.ParseIP(host) != nil
}

// GeneratePassword containChars : "lowercase|uppercase|numbers|hyphen|underline|space|specials|brackets|no_similar"
func GeneratePassword(length int, containChars string) string {
	asciiLowercase := "abcdefghijklmnopqrstuvwxyz"
	asciiUppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits := "0123456789"
	punctuation := "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"

	charsStr := ""
	if strings.Contains(containChars, "lowercase") {
		charsStr += asciiLowercase
	}
	if strings.Contains(containChars, "uppercase") {
		charsStr += asciiUppercase
	}
	if strings.Contains(containChars, "numbers") {
		charsStr += digits
	}
	if strings.Contains(containChars, "hyphen") {
		charsStr += "-"
	}
	if strings.Contains(containChars, "underline") {
		charsStr += "_"
	}
	if strings.Contains(containChars, "space") {
		charsStr += " "
	}
	if strings.Contains(containChars, "specials") {
		charsStr += punctuation
	}
	if strings.Contains(containChars, "brackets") {
		charsStr += "{}[]()<>"
	}
	if strings.Contains(containChars, "no_similar") {
		noSimilarChars := []byte("0ODQ1lLj8B5S2Z")
		for _, c := range noSimilarChars {
			charsStr = strings.Replace(charsStr, string(c), "", 1)
		}
	}
	if charsStr == "" {
		return ""
	}

	charsStrSlice := []byte(charsStr)
	charsStrLength := len(charsStrSlice)
	var password strings.Builder
	for i := 0; i < length; i++ {
		randNumber := rand.Intn(charsStrLength)
		password.WriteByte(charsStrSlice[randNumber])
	}
	return password.String()
}
