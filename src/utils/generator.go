package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/twinj/uuid"
)

//*Return a string uuid
func GenerateUuid() string {
	uuid := uuid.NewV4()

	return uuid.String()
}

func GenerateSKU(productName string) (firstName string) {
	if strings.Contains(productName, " ") {
		firstName = strings.Split(productName, " ")[0]
	} else {
		firstName = productName
	}

	firstName = firstName + "-" + fmt.Sprint(time.Now().UnixNano()/int64(time.Millisecond))
	return
}
