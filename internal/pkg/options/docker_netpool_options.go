package options

import (
	"github.com/spf13/pflag"
)

type DockerNetpoolOptions struct {
	Docker0Net          string `json:"docker0" mapstructure:"docker0"`
	DockerGwbridgeNet   string `json:"docker" mapstructure:"docker_gwbridge"`
	MacrowingOverlayNet string `json:"macrowing" mapstructure:"macrowing"`
}

func NewDockerNetpoolOptions() *DockerNetpoolOptions {
	return &DockerNetpoolOptions{
		Docker0Net:          "172.17.0.1/24",
		DockerGwbridgeNet:   "172.18.0.0/16",
		MacrowingOverlayNet: "10.1.0.0/24",
	}
}

func (o *DockerNetpoolOptions) Validate() []error {
	errs := []error{}

	return errs
}

func (o *DockerNetpoolOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.Docker0Net, "dockernet.docker0", o.Docker0Net, ""+
		"Specify the network of docker0 virtual bridge")

	fs.StringVar(&o.DockerGwbridgeNet, "dockernet.docker_gwbridge", o.DockerGwbridgeNet, ""+
		"Specify the network of docker_gwbridge virtual bridge")

	fs.StringVar(&o.MacrowingOverlayNet, "dockernet.macrowing", o.MacrowingOverlayNet, ""+
		"Specify the network of macrowing overlay network")
}
