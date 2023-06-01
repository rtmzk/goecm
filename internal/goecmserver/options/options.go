package options

import (
	"encoding/json"
	genericoptions "go-ecm/internal/pkg/options"
	"go-ecm/internal/pkg/server"
	"go-ecm/pkg/log"
	"go-ecm/utils"
	cliflag "go-ecm/utils/flag"
)

type Options struct {
	GenericServerRunOptions *genericoptions.ServerRunOptions       `json:"server"   mapstructure:"server"`
	InsecureServing         *genericoptions.InsecureServingOptions `json:"insecure" mapstructure:"insecure"`
	SecureServing           *genericoptions.SecureServingOptions   `json:"secure"   mapstructure:"secure"`
	SQLiteOptions           *genericoptions.SQLiteOptions          `json:"sqlite" mapstructure:"sqlite"`
	JwtOptions              *genericoptions.JwtOptions             `json:"jwt"      mapstructure:"jwt"`
	Log                     *log.Options                           `json:"log"      mapstructure:"log"`
	DockerNetwork           *genericoptions.DockerNetpoolOptions   `json:"dockernet" mapstructure:"dockernet"`
	//MySQLOptions            *genericoptions.MySQLOptions           `json:"mysql"    mapstructure:"mysql"`
}

// NewOptions creates a new Options object with default parameters.
func NewOptions() *Options {
	o := Options{
		GenericServerRunOptions: genericoptions.NewServerRunOptions(),
		InsecureServing:         genericoptions.NewInsecureServingOptions(),
		SecureServing:           genericoptions.NewSecureServingOptions(),
		DockerNetwork:           genericoptions.NewDockerNetpoolOptions(),
		SQLiteOptions:           genericoptions.NewSqliteOptions(),
		Log:                     log.NewOptions(),
		//MySQLOptions:            genericoptions.NewMySQLOptions(),
	}

	return &o
}

// ApplyTo applies the run options to the method receiver and returns self.
func (o *Options) ApplyTo(c *server.Config) error {
	return nil
}

// Flags returns flags for a specific APIServer by section name.
func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
	o.GenericServerRunOptions.AddFlags(fss.FlagSet("generic"))
	//o.MySQLOptions.AddFlags(fss.FlagSet("mysql"))
	o.SQLiteOptions.AddFlags(fss.FlagSet("sqlite"))
	o.InsecureServing.AddFlags(fss.FlagSet("insecure serving"))
	o.SecureServing.AddFlags(fss.FlagSet("secure serving"))
	o.Log.AddFlags(fss.FlagSet("logs"))
	o.DockerNetwork.AddFlags(fss.FlagSet("dockernet"))

	return fss
}

func (o *Options) String() string {
	data, _ := json.Marshal(o)

	return string(data)
}

// Complete set default Options.
func (o *Options) Complete() error {
	if o.JwtOptions.Key == "" {
		o.JwtOptions.Key = utils.RandString(utils.Alphabet62, 32)
	}

	return o.SecureServing.Complete()
}
