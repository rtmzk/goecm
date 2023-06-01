package options

import (
	"github.com/spf13/pflag"
	"go-ecm/utils"
	"time"
)

type AgentRegistryOptions struct {
	ServerAddr        string        `json:"server-addr" mapstructure:"server-addr"`
	Token             string        `json:"token" mapstructure:"token"`
	ReportInterval    time.Duration `json:"report-interval" mapstructure:"report-interval"`
	HeartBeatInterval time.Duration `json:"heartbeat-interval" mapstructure:"heartbeat-interval"`
}

func NewAgentRegistryOptions() *AgentRegistryOptions {
	return &AgentRegistryOptions{
		ServerAddr:     "",
		Token:          utils.RandString(utils.Alphabet62, 16),
		ReportInterval: time.Duration(30) * time.Second,
	}
}

func (a *AgentRegistryOptions) Validate() []error {
	var errors []error

	return errors
}

func (a *AgentRegistryOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&a.ServerAddr, "agent.server-addr", a.ServerAddr, "")
	fs.StringVar(&a.Token, "agent.token", a.Token, "")
	fs.DurationVar(&a.ReportInterval, "agent.report-interval", a.ReportInterval, "")
	fs.DurationVar(&a.HeartBeatInterval, "agent.heartbeat-interval", a.HeartBeatInterval, "")
}
