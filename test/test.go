package test

import (
	"../httpcache"
	"../memorycache"
	"bytes"
	"testing"
)

// Cache excercises a httpcache.Cache implementation.
func Cache(t *testing.T, cache httpcache.Cache) {
	key := "testKey"
	_, ok := cache.Get(key)
	if ok {
		t.Fatal("retrieved key before adding it")
	}

	val := []byte("some bytes")
	cache.Set(key, val)

	retVal, ok := cache.Get(key)
	if !ok {
		t.Fatal("could not retrieve an element we just added")
	}
	if !bytes.Equal(retVal, val) {
		t.Fatal("retrieved a different value than what we put in")
	}

	cache.Delete(key)

	_, ok = cache.Get(key)
	if ok {
		t.Fatal("deleted key still present")
	}
}

// NewMemoryCacheTransport returns a new Transport using the in-memory cache implementation
func NewMemoryCacheTransport() *httpcache.Transport {
	c := memorycache.NewMemoryCache()
	t := httpcache.NewTransport(c)
	return t
}
