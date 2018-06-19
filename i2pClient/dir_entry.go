package client

import (
	"github.com/xaionaro-go/errors"
)

type DirEntry interface {
	GetMode() uint32
	GetName() string
	GetIno() uint64
}

type dirEntry struct {
}

func NewDirEntry() DirEntry {
	return &dirEntry{}
}

type DirEntries interface {
	Slice() []DirEntry
}

type dirEntries []DirEntry

func (dirEntries dirEntries) Slice() []DirEntry {
	return []DirEntry(dirEntries)
}

func (dirEntry dirEntry) GetMode() uint32 {
	panic(errors.NotImplemented)
	return 0
}
func (dirEntry dirEntry) GetName() string {
	panic(errors.NotImplemented)
	return ""
}
func (dirEntry dirEntry) GetIno() uint64 {
	panic(errors.NotImplemented)
	return 0
}
