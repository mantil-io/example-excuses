package excuses

import (
	"io"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	testURL          = "https://gist.githubusercontent.com/ianic/f3335ba0b7ec63cbb821f8a7b735d86e/raw/066e44b04682295781164c538774db645dfe4cc6/excuses.txt"
	excusesInTestURL = 63
)

func TestLoadClear(t *testing.T) {
	log.SetOutput(io.Discard) // silence logs in test
	ex := New()

	cr, err := ex.Count(nil)
	require.NoError(t, err)
	require.Equal(t, 0, cr.Count)

	cr, err = ex.Load(nil, LoadRequest{URL: testURL})
	require.NoError(t, err)
	require.Equal(t, excusesInTestURL, cr.Count)

	// repeat load to test that we are not adding duplicates
	cr, err = ex.Load(nil, LoadRequest{URL: testURL})
	require.NoError(t, err)
	require.Equal(t, excusesInTestURL, cr.Count)

	ex.Clear(nil)
	cr, err = ex.Count(nil)
	require.NoError(t, err)
	require.Equal(t, 0, cr.Count)
}

func TestRandom(t *testing.T) {
	ex := New()
	cr, err := ex.Load(nil, LoadRequest{URL: testURL})
	require.NoError(t, err)
	require.Equal(t, excusesInTestURL, cr.Count)

	seen := make(map[string]int)

	for i := 0; i < excusesInTestURL*10; i++ {
		rr, err := ex.Random(nil)
		require.NoError(t, err)
		require.True(t, len(rr.Excuse) > 0)
		seen[rr.Excuse] = seen[rr.Excuse] + 1
	}

	// show random distirbution
	for k, v := range seen {
		t.Logf("%2d %s", v, k)
	}
}

func TestPreload(t *testing.T) {
	t.Setenv(preloadURLEnv, testURL)

	ex := New()
	cr, err := ex.Count(nil)
	require.NoError(t, err)
	require.Equal(t, excusesInTestURL, cr.Count)
}
