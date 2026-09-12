package interfaces

type ManagedItem interface {
	ID() int
}

type ClientManager interface {
	AddClient(item ManagedItem)

	GetClient(id int) (ManagedItem, error)

	DisconnectClient(id int)
}
