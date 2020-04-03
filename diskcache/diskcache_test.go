package diskcache

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/micaelAlastor/httpcache/test"
)

func TestDiskCache(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "httpcache")
	if err != nil {
		t.Fatalf("TempDir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	test.Cache(t, New(tempDir))
}
