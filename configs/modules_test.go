package configs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootModulesDirs(t *testing.T) {
	dirs, err := RootModulesDirs("fixtures")

	assert.NoError(t, err)
	assert.Equal(t, []string{"fixtures/root", "fixtures/root/nested"}, dirs)
}
