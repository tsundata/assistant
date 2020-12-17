package utils

import (
	"github.com/tsundata/assistant/internal/pkg/model"
	"math/rand"
	"net"
	"regexp"
	"strings"
)

func IsUrl(text string) bool {
	re := regexp.MustCompile(`^https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`)
	return re.MatchString(text)
}

func IsIPv4(host string) bool {
	return net.ParseIP(host) != nil
}

func IsMessageOfScript(text string) bool {
	lines := strings.Split(text, "\n")
	if len(lines) >= 1 {
		re := regexp.MustCompile(`^#!script:\w+$`)
		return re.MatchString(strings.TrimSpace(lines[0]))
	}
	return false
}

func IsMessageOfAction(text string) bool {
	lines := strings.Split(text, "\n")
	if len(lines) >= 1 {
		re := regexp.MustCompile(`^#!action$`)
		return re.MatchString(strings.TrimSpace(lines[0]))
	}
	return false
}

func MessageScriptKind(text string) string {
	if !IsMessageOfScript(text) {
		return model.MessageScriptOfUndefined
	}

	lines := strings.Split(text, "\n")
	if len(lines) >= 1 {
		return strings.Replace(strings.TrimSpace(lines[0]), "#!script:", "", -1)
	}
	return model.MessageScriptOfUndefined
}

func SliceDiff(s1, s2 []string) []string {
	mb := make(map[string]struct{}, len(s2))
	for _, x := range s2 {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range s1 {
		if _, ok := mb[x]; !ok {
			diff = append(diff, x)
		}
	}
	return diff
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
