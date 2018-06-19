package main

import (
	"github.com/i2pfs/mount.i2pfs/consts"
	"github.com/i2pfs/mount.i2pfs/fuse"
	"github.com/i2pfs/mount.i2pfs/signals"
	"github.com/i2pfs/mount.i2pfs/i2pClient"
	"github.com/xaionaro-go/errors"
	"github.com/xaionaro-go/log"
)

func doMount(samUrl, peersFilePath, mountpoint string) error {
	err := setRLimitNoFile(consts.RLIMIT_NOFILE)
	if err != nil {
		return log.WarningWrapper(errors.CannotSetRLimitNoFile, err)
	}
	cluster, err := client.NewCluster(samUrl, peersFilePath)
	if err != nil {
		return log.WarningWrapper(errors.UnableToConnect, err)
	}
	fuseServer := fuse.NewServer(cluster, mountpoint)
	signals.Init(fuseServer)
	fuseServer.Serve()
	return nil
}
