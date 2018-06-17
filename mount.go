package main

import (
	"github.com/i2pfs/mount.i2pfs/consts"
	"github.com/i2pfs/mount.i2pfs/fuse"
	"github.com/i2pfs/mount.i2pfs/signals"
)

func doMount(samUrl, peersFilePath string) error {
	err := setRLimitNoFile(consts.RLIMIT_NOFILE)
	if err != nil {
		return err
	}
	fuseServer := fuse.NewServer()
	signals.Init(fuseServer)
	fuseServer.Serve()
	return nil
}