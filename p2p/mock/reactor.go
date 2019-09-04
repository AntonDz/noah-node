package mock

import (
	p2p2 "github.com/noah-blockchain/noah-node/p2p"
	"github.com/noah-blockchain/noah-node/p2p/conn"
	"github.com/tendermint/tendermint/libs/log"
)

type Reactor struct {
	p2p2.BaseReactor
}

func NewReactor() *Reactor {
	r := &Reactor{}
	r.BaseReactor = *p2p2.NewBaseReactor("Reactor", r)
	r.SetLogger(log.TestingLogger())
	return r
}

func (r *Reactor) GetChannels() []*conn.ChannelDescriptor             { return []*conn.ChannelDescriptor{} }
func (r *Reactor) AddPeer(peer p2p2.Peer)                             {}
func (r *Reactor) RemovePeer(peer p2p2.Peer, reason interface{})      {}
func (r *Reactor) Receive(chID byte, peer p2p2.Peer, msgBytes []byte) {}
