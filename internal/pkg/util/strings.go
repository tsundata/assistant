package util

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/google/uuid"
	"math"
	"math/big"
	"net"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

const (
	UrlRegex   = `https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`
	EmailRegex = `[a-z0-9_\.-]+\@[\da-z\.-]+\.[a-z\.]{2,6}`
)

func IsUrl(text string) bool {
	re := regexp.MustCompile("^" + UrlRegex + "$")
	return re.MatchString(text)
}

func IsIPv4(host string) bool {
	return net.ParseIP(host) != nil
}

// RandString containChars : "lowercase|uppercase|numbers|hyphen|underline|space|specials|brackets|no_similar"
func RandString(length int, containChars string) string {
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
		randNumber, err := rand.Int(rand.Reader, big.NewInt(int64(charsStrLength)))
		if err != nil {
			return ""
		}

		password.WriteByte(charsStrSlice[randNumber.Int64()])
	}
	return password.String()
}

// UUID generates a random ID for a message
func UUID() string {
	u := uuid.New()
	return u.String()
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

//ToBase64 base64 encoding
func ToBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

//ImageToBase64 image base64 encoding
func ImageToBase64(b []byte) string {
	mimeType := http.DetectContentType(b)
	base64Encoding := strings.Builder{}
	base64Encoding.WriteString("data:")
	switch mimeType {
	case "image/jpeg":
		base64Encoding.WriteString("image/jpeg")
	case "image/png":
		base64Encoding.WriteString("image/png")
	default:
		return ""
	}
	base64Encoding.WriteString(";base64,")
	base64Encoding.WriteString(ToBase64(b))

	return base64Encoding.String()
}

func FirstToUpper(str string) string {
	if len(str) < 1 {
		return ""
	}
	strA := []rune(str)
	if strA[0] >= 97 && strA[0] <= 122 {
		strA[0] -= 32
	}
	return string(strA)
}

func TypeName(s interface{}) string {
	return reflect.TypeOf(s).String()
}

func ModelName(s interface{}) string {
	t := TypeName(s)
	t = strings.ReplaceAll(t, "*", "")
	t = strings.ReplaceAll(t, "[]", "")
	t = strings.ReplaceAll(t, "pb.", "")
	return strings.ToLower(t)
}

func SubString(s string, start, end int) string {
	if start > end {
		return ""
	}
	r := []rune(s)
	l := len(r)
	if end >= l {
		end = l
	}
	return string(r[start:end])
}

const DefaultAllocate int32 = 256

func ParseInt32(desired string) int32 {
	parsed, err := strconv.Atoi(desired)
	if err != nil {
		return DefaultAllocate
	}
	// check for lower and upper bounds
	if parsed > 0 && parsed <= math.MaxInt32 {
		return int32(parsed)
	}
	return DefaultAllocate
}
