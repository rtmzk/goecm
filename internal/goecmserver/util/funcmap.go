package util

import (
	v1 "go-ecm/internal/goecmserver/model/v1"
	"strings"
)

func NewFuncMap() map[string]any {
	return map[string]any{
		"setLocalOrNasDir": func(data v1.DeploySpec) string {
			var out string
			if data.StorageType != "0" {
				out = "  #- /home/edoc2/macrowing/edoc2v5/edoc2Docs"
			} else {
				out = `  - ` + data.StoragePath
			}
			return out
		},
		"setDockerNet": func(data string, flag int) string {
			net := strings.Split(data, "/")
			return net[flag]
		},
	}
}
