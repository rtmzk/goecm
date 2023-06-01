package docker

import "github.com/docker/docker/client"

var cli *client.Client

func Client() *client.Client {
	return cli
}

func SetClient(c *client.Client) {
	cli = c
}
