package client

import (
	pb "github.com/i2pfs/i2pfsd/serverToClient/protobuf/generated"
	"github.com/xaionaro-go/errors"
	"github.com/xaionaro-go/log"
)

func (cluster *cluster) OpenDir(name string) (DirEntries, error) {
	ch := make(chan *pb.Message)
	err := cluster.RequestRead(name, &pb.OpenDir{}, ch)
	if err != nil {
		return nil, log.WarningWrapper(errors.UnableToMakeRequest, err)
	}
	msg := <-ch
	log.Debugf(`answer :)… %v`, msg)
	return nil, errors.NotImplemented
}
