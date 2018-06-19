package client

import (
	"github.com/xaionaro-go/errors"
	"github.com/xaionaro-go/log"
)

func (cluster *cluster) OpenDir(name string) (DirEntries, error) {
	log.Debugf(`cluster.OpenDir("%v")`, name)
	return nil, errors.NotImplemented
}
