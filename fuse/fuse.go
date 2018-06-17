package fuse

import (
	"fmt"
	"path/filepath"
	"syscall"

	"github.com/hanwen/go-fuse/fuse"
	"github.com/i2pfs/mount.i2pfs/fuse/nodefs"
	"github.com/i2pfs/mount.i2pfs/fuse/pathfs"
	"github.com/xaionaro-go/log"
)

type Server struct {
	*fuse.Server
	mountpoint string
}

func NewServer(mountpoint string) *Server {
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
	server := &Server{}
	var err error
	server.mountpoint, err = filepath.Abs(mountpoint)
	if err != nil {
		log.Panic(err)
	}
	server.Server, err = fuse.NewServer(conn.RawFS(), server.mountpoint, &mountOptions)
	if err != nil {
		log.Panic(err)
	}
	syscall.Umask(0000)
	return server
}
