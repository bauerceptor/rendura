package pisnap_test

import (
	"testing"

	"github.com/bauerceptor/rendura/rendura_snap"

	"github.com/stretchr/testify/assert"
)

func TestCaptureOrErr(t *testing.T) {
	file, err := pisnap.CaptureOrErr()
	assert.Error(t, err)
	assert.Empty(t, file)
}
