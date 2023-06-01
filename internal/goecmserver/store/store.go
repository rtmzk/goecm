package store

var client Factory

type Factory interface {
	Users() UserStore
	Swarm() SwarmStore
	Menu() MenuStore
	Docker() DockerStore
	SSHKey() SSHkeyStore
	SystemConfig() SystemConfigStore
	Agent() AgentStore
	Close() error
}

func Client() Factory {
	return client
}

func SetClient(factory Factory) {
	client = factory
}
