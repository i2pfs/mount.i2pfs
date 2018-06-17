// +build linux freebsd

package fuse

import (
	"os"
	"os/exec"
)

func (server *Server) KillUsers() error {
	cmd := exec.Command("fuser", "-k", server.mountpoint)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (server *Server) LazyUnmount() error {
	cmd := exec.Command("fusermount", "-u", "-z", server.mountpoint)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
