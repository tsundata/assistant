package classifier

import (
	"github.com/stretchr/testify/require"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/model"
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

	a, err := c.Do("demo2")
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, model.CultureAttr, a)

	_, err = c.Do("demo8")
	require.ErrorIs(t, ErrEmpty, err)
}
