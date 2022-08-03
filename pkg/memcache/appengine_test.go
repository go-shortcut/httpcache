//go:build appengine
// +build appengine

package memcache

import (
	"github.com/go-shortcut/httpcache/internal/testsuit"
	"testing"

	"appengine/aetest"
)

func TestAppEngine(t *testing.T) {
	ctx, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer ctx.Close()

	testsuit.Cache(t, New(ctx))
}
