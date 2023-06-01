package server

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-ecm/pkg/log"
	"go-ecm/utils"
	"net"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	RecommendedHomeDir = ".goecm"
	RecommendEnvPrefix = "GOECM"
)

type Config struct {
	SecureServing   *SecureServingInfo
	InSecureServing *InSecureServingInfo
	Mode            string
	Middlewares     []string
	Healthz         bool
	EnableMetrics   bool
	EnableProfiling bool
}

type SecureServingInfo struct {
	BindAddress string
	BindPort    int
	CertKey     CertKey
}

func (s *SecureServingInfo) Address() string {
	return net.JoinHostPort(s.BindAddress, strconv.Itoa(s.BindPort))
}

type InSecureServingInfo struct {
	Address string
}

type CertKey struct {
	CertFile string
	KeyFile  string
}

func NewConfig() *Config {
	return &Config{
		Healthz:     true,
		Mode:        gin.ReleaseMode,
		Middlewares: []string{},
	}
}

type CompletedConfig struct {
	*Config
}

func (c *Config) Complete() CompletedConfig {
	return CompletedConfig{c}
}

func (c CompletedConfig) New() (*GenericServer, error) {
	s := &GenericServer{
		SecureServingInfo:   c.SecureServing,
		InSecureServingInfo: c.InSecureServing,
		mode:                c.Mode,
		healthz:             c.Healthz,
		enableMetric:        c.EnableMetrics,
		enableProfiling:     c.EnableProfiling,
		middlewares:         c.Middlewares,
		Engine:              gin.New(),
	}

	initGenericServer(s)
	return s, nil
}

func LoadConfig(cfg string, defaultName string) {
	if cfg != "" {
		viper.SetConfigFile(cfg)
	} else {
		viper.AddConfigPath(".")
		viper.AddConfigPath(filepath.Join(utils.UserHome(), RecommendedHomeDir))
		viper.SetConfigFile(defaultName)
	}

	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix(RecommendEnvPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	if err := viper.ReadInConfig(); err != nil {
		log.Warnf("WARNING: viper failed to discover and load configuration file: %s", err.Error())
	}
}
