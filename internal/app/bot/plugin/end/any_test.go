package end

import (
	"context"
	"testing"
)

func TestAny(t *testing.T) {
	e := &End{}
	_, err := e.Run(context.TODO(), nil)
	if err != nil {
		t.Fatal(err)
	}
}