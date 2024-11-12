package info

import (
	"encoding/json"
	"testing"

	"github.com/streamingfast/substreams/manifest"
	"github.com/stretchr/testify/require"
)

func TestBasicInfo(t *testing.T) {
	reader, err := manifest.NewReader("https://github.com/streamingfast/substreams-uniswap-v3/releases/download/v0.2.8/substreams.spkg")
	require.NoError(t, err)

	pkgBundle, err := reader.Read()
	require.NoError(t, err)
	require.NotNil(t, pkgBundle)

	info, err := Basic(pkgBundle.Package, pkgBundle.Graph)
	require.NoError(t, err)

	r, err := json.MarshalIndent(info, "", "  ")
	require.NoError(t, err)
	require.NotNil(t, r)
}

func TestExtendedInfo(t *testing.T) {
	info, err := Extended("https://github.com/streamingfast/substreams-uniswap-v3/releases/download/v0.2.8/substreams.spkg", "graph_out")
	require.NoError(t, err)

	r, err := json.MarshalIndent(info, "", "  ")
	require.NoError(t, err)
	require.NotNil(t, r)
}
