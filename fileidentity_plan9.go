// +build plan9

package fileidentity

import (
	"errors"
	"os"
	"syscall"
)

type FileIdentity struct {
	Path uint64
	Type uint16
	Dev uint32
}

func (f *FileIdentity) Equals(o *FileIdentity) bool {
	return f.Path == o.Path && f.Type == o.Type && f.Dev == o.Dev
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
	stat := (*info).Sys().(*syscall.Dir)
	if stat == nil {
		return nil, errors.New("fileidentity.NewFileIdentity: can't get internal data")
	}
	return &FileIdentity{
		Path: stat.Qid.Path,
		Type: stat.Type,
		Dev: stat.Dev,
	}, nil
}
