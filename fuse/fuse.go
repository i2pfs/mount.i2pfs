package fuse

import (
	"fmt"
	"syscall"

	"github.com/hanwen/go-fuse/fuse"
	"github.com/i2pfs/mount.i2pfs/fuse/nodefs"
	"github.com/i2pfs/mount.i2pfs/fuse/pathfs"
	"github.com/i2pfs/mount.i2pfs/log"
)

func NewServer() *fuse.Server {
	fs := pathfs.NewFs()
	pathFs := pathfs.NewPathNodeFs(fs)
	conn := nodefs.NewFileSystemConnector(pathFs)
	mountOptions := fuse.MountOptions{
		MaxWrite: fuse.MAX_KERNEL_WRITE,
		Options: []string{
			fmt.Sprintf("max_read=%d", fuse.MAX_KERNEL_WRITE),
			"fsname=i2pfs",
		},
	}
	fuseServer, err := fuse.NewServer(conn.RawFS(), "/tmp/test", &mountOptions)
	if err != nil {
		log.Panic(err)
	}
	syscall.Umask(0000)
	return fuseServer
}
