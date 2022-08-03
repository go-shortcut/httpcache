//go:build !appengine
// +build !appengine

package memcache

import (
	"github.com/go-shortcut/httpcache/internal/testsuit"
	"net"
	"testing"
)

const testServer = "localhost:11211"

func TestMemCache(t *testing.T) {
	conn, err := net.Dial("tcp", testServer)
	if err != nil {
		// TODO: rather than skip the test, fall back to a faked memcached server
		t.Skipf("skipping test; no server running at %s", testServer)
	}
	conn.Write([]byte("flush_all\r\n")) // flush memcache
	conn.Close()

	testsuit.Cache(t, New(testServer))
}
