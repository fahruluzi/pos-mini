package utils

import (
	"crypto/rand"
	"io"
	"strings"

	"github.com/twinj/uuid"
)

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

//*Return a random number based on inputed max param
func RandomNumber(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

//*Return a string uuid
func GenerateUuid() string {
	uuid := uuid.NewV4()

	return uuid.String()
}

//*Return a unique mixed string number with separator underscore
func GenerateUsername(name string) string {
	splitName:=strings.Split(name," ")

	getRandNumber:=RandomNumber(4)

	return strings.Join([]string{splitName[0],getRandNumber},"_")
}

