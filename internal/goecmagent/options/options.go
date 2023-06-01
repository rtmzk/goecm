package options

import (
	commonoptions "go-ecm/internal/pkg/options"
	"go-ecm/internal/pkg/server"
	"go-ecm/pkg/log"
	cliflag "go-ecm/utils/flag"
)

type Options struct {
	GenericServerRunOptions *commonoptions.ServerRunOptions       `json:"server" mapstructure:"server"`
	InsecureServing         *commonoptions.InsecureServingOptions `json:"insecure" mapstructure:"insecure"`
	SecureServing           *commonoptions.SecureServingOptions   `json:"secure" mapstructure:"secure"`
	Log                     *log.Options                          `json:"log" mapstructure:"log"`
	Agent                   *commonoptions.AgentRegistryOptions   `json:"agent" mapstructure:"agent"`
}

func NewOptions() *Options {
	o := Options{
		GenericServerRunOptions: commonoptions.NewServerRunOptions(),
		InsecureServing:         commonoptions.NewInsecureServingOptions(),
		SecureServing:           commonoptions.NewSecureServingOptions(),
		Log:                     log.NewOptions(),
		Agent:                   commonoptions.NewAgentRegistryOptions(),
	}

	return &o
}

func (o *Options) ApplyTo(c *server.Config) error {
	return nil
}

func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
	o.GenericServerRunOptions.AddFlags(fss.FlagSet("generic"))
	o.InsecureServing.AddFlags(fss.FlagSet("insecure serving"))
	o.SecureServing.AddFlags(fss.FlagSet("secure serving"))
	o.Log.AddFlags(fss.FlagSet("logs"))
	o.Agent.AddFlags(fss.FlagSet("agent"))

	return fss
}
