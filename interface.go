package p2p

import (
	"github.com/libsv/go-p2p/wire"
)

type Status int32

var (
	StatusSent     Status = 5
	StatusSeen     Status = 6
	StatusRejected Status = 109
)

type PeerManagerI interface {
	AnnounceTransaction(txID []byte, peers []PeerI) []PeerI
	GetTransaction(txID []byte)
	AddPeer(peerURL string, peerStore PeerHandlerI) error
	RemovePeer(peerURL string) error
	GetPeers() []PeerI
	PeerCreator(peerCreator func(peerAddress string, peerStore PeerHandlerI) (PeerI, error))
	addPeer(peer PeerI) error
}

type PeerI interface {
	Connected() bool
	WriteMsg(msg wire.Message) error
	String() string
	AnnounceTransaction(txID []byte)
}

type PeerHandlerI interface {
	HandleTransactionGet(msg *wire.InvVect, peer PeerI) ([]byte, error)
	HandleTransactionSent(msg *wire.MsgTx, peer PeerI) error
	HandleTransactionAnnouncement(msg *wire.InvVect, peer PeerI) error
	HandleTransactionRejection(rejMsg *wire.MsgReject, peer PeerI) error
	HandleTransaction(msg *wire.MsgTx, peer PeerI) error
	HandleBlockAnnouncement(msg *wire.InvVect, peer PeerI) error
	HandleBlock(msg *BlockMessage, peer PeerI) error
}
