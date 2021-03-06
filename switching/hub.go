package switching

import (
	bh "github.com/kandoo/beehive"
	"github.com/kandoo/beehive-netctrl/nom"
)

type Hub struct{}

func (h Hub) Rcv(msg bh.Msg, ctx bh.RcvContext) error {
	in := msg.Data().(nom.PacketIn)
	out := nom.PacketOut{
		Node:     in.Node,
		InPort:   in.InPort,
		BufferID: in.BufferID,
		Packet:   in.Packet,
		Actions:  []nom.Action{nom.ActionFlood{}},
	}
	ctx.ReplyTo(msg, out)
	return nil
}

func (h Hub) Map(msg bh.Msg, ctx bh.MapContext) bh.MappedCells {
	return bh.MappedCells{{"N", string(msg.Data().(nom.PacketIn).Node)}}
}
