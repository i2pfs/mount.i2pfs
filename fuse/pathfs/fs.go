package pathfs

import (
	"github.com/hanwen/go-fuse/fuse/pathfs"
)

type fs struct {
	pathfs.FileSystem
}

func NewFs() *fs {
	return &fs{
		FileSystem: pathfs.NewDefaultFileSystem(),
	}
}
