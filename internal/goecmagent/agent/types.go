package agent

import (
	"github.com/docker/docker/api/types"
)

type ReportItems struct {
	Host  string     `json:"host"`
	Items ReportItem `json:"items"`
}

type ReportItem struct {
	ContainerItem []types.Container    `json:"containers"`
	ImageItem     []types.ImageSummary `json:"images"`
	//	ServiceItem   []swarm.Service      `json:"services"`
	//	NetworkItem   []swarm.Network      `json:"networks"`
}
