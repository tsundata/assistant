package utils

import (
	cRand "crypto/rand"
	"fmt"
	"io"
	"math/rand"
	"net"
	"regexp"
	"strings"
)

const (
	UrlRegex = `https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`
)

func IsUrl(text string) bool {
	re := regexp.MustCompile("^" + UrlRegex + "$")
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
		noSimilarChars := StringToByte("0ODQ1lLj8B5S2Z")
		for _, c := range noSimilarChars {
			charsStr = strings.ReplaceAll(charsStr, string(c), "")
		}
	}
	if charsStr == "" {
		return ""
	}

	charsStrSlice := StringToByte(charsStr)
	charsStrLength := len(charsStrSlice)
	var password strings.Builder
	for i := 0; i < length; i++ {
		randNumber := rand.Intn(charsStrLength)
		password.WriteByte(charsStrSlice[randNumber])
	}
	return password.String()
}

// GenerateUUID generates a random ID for a message
func GenerateUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(cRand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

// ExtractUUID extract path
func ExtractUUID(path string) string {
	re := regexp.MustCompile(`(\w{8}-\w{4}-\w{4}-\w{4}-\w{12})`)
	return re.FindString(path)
}

// DataMasking Data Masking
func DataMasking(data string) string {
	if len(data) > 3 {
		var res strings.Builder
		res.WriteString(data[:3])
		res.WriteString("******")
		res.WriteString(data[len(data)-3:])
		return res.String()
	} else if len(data) >= 1 {
		var res strings.Builder
		res.WriteString(data[:1])
		res.WriteString("******")
		res.WriteString(data[len(data)-1:])
		return res.String()
	} else {
		return ""
	}
}
