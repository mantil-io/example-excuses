package excuses

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

// unit test mocks outside resource
func TestUnit(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `It works, but it's not been tested.
It works for me.
I thought I fixed that.
`)
	}))
	defer svr.Close()

	run(t, svr.URL, 3)
}

// integration test uses resource outside of development evironment
func TestIntegration(t *testing.T) {
	url := "https://gist.githubusercontent.com/ianic/f3335ba0b7ec63cbb821f8a7b735d86e/raw/066e44b04682295781164c538774db645dfe4cc6/excuses.txt"
	run(t, url, 63)
}

func run(t *testing.T, url string, expectedCount int) {
	log.SetOutput(io.Discard) // silence logs in test

	t.Run("load", func(t *testing.T) {
		ex := New()
		cr, err := ex.Load(nil, LoadRequest{URL: url})
		require.NoError(t, err)
		require.Equal(t, expectedCount, cr.Count)

		cr, err = ex.Load(nil, LoadRequest{URL: url})
		require.NoError(t, err)
		require.Equal(t, expectedCount, cr.Count)
	})

	t.Run("preload", func(t *testing.T) {
		t.Setenv(preloadURLEnv, url)

		ex := New()
		cr, err := ex.Count(nil)
		require.NoError(t, err)
		require.Equal(t, expectedCount, cr.Count)
	})

	t.Run("clear", func(t *testing.T) {
		ex := New()
		cr, err := ex.Load(nil, LoadRequest{URL: url})
		require.NoError(t, err)
		require.Equal(t, expectedCount, cr.Count)

		ex.Clear(nil)
		cr, err = ex.Count(nil)
		require.NoError(t, err)
		require.Equal(t, 0, cr.Count)
	})

	t.Run("random", func(t *testing.T) {
		ex := New()
		cr, err := ex.Load(nil, LoadRequest{URL: url})
		require.NoError(t, err)
		require.Equal(t, expectedCount, cr.Count)

		seen := make(map[string]int)

		for i := 0; i < expectedCount*10; i++ {
			rr, err := ex.Random(nil)
			require.NoError(t, err)
			require.True(t, len(rr.Excuse) > 0)
			seen[rr.Excuse] = seen[rr.Excuse] + 1
		}

		// show random distirbution
		for k, v := range seen {
			t.Logf("%2d %s", v, k)
		}
	})

}
