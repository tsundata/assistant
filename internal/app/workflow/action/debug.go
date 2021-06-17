package action

import (
	"encoding/json"
	"fmt"
	"github.com/tsundata/assistant/internal/pkg/util"
	"strings"
)

var Debug bool

func debugLog(out string) {
	if Debug {
		fmt.Println("[Action Debug]", out)
	}
}

func opcodeLog(name string, parameters []interface{}, input, output interface{}, err error) string {
	var result strings.Builder
	result.WriteString(fmt.Sprintf("%s\n", name))
	result.WriteString("====================================\n")
	p, _ := json.Marshal(parameters)
	result.WriteString(fmt.Sprintf("parameters: %v\n", util.ByteToString(p)))
	i, _ := json.Marshal(input)
	o, _ := json.Marshal(output)
	result.WriteString(fmt.Sprintf("%s ---> %s\n", util.ByteToString(i), util.ByteToString(o)))
	result.WriteString(fmt.Sprintf("error: %v\n", err))
	result.WriteString("------------------------------------\n")

	return result.String()
}
