package i2pClientHandler

import (
	"github.com/i2pfs/i2pfsd/i2p"
	"github.com/xaionaro-go/errors"
)

func MessageHandler(conn i2p.Connection, buf []byte) error {
	return errors.NotImplemented.New(nil, "i2pClientHandler.MessageHandler")
}
