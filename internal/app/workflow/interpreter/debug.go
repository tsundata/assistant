package interpreter

import (
	"encoding/json"
	"fmt"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"log"
	"strings"
)

var Debug bool

func debugLog(out string) {
	if Debug {
		log.Println(out)
	}
}

func nodeLog(name string, regular string, parameters map[string]interface{}, secret string, input, output interface{}, err error) string {
	var result strings.Builder
	result.WriteString(fmt.Sprintf("%s (%s)\n", name, regular))
	result.WriteString("====================================\n")
	p, _ := json.Marshal(parameters)
	result.WriteString(fmt.Sprintf("parameters: %v\n", utils.ByteToString(p)))
	result.WriteString(fmt.Sprintf("secret: %v\n", secret))
	i, _ := json.Marshal(input)
	o, _ := json.Marshal(output)
	result.WriteString(fmt.Sprintf("%s ---> %s\n", utils.ByteToString(i), utils.ByteToString(o)))
	result.WriteString(fmt.Sprintf("error: %v\n", err))
	result.WriteString("------------------------------------\n")

	return result.String()
}
