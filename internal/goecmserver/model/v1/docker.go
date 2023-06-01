package v1

import (
	"encoding/json"
	"github.com/docker/docker/api/types"
	"github.com/marmotedu/errors"
	"go-ecm/pkg/log"
	"go-ecm/utils"
	myssh "go-ecm/utils/ssh"
)

type DockersInstanceQueryOption struct {
	OperationAPI
	SshClient *myssh.Client
}

type Option func(*DockersInstanceQueryOption)

func NewDockerInstanceQueryOption(api OperationAPI) *DockersInstanceQueryOption {
	d := &DockersInstanceQueryOption{}
	d.OperationAPI = api

	return d
}

func (d *DockersInstanceQueryOption) ImageListByNode(h string, port int) (*Images, error) {
	var res = new(Images)
	var dockerTypes = new([]types.ImageSummary)
	client := myssh.DefaultSSHClient(h, port)
	d.SshClient = client
	summary, err := d.SshClient.Output("curl -s --unix-socket /var/run/docker.sock -X" + d.Method + " " + d.Url)
	if err != nil {
		log.Info(err.Error())
	}
	err = json.Unmarshal(summary, &dockerTypes)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	res.Host = h
	res.Summary = dockerTypes

	return res, nil
}

func (d *DockersInstanceQueryOption) ImageLoadByNode(ip, path string, port int) error {
	client := myssh.DefaultSSHClient(ip, port)
	d.SshClient = client

	// TODO... scp到其他机器
	dstPath := "/tmp/" + utils.RandString(utils.Alphabet36, 7)
	err := d.SshClient.Upload(path, dstPath)
	if err != nil {
		return err
	}
	err = d.SshClient.RunWithBindStd("curl -s --unix-socket /var/run/docker.sock " +
		"-H 'Content-Type: application/x-tar' " +
		"--data-binary @" + dstPath +
		" -X" + d.Method + " " + d.Url)
	if err != nil {
		log.Errorf("host: %s load image failed. Error: %s", ip, err.Error())
		return err
	}

	return nil
}

func (d *DockersInstanceQueryOption) ImageDeleteByNode(h string, port int) error {
	client := myssh.DefaultSSHClient(h, port)
	d.SshClient = client
	_, err := d.SshClient.Output("curl -s --unix-socket /var/run/docker.sock -X" + d.Method + " " + d.Url)
	if err != nil {
		err = errors.Wrap(err, d.Url+"镜像删除失败")
		return err
	}

	return nil
}
