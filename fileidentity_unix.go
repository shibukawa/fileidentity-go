// +build darwin dragonfly freebsd linux nacl netbsd openbsd solaris

package fileidentity

import (
	"errors"
	"os"
	"syscall"
)

type FileIdentity struct {
	Inode uint64
	Dev int32
}

func (f *FileIdentity) Equals(o *FileIdentity) bool {
	return f.Inode == o.Inode && f.Dev == o.Dev
}

func NewFileIdentity(file *os.File) (*FileIdentity, error) {
	if file == nil {
		return nil, errors.New("fileidentity.NewFileIdentityFromFile: input file is nil")
	}
	stat, err :=  file.Stat()
	if err != nil {
		return nil, errors.New("fileidentity.NewFileIdentityFromFile: can't get info")
	}
	return NewFileIdentityFromFileInfo(&stat)
}

func NewFileIdentityFromFileInfo(info *os.FileInfo) (*FileIdentity, error) {
	if info == nil {
		return nil, errors.New("fileidentity.NewFileIdentity: input file info is nil")
	}
	stat := (*info).Sys().(*syscall.Stat_t)
	if stat == nil {
		return nil, errors.New("fileidentity.NewFileIdentity: can't get internal data")
	}
	return &FileIdentity{
		Inode: stat.Ino,
		Dev: stat.Dev,
	}, nil
}