package interpreter

import "log"

var Debug bool

func debugLog(out string) {
	if Debug {
		log.Println(out)
	}
}
