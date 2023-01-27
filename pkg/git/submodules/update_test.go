package submodules

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/go-git/go-git/v5/config"

	"github.com/stretchr/testify/assert"
)

func TestSubmodules_updateConfigFile(t *testing.T) {
	dir := t.TempDir()

	data, err := os.ReadFile(filepath.Join("testdata", "multiple_submodules", ".gitmodules"))
	assert.NoError(t, err)

	err = os.WriteFile(filepath.Join(dir, ".gitmodules"), data, 0o600)
	assert.NoError(t, err)

	_, err = updateConfigFile(dir, "foo", "bar", "v1.2.4")
	assert.NoError(t, err)

	data, err = os.ReadFile(filepath.Join(dir, ".gitmodules"))
	assert.NoError(t, err)

	cfg := config.NewModules()
	err = cfg.Unmarshal(data)
	assert.NoError(t, err)

	assert.Equal(t, "v1.2.4", cfg.Submodules["images/cheese/mount/bar"].Branch)
	assert.Equal(t, "v1.2.4", cfg.Submodules["images/wine/mount/bar"].Branch)
}
