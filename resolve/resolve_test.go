package resolve

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)


func TestResolveGet(t *testing.T) {
	modules, err := ResolveGet(context.Background(), KnownImports, "golang.org/x/mod/...")
	require.NoError(t, err)

	assert.Len(t, modules, 3)



}