package client

import (
	"github.com/i2pfs/i2pfsd/i2p"
	pb "github.com/i2pfs/i2pfsd/protobuf/generated"
	handler "github.com/i2pfs/mount.i2pfs/i2pClient/handler"
	"github.com/majestrate/i2p-tools/sam3"
	"github.com/xaionaro-go/errors"
	"github.com/xaionaro-go/log"
)

type peer struct {
	address       sam3.I2PAddr
	streamSession *sam3.StreamSession
	conn          i2p.Connection
}
type peers map[string]*peer

func (peers peers) OnlineCount() (count int) {
	for _, peer := range peers {
		if peer == nil {
			continue
		}
		if peer.conn != nil {
			count++
		}
	}
	return
}

func helloMessageHandler(conn i2p.Connection, buf []byte) error {
	hello := pb.Hello{}
	err := hello.Unmarshal(buf)
	if err != nil {
		return err
	}
	log.Debugln("hello connection: received hello: ", hello)
	switch hello.ConnectionType {
	case pb.ConnectionType_Client:
		conn.SetMessageHandler(handler.MessageHandler)
	default:
		return errors.ProtocolMismatch
	}
	return err
}

func (peer *peer) CreateStreamSession(cluster *cluster) error {
	var err error
	peer.streamSession, err = cluster.sam.NewStreamSession("mount.i2pfs." + peer.address.Base32(), cluster.keys, sam3.Options_Medium)
	if err != nil {
		peer.streamSession = nil
		return errors.UnableToConnect.SetArgs("NewStreamSession", err.Error())
	}
	return nil
}

func (peer *peer) Connect(cluster *cluster) error {
	if peer.conn != nil {
		return nil
	}
	if peer.streamSession == nil {
		err := peer.CreateStreamSession(cluster)
		if err != nil {
			return err
		}
	}

	conn, err := peer.streamSession.DialI2P(peer.address)
	if err != nil {
		peer.conn = nil
		return errors.UnableToConnect.SetArgs("DialI2P", err.Error())
	}

	peer.conn = i2p.NewConnection(conn)
	peer.conn.SendMessage(pb.Hello{ConnectionType: pb.ConnectionType_Client})
	peer.conn.SetMessageHandler(helloMessageHandler)

	go peer.conn.Loop()
	return nil
}
