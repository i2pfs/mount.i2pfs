// +build !linux,!freebsd

package fuse

import (
	"github.com/xaionaro-go/errors"
)

func (server *Server) KillUsers() error {
	return errors.NotImplemented.SetArgs("server.KillUsers()")
}

func (server *Server) LazyUnmount() error {
	return errors.NotImplemented.SetArgs("server.LazyUnmount()")
}
