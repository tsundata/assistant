package component

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/tsundata/assistant/mock"
	"testing"
)

func TestMockComponent(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	message := mock.NewMockMessageSvcClient(ctl)
	middle := mock.NewMockMiddleSvcClient(ctl)
	comp := MockComponent(message, middle)
	assert.NotNil(t, comp.Message())
	assert.NotNil(t, comp.Middle())
	assert.Nil(t, comp.User())
}
