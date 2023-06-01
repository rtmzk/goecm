package options

import (
	"github.com/spf13/pflag"
	"go-ecm/internal/pkg/server"
)

type ServerRunOptions struct {
	Mode        string   `json:"mode" mapstructure:"mode"`
	Healthz     bool     `json:"healthz" mapstructure:"healthz"`
	Middlewares []string `json:"middlewares" mapstructure:"middlewares"`
}

func NewServerRunOptions() *ServerRunOptions {
	defaults := server.NewConfig()

	return &ServerRunOptions{
		Mode:        defaults.Mode,
		Healthz:     defaults.Healthz,
		Middlewares: defaults.Middlewares,
	}
}

func (s *ServerRunOptions) Validate() []error {
	var errors []error

	return errors
}

func (s *ServerRunOptions) Apply(c *server.Config) error {
	c.Mode = s.Mode
	c.Healthz = s.Healthz
	c.Middlewares = s.Middlewares

	return nil
}

func (s *ServerRunOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.Mode, "server.mode", s.Mode, ""+
		"Start the server in a specified server mode. Supported: debug, test, release")

	fs.BoolVar(&s.Healthz, "server.healthz", s.Healthz, ""+
		"Add self readiness check and install /healthz router")

	fs.StringSliceVar(&s.Middlewares, "server.middlewares", s.Middlewares, ""+
		"List of allowed middlewares for server,comma separated.If this list is empty default middlewares will be used")
}
