// +build darwin dragonfly freebsd linux nacl netbsd openbsd solaris

package fileidentity

import (
	"os"
	"testing"
)

func Test_FileIdentity_Equal(t *testing.T) {
	src, _ := os.Open("testdata/test_file.txt")
	samefile, _ := os.Open("testdata/hardlink.txt")
	id1, _ := NewFileIdentity(src)
	id2, _ := NewFileIdentity(samefile)

	if !id1.Equals(id2) {
		t.Error("hard link files should be equal")
	}
}

func Test_FileIdentity_NotEqual(t *testing.T) {
	src, _ := os.Open("testdata/test_file.txt")
	samefile, _ := os.Open("testdata/different_file.txt")
	id1, _ := NewFileIdentity(src)
	id2, _ := NewFileIdentity(samefile)

	if id1.Equals(id2) {
		t.Error("not hard link files should not be equal")
	}
}