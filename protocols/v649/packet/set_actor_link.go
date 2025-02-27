package packet

import (
	"github.com/oomph-ac/new-mv/protocols/v649/types"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

// SetActorLink is sent by the server to initiate an entity link client-side, meaning one entity will start
// riding another.
type SetActorLink struct {
	// EntityLink is the link to be set client-side. It links two entities together, so that one entity rides
	// another. Note that players that see those entities later will not see the link, unless it is also sent
	// in the AddActor and AddPlayer packets.
	EntityLink types.EntityLink
}

// ID ...
func (*SetActorLink) ID() uint32 {
	return packet.IDSetActorLink // 41
}

func (pk *SetActorLink) Marshal(io protocol.IO) {
	protocol.Single(io, &pk.EntityLink)
}
