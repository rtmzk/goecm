package app

import (
	"strings"

	"github.com/spf13/pflag"
)

// noused
func initFlag() {
	pflag.CommandLine.SetNormalizeFunc(WordSepNormalizeFunc)
}

// WordSepNormalizeFunc changes all flags that contain "_" separators.
func WordSepNormalizeFunc(_ *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name, "_") {
		return pflag.NormalizedName(strings.Replace(name, "_", "-", -1))
	}

	return pflag.NormalizedName(name)
}
