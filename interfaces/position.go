package interfaces

import responses "multispore/types"

// TEMP!!!
type Location interface {
	Join(id int, channel chan<- responses.Response)

	// Disconnects the player by id
	Leave(id int)

	// Send message to everyone in the location
	Broadcast(arg ...string)

	// Send message to other clients in the location (Not going to send to the source)
	Announce(id int, arg ...string)
}
