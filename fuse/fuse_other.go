// +build !linux,!freebsd

package fuse

import (
	"github.com/xaionaro-go/errors"
)

func (server *Server) KillUsers() error {
	return errors.ErrNotImplemented.SetArgs("server.LazyUnmount()")
}

func (server *Server) LazyUnmount() error {
	cmd := exec.Command("fusermount", "-u", "-z", mountpoint)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return errors.ErrNotImplemented.SetArgs("server.LazyUnmount()")
}


