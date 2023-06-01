package docker

import (
	"context"
	"github.com/docker/docker/api/types"
	"go-ecm/pkg/log"
	"io"
)

func (cli *dockerClient) GetImages(ctx context.Context, options types.ImageListOptions) ([]types.ImageSummary, error) {
	c := cli.cli
	data, err := c.ImageList(ctx, options)
	if err != nil {
		return nil, err
	}
	log.Debugf("docker/GetImages: %#v", &data)

	return data, nil
}

func (cli *dockerClient) GetImageInspect(id string) (*types.ImageInspect, error) {
	c := cli.cli
	data, _, err := c.ImageInspectWithRaw(context.Background(), id)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (cli *dockerClient) LoadImage(ctx context.Context, reader io.Reader, quite bool) (types.ImageLoadResponse, error) {
	var c = cli.cli
	return c.ImageLoad(ctx, reader, quite)
}

func (cli *dockerClient) DeleteImage(ctx context.Context, imageId string, option types.ImageRemoveOptions) error {
	var c = cli.cli
	_, err := c.ImageRemove(ctx, imageId, option)
	return err
}

func (cli *dockerClient) ExportImage(ctx context.Context, imageId []string) (io.ReadCloser, error) {
	var c = cli.cli
	return c.ImageSave(ctx, imageId)
}

func (cli *dockerClient) PullImage(ctx context.Context, imageName string, option types.ImagePullOptions) (io.ReadCloser, error) {
	return cli.cli.ImagePull(ctx, imageName, option)
}
