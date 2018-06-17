package signals

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/i2pfs/mount.i2pfs/fuse"
	"github.com/xaionaro-go/log"
)

func Init(fuseServer *fuse.Server) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, syscall.SIGTERM)
	go func() {
		<-ch
		defer os.Exit(0)
		err := fuseServer.Unmount()
		if err == nil {
			return
		}
		log.Warningf("[signals] Got an error while fuseServer.Unmount(): %v. Killing all users of the mountpoint.", err)
		err = fuseServer.KillUsers()
		if err != nil {
			log.Warningf("[signals] Got an error while fuseServer.KillUsers(): %v", err)
		} else {
			err = fuseServer.Unmount()
			if err == nil {
				return
			}
			log.Warningf("[signals] Got an error while fuseServer.Unmount(): %v", err)
		}
		log.Errorf("[signals] Give up. Lazy unmount.")
		err = fuseServer.LazyUnmount()
		if err != nil {
			log.Errorf("[signals] Got an error while fuseServer.LazyUnmount(): %v", err)
		}
	}()
	return
}
