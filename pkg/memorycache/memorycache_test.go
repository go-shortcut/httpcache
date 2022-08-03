package memorycache_test

import (
	"github.com/go-shortcut/httpcache/internal/testsuit"
	"github.com/go-shortcut/httpcache/pkg/memorycache"
	"testing"
)

func TestMemoryCache(t *testing.T) {
	testsuit.Cache(t, memorycache.NewMemoryCache())
}
