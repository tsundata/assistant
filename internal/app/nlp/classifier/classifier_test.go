package classifier

import (
	"github.com/stretchr/testify/require"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"testing"
)

func TestClassifier(t *testing.T) {
	conf, err := config.CreateAppConfig(app.NLP)
	if err != nil {
		t.Fatal(err)
	}

	c := NewClassifier(conf)
	err = c.LoadRule()
	if err != nil {
		t.Fatal(err)
	}
	require.Len(t, c.rules, 2)
}
