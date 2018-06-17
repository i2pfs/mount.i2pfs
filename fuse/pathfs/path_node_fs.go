package pathfs

import (
	"github.com/hanwen/go-fuse/fuse/pathfs"
)

type PathNodeFs struct {
	*pathfs.PathNodeFs
}

func NewPathNodeFs(fs pathfs.FileSystem) *PathNodeFs {
	return &PathNodeFs{
		PathNodeFs: pathfs.NewPathNodeFs(fs, nil),
	}
}
