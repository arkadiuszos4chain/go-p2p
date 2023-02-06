package p2p

import (
	"sync"

	"github.com/libsv/go-p2p/wire"
)

type PeerMock struct {
	mu            sync.Mutex
	address       string
	peerStore     PeerHandlerI
	writeChan     chan wire.Message
	messages      []wire.Message
	announcements [][]byte
}

func NewPeerMock(address string, peerStore PeerHandlerI) (*PeerMock, error) {
	writeChan := make(chan wire.Message)

	p := &PeerMock{
		peerStore: peerStore,
		address:   address,
		writeChan: writeChan,
	}

	go func() {
		for msg := range writeChan {
			p.message(msg)
		}
	}()

	return p, nil
}

func (p *PeerMock) Connected() bool {
	return true
}

func (p *PeerMock) Len() int {
	p.mu.Lock()
	defer p.mu.Unlock()

	return len(p.messages)
}

func (p *PeerMock) AnnounceTransaction(txID []byte) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.announcements = append(p.announcements, txID)
}

func (p *PeerMock) getAnnouncements() [][]byte {
	p.mu.Lock()
	defer p.mu.Unlock()

	return p.announcements
}

func (p *PeerMock) message(msg wire.Message) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.messages = append(p.messages, msg)
}

// func (p *PeerMock) getMessages() []wire.Message {
// 	p.mu.Lock()
// 	defer p.mu.Unlock()

// 	return p.messages
// }

func (p *PeerMock) WriteMsg(msg wire.Message) error {
	p.writeChan <- msg
	return nil
}

func (p *PeerMock) String() string {
	return p.address
}
