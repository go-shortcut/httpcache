package test_test

import (
	"../memorycache"
	"../test"
	"testing"
)

func TestMemoryCache(t *testing.T) {
	test.Cache(t, memorycache.NewMemoryCache())
}
