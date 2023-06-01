package v1

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
	uuid "github.com/satori/go.uuid"
	metav1 "go-ecm/internal/pkg/meta/v1"
	"time"
)

type Agent struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	AgentAddr         string `json:"address" gorm:"column:address;NOTNULL"`
	AgentPort         int    `json:"port" gorm:"column:port;NOTNULL"`
	ImageLoaded       bool   `json:"image_loaded" gorm:"column:loaded;default:0"`
	Deployed          bool   `json:"deployed" gorm:"column:deployed;default:0"`
	IsApp             bool   `json:"is_app,omitempty"`
	IsMiddle          bool   `json:"is_middle,omitempty"`
	Joined            bool   `json:"joined,omitempty" gorm:"column:joined;default:0"`
	Prepared          bool   `json:"prepared,omitempty" gorm:"column:prepared"`
}

type Status struct {
	Status int `json:"status" gorm:"column:status;NOTNULL"`
}

type RegistryResponse struct {
	Success   string    `json:"success,omitempty"`
	Code      int       `json:"code,omitempty"`
	Message   string    `json:"message,omitempty"`
	AgentUUID uuid.UUID `json:"uuid,omitempty"`
	SwarmKey  string    `json:"swarmKey,omitempty"`
}
type Actions struct {
	Action string      `json:"type"`
	Data   interface{} `json:"data"`
	Query  string      `json:"query,omitempty"`
}

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

type ImageUpload struct {
	Chunk     string `json:"chunk"`
	Chunks    string `json:"chunks"`
	ChunkHash string `json:"chunkHash"`
	FileHash  string `json:"fileHash"`
	FileName  string `json:"fileName"`
}

type AllReportItems struct {
	ImageReportItem
	ContainerReportItem
	ServiceReportItem
}

type ImageReportItem struct {
	types.ImageSummary
	Host string `json:"host"`
}

type ContainerReportItem struct {
	types.Container
	Host string `json:"host"`
}

type ServiceReportItem struct {
	swarm.Service
	Host string `json:"host"`
}

type GetImagesReportItems struct {
	ImageItem []ImageReportItem
}

type GetContainerReportItems struct {
	ContainerItem []ContainerReportItem
}

type GetServiceReportItems struct {
	ServiceItem []ServiceReportItem
}

type ContainerOperation struct {
	Host        string   `json:"host"`
	ContainerId []string `json:"containerId"`
}
type ImageOperation struct {
	Host    string   `json:"host"`
	ImageId []string `json:"imageId"`
}

type Heartbeat struct {
	Token             string        `json:"token"`
	AgentAddr         string        `json:"address"`
	AgentPort         int           `json:"port"`
	RemoteServerAddr  string        `json:"-"`
	HeartbeatInterval time.Duration `json:"-"`
}

type ImagePull struct {
	Node      string `json:"node"`
	ImageName string `json:"imageName"`
}

type MultiOperationErr struct {
	Node    []string `json:"node"`
	Message string   `json:"message"`
}

type Rule struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Func        string `json:"func"`
	ScriptOut
}

type CheckRules struct {
	Rules []Rule `json:"rules"`
}

type ScriptOut struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
