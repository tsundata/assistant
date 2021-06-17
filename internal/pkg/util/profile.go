package util

import (
	"fmt"
	"time"
)

func Duration(invocation time.Time, name string) {
	elapsed := time.Since(invocation)
	fmt.Printf("%s lasted %s", name, elapsed)
}
