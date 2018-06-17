package main

import (
	"github.com/i2pfs/mount.i2pfs/config"
	"github.com/i2pfs/mount.i2pfs/log"
)

func main() {
	cfg := config.GetDefaults()
	cfg.ReloadConfig()

	err := doMount(cfg.SamUrl, cfg.PeersFilePath)
	if err != nil {
		log.Panic(err)
	}
}
