package discovery

import (
	"fmt"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/util"
	"reflect"
)

func SvcAddr(conf *config.AppConfig, service string) string {
	t := reflect.ValueOf(conf.SvcAddr)
	field := util.FirstToUpper(service)
	v := t.FieldByName(field)

	res := v.String()
	if res == "<invalid Value>" {
		return fmt.Sprintf("%s-error-svc-addr", service)
	}

	return v.String()
}
