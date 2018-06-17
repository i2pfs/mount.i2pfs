package signals

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/hanwen/go-fuse/fuse"
	"github.com/i2pfs/mount.i2pfs/log"
)

func Init(fuseServer *fuse.Server) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, syscall.SIGTERM)
	go func() {
		<-ch
		err := fuseServer.Unmount()
		if err != nil {
			log.Warningf("[signals] Got an error while fuseServer.Unmount(): %v", err)
		}
		os.Exit(0)
	}()
	return
}
