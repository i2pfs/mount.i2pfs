package pathfs

import (
	"github.com/i2pfs/mount.i2pfs/i2pClient"
	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/pathfs"
	"github.com/xaionaro-go/errors"
	"github.com/xaionaro-go/log"
)

type fs struct {
	pathfs.FileSystem
	cluster client.Cluster
}

func NewFs(cluster client.Cluster) *fs {
	return &fs{
		FileSystem: pathfs.NewDefaultFileSystem(),
		cluster: cluster,
	}
}

func (fs *fs) OpenDir(name string, context *fuse.Context) (stream []fuse.DirEntry, status fuse.Status) {
	log.Debugf(`fs.OpenDir("%v", context)`, name)
	dirEntries, err := fs.cluster.OpenDir(name)
	if err != nil {
		log.Warning(errors.CannotOpenDir.New(nil, err, name))
		return nil, fuse.EAGAIN
	}
	for _, dirEntry := range dirEntries.Slice() {
		stream = append(stream, fuse.DirEntry{
			Mode: dirEntry.GetMode(),
			Name: dirEntry.GetName(),
			Ino: dirEntry.GetIno(),
		})
	}
	return stream, fuse.OK
}
