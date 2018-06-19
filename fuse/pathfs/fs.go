package pathfs

import (
	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/pathfs"
	"github.com/xaionaro-go/log"
)

type fs struct {
	pathfs.FileSystem
}

func NewFs() *fs {
	return &fs{
		FileSystem: pathfs.NewDefaultFileSystem(),
	}
}

func (fs *fs) OpenDir(name string, context *fuse.Context) (stream []fuse.DirEntry, status fuse.Status) {
	log.Debugf(`fs.OpenDir("%v", context)`, name)
	return nil, fuse.ENOSYS
}
