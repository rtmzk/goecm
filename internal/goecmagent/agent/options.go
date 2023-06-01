package agent

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type registryOption struct {
	Token        string `json:"token"`
	AgentAddr    string `json:"address"`
	AgentPort    int    `json:"port"`
	RemoteServer string `json:"-"`
}

type registryResponse struct {
	Success   string    `json:"success,omitempty"`
	Code      int       `json:"code,omitempty"`
	Message   string    `json:"message"`
	AgentUUID uuid.UUID `json:"uuid,omitempty"`
	SwarmKey  string    `json:"swarmKey"`
}

type heartbeatRequest struct {
	Token             string        `json:"token"`
	AgentAddr         string        `json:"address"`
	AgentPort         int           `json:"port"`
	RemoteServerAddr  string        `json:"-"`
	HeartbeatInterval time.Duration `json:"-"`
}
