package docker

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
	"io"
)

type DockerOperator interface {
	GetImages(ctx context.Context, options types.ImageListOptions) ([]types.ImageSummary, error)
	GetImageInspect(id string) (*types.ImageInspect, error)
	GetNodes(options types.NodeListOptions) ([]swarm.Node, error)
	GetServices(ctx context.Context, options types.ServiceListOptions) ([]swarm.Service, error)
	GetContainers(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error)
	GetContainerInspect(id string) (*types.ContainerJSON, error)
	LoadImage(ctx context.Context, reader io.Reader, quite bool) (types.ImageLoadResponse, error)
	UpdateNode(node *swarm.Node) error
	Exec(ctx context.Context, containerId, command string) (types.HijackedResponse, error)
	Resize(ctx context.Context, containerId, w, h string) error
	DeleteContainer(ctx context.Context, containerId string, option types.ContainerRemoveOptions) error
	DeleteImage(ctx context.Context, imageId string, option types.ImageRemoveOptions) error
	ExportImage(ctx context.Context, imageIds []string) (io.ReadCloser, error)
	PullImage(ctx context.Context, imageName string, option types.ImagePullOptions) (io.ReadCloser, error)
	SwarmJoin(ctx context.Context, req swarm.JoinRequest) error
	NetworkList(ctx context.Context, options types.NetworkListOptions) ([]types.NetworkResource, error)
	NetworkCreate(ctx context.Context, name string, options types.NetworkCreate) error
	NodeInspect(ctx context.Context, nodeId string) (swarm.Node, error)
}
