package diskcache

import (
	"github.com/go-shortcut/httpcache/internal/testsuit"
	"io/ioutil"
	"os"
	"testing"
)

func TestDiskCache(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "httpcache")
	if err != nil {
		t.Fatalf("TempDir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	testsuit.Cache(t, New(tempDir))
}
