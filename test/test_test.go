package test_test

import (
	"testing"

	"github.com/micaelAlastor/httpcache"
	"github.com/micaelAlastor/httpcache/test"
)

func TestMemoryCache(t *testing.T) {
	test.Cache(t, httpcache.NewMemoryCache())
}
