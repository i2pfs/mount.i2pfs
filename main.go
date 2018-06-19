package main

import (
	"os"
	"syscall"

	"github.com/i2pfs/mount.i2pfs/config"
	"github.com/pborman/getopt/v2"
	"github.com/sevlyar/go-daemon"
	"github.com/xaionaro-go/log"
)

func usage() {
	getopt.Usage()
	os.Exit(int(syscall.EINVAL))
}

func main() {
	cfg := config.GetDefaults()
	cfg.ReloadConfig()

	getopt.SetParameters("mountpoint")
	helpFlag := getopt.BoolLong("help", 'h', "display help")
	foregroundFlag := getopt.BoolLong("foreground", 'f', "run in foreground")
	getopt.Parse()

	if *helpFlag {
		usage()
		return
	}

	args := getopt.Args()
	if len(args) == 0 {
		usage()
		return
	}

	mountpoint := args[0]

	if *foregroundFlag {
		err := doMount(cfg.SamUrl, cfg.PeersFilePath, mountpoint)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	ctx := &daemon.Context{}

	d, err := ctx.Reborn()
	if err != nil {
		log.Fatal("Unable to daemonize: ", err)
	}
	if d != nil {
		return
	}
	defer ctx.Release()

	log.Debug("Daemonized")
	err = doMount(cfg.SamUrl, cfg.PeersFilePath, mountpoint)
	if err != nil {
		log.Fatal(err)
	}
}
