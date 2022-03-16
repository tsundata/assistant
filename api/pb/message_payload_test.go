package pb

import (
	"github.com/stretchr/testify/assert"
	"github.com/tsundata/assistant/api/enum"
	"testing"
)

func TestPayload(t *testing.T) {
	var payload MsgPayload
	payload = &TextMsg{Text: "test"}
	assert.Equal(t, enum.MessageTypeText, payload.Type())

	payload = &ImageMsg{Src: "/test"}
	assert.Equal(t, enum.MessageTypeImage, payload.Type())
}
