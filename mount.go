package main

import (
	"github.com/i2pfs/mount.i2pfs/consts"
	"github.com/i2pfs/mount.i2pfs/fuse"
	"github.com/i2pfs/mount.i2pfs/signals"
	"github.com/i2pfs/mount.i2pfs/i2pClient"
)

func doMount(samUrl, peersFilePath, mountpoint string) error {
	err := setRLimitNoFile(consts.RLIMIT_NOFILE)
	if err != nil {
		return err
	}
	cluster, err := client.NewCluster(samUrl, peersFilePath)
	if err != nil {
		return err
	}
	fuseServer := fuse.NewServer(cluster, mountpoint)
	signals.Init(fuseServer)
	fuseServer.Serve()
	return nil
}
