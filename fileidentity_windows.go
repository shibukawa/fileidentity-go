// +build windows

package fileidentity

import (
	"os"
	"errors"
	"syscall"
)

type FileIdentity struct {
	VolumeSerialNumber int
	FileIndex uint64
}

func (f *FileIdentity) Equals(o *FileIdentity) bool {
	return f.FileIndex == o.FileIndex && f.Dev == o.Dev
}

func NewFileIdentity(file *os.File) (*FileIdentity, error) {
	if file == nil {
		return nil, errors.New("fileidentity.NewFileIdentityFromFile: input file is nil")
	}
	var d syscall.ByHandleFileInformation
	e := syscall.GetFileInformationByHandle(syscall.Handle(file.Fd()), &d)
	if e != nil {
		return nil, errors.New("fileidentity.NewFileIdentityFromFile: can't get file information from " + file.Name())
	}
	return *NewFileIdentity{
		VolumeSerialNumber: d.VolumeSerialNumber,
		FileIndex: uint64(d.FileIndexHigh) << 32 + uint64(d.FileIndexLow),
	}, nil
}
