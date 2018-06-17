package nodefs

import (
	"github.com/hanwen/go-fuse/fuse/nodefs"
	my_pathfs "github.com/i2pfs/mount.i2pfs/fuse/pathfs"
)

func NewFileSystemConnector(pathFs *my_pathfs.PathNodeFs) *nodefs.FileSystemConnector {
	return nodefs.NewFileSystemConnector(pathFs.Root(), nil)
}
