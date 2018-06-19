package client

import (
	"bufio"
	"os"

	"github.com/majestrate/i2p-tools/sam3"
	"github.com/xaionaro-go/errors"
	"github.com/xaionaro-go/log"
)

type Cluster interface {
	OpenDir(name string) (DirEntries, error)
}

type cluster struct {
	sam   *sam3.SAM
	keys  sam3.I2PKeys
	peers peers
}

func (cluster *cluster) addPeer(peerAddress string) error {
	var err error
	peer := &peer{}
	peer.address, err = cluster.sam.Lookup(peerAddress)
	if err != nil {
		return errors.CannotResolveAddress.New(err, "cluster.addPeer", peerAddress)
	}
	cluster.peers[peerAddress] = peer
	return nil
}

func (cluster *cluster) removePeer(peerAddress string) error {
	if cluster.peers[peerAddress] == nil {
		return errors.NotFound.New(nil, "removePeer", peerAddress)
	}
	cluster.peers[peerAddress] = nil
	return nil
}

func (cluster *cluster) readPeersFromFile(peersFilePath string) error {
	file, err := os.Open(peersFilePath)
	if err != nil {
		return errors.CannotOpenFile.New(err, "readPeersFromFile", peersFilePath)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		peerAddress := scanner.Text()
		err = cluster.addPeer(peerAddress)
		if err != nil {
			return err
		}
	}

	if err = scanner.Err(); err != nil {
		return err
	}

	return nil
}

func (cluster *cluster) connectToPeers() error {
	for idx, _ := range cluster.peers {
		peer := cluster.peers[idx]
		if peer == nil {
			continue
		}
		err := peer.Connect(cluster)
		if err != nil {
			smartError := err.(errors.SmartError)
			switch smartError.ToOriginal() {
			case errors.UnableToStartSession:
				return log.WarningWrapper(errors.UnableToStartSession, err)
			}
			log.Warningf("Cannot connect to %v: %v", idx, err.(errors.SmartError).ErrorShort())
			cluster.removePeer(idx)
		}
	}
	if cluster.peers.OnlineCount() == 0 {
		return errors.UnableToConnect.New(nil, "connectToPeers(): There're no reachable peers. (All peers are down? Or maybe empty peers list file?)")
	}
	log.Debugf("Online peers: %v", cluster.peers.OnlineCount())
	return nil
}

func NewCluster(samUrl, peersFilePath string) (Cluster, error) {
	var err error
	result := &cluster{peers: peers{}}

	log.Debugln("Client: NewSAM")
	result.sam, err = sam3.NewSAM(samUrl)
	if err != nil {
		return nil, err
	}

	log.Debugln("Client: readPeersFromFile")
	err = result.readPeersFromFile(peersFilePath)
	if err != nil {
		return nil, err
	}

	log.Debugln("Client: NewKeys")
	result.keys, err = result.sam.NewKeys()
	if err != nil {
		return nil, err
	}

	log.Debugln("Client: connectToPeers")
	err = result.connectToPeers()
	if err != nil {
		return nil, err
	}
	return result, nil
}
