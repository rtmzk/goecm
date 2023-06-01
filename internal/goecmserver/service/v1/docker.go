package v1

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
	"github.com/marmotedu/errors"
	v1 "go-ecm/internal/goecmserver/model/v1"
	"go-ecm/internal/goecmserver/store"
	"go-ecm/internal/pkg/code"
	metav1 "go-ecm/internal/pkg/meta/v1"
	"go-ecm/pkg/log"
	"sync"
)

type DockerSrv interface {
	ImageGet(ctx context.Context, hosts []string, options metav1.GetOptions) ([]*v1.Images, error)
	ImageDelete(ctx context.Context, Ids *v1.ImageDeleteRequest, options metav1.DeleteOptions) error
	ImageLoad(ctx context.Context, path []string) error
	ServiceGet(ctx context.Context, options metav1.ListOptions) ([]swarm.Service, error)
	ServiceInspect(ctx context.Context, id string, options metav1.GetOptions) (*swarm.Service, error)
	ServiceUpdate(ctx context.Context, service *swarm.Service, options metav1.UpdateOptions) error
	NodeGet(ctx context.Context, options metav1.GetOptions) ([]swarm.Node, error)
	NodeInspect(ctx context.Context, nodeId string, options metav1.GetOptions) (*swarm.Node, error)
	NodeUpdate(ctx context.Context, node *swarm.Node, options metav1.UpdateOptions) error
	NodeDelete(ctx context.Context, nodeId string, options metav1.DeleteOptions) error
}

type dockerService struct {
	store store.Factory
	cli   *client.Client
}

var _ DockerSrv = (*dockerService)(nil)

func NewDockerSrv(srv *service) *dockerService {
	return &dockerService{srv.store, srv.cli}
}

// ImageGet  bad performance.
func (s *dockerService) ImageGet(ctx context.Context, hosts []string, options metav1.GetOptions) ([]*v1.Images, error) {
	var diqo *v1.DockersInstanceQueryOption
	var operation = v1.ImageList
	var result []*v1.Images
	diqo = v1.NewDockerInstanceQueryOption(operation)
	for _, h := range hosts {
		res, err := diqo.ImageListByNode(h, 22)
		if err != nil {
			return nil, err
		}
		result = append(result, res)
	}

	return result, nil
}

// ImageDelete Bad performance
// TODO... 判断镜像是否有容器在使用，如果有在使用 会返回http 409.
func (s *dockerService) ImageDelete(ctx context.Context, Ids *v1.ImageDeleteRequest, options metav1.DeleteOptions) error {
	var diqo *v1.DockersInstanceQueryOption
	var operation = v1.ImageDelete
	var err error
	for _, h := range Ids.Host {
		for _, n := range Ids.Ids {
			operation.Url = v1.ImageDelete.Url + n
			diqo = v1.NewDockerInstanceQueryOption(operation)
			if err = diqo.ImageDeleteByNode(h, 22); err != nil {
				err = errors.Wrap(err, "镜像删除失败")
			}
		}
	}
	return err
}

// TODO... 批量删除
func (s *dockerService) ImageDeleteCollection(ctx context.Context) {}

// ImageLoad 并行导入 good performance.
func (s *dockerService) ImageLoad(ctx context.Context, name []string) error {
	// 获取所有节点
	ret, err := s.store.Swarm().NodeList(ctx, metav1.ListOptions{})
	if err != nil {
		return errors.WithCode(code.ErrDatabase, "无法拉取节点列表")
	}

	var wg sync.WaitGroup
	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	wg.Add(len(ret.Items))

	for _, v := range ret.Items {
		clone := v
		go func(n []string) {
			defer wg.Done()
			var diqo *v1.DockersInstanceQueryOption
			var operation = v1.ImageLoad
			for _, img := range n {
				diqo = v1.NewDockerInstanceQueryOption(operation)
				err := diqo.ImageLoadByNode(clone.HostIP, img, clone.SSHPort)
				if err != nil {
					errChan <- errors.WithCode(code.ErrUnknown, "Unknown")
				}
			}
		}(name)
	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChan:
		return err
	}

	return nil
}

func (s *dockerService) ServiceGet(ctx context.Context, opts metav1.ListOptions) ([]swarm.Service, error) {
	summary, err := s.cli.ServiceList(ctx, types.ServiceListOptions{})
	if err != nil {
		return nil, errors.WithCode(code.ErrGetServiceList, "无法获取服务列表，请检查当前docker是否启用了swarm模式. Error: %s", err.Error())
	}

	return summary, nil
}

func (s *dockerService) ServiceInspect(ctx context.Context, id string, options metav1.GetOptions) (*swarm.Service, error) {
	data, _, err := s.cli.ServiceInspectWithRaw(ctx, id, types.ServiceInspectOptions{})
	if err != nil {
		log.Errorf("无法获取服务: %s的详细信息", id)
		return nil, errors.WithCode(code.ErrGetServiceSpec, "无法获取服务详细信息. Error: %s", err.Error())
	}

	return &data, nil
}

func (s *dockerService) ServiceUpdate(ctx context.Context, service *swarm.Service, options metav1.UpdateOptions) error {
	_, err := s.cli.ServiceUpdate(ctx, service.ID, service.Version, service.Spec, types.ServiceUpdateOptions{})
	if err != nil {
		log.Errorf("服务: %s 更新失败", service.Spec.Name)
		return errors.WithCode(code.ErrServiceUpdate, "服务更新失败")
	}

	return nil
}

func (s *dockerService) NodeGet(ctx context.Context, options metav1.GetOptions) ([]swarm.Node, error) {
	nodes, err := s.cli.NodeList(ctx, types.NodeListOptions{})
	if err != nil {
		return nil, errors.WithCode(code.ErrNodeList, "获取节点列表失败")
	}

	return nodes, nil
}

func (s *dockerService) NodeDelete(ctx context.Context, nodeId string, options metav1.DeleteOptions) error {
	return s.cli.NodeRemove(ctx, nodeId, types.NodeRemoveOptions{})
}

func (s *dockerService) NodeInspect(ctx context.Context, nodeId string, options metav1.GetOptions) (*swarm.Node, error) {
	resp, _, err := s.cli.NodeInspectWithRaw(ctx, nodeId)
	if err != nil {
		return nil, errors.WithCode(code.ErrNodeInspect, "无法获取nodeId: %s的详细信息. Error: %s", nodeId, err.Error())
	}

	return &resp, nil
}

func (s *dockerService) NodeUpdate(ctx context.Context, node *swarm.Node, options metav1.UpdateOptions) error {
	err := s.cli.NodeUpdate(ctx, node.ID, node.Version, node.Spec)
	if err != nil {
		return errors.WithCode(code.ErrNodeUpdate, "更新服务时发生错误. Error: %s", err.Error())
	}

	return nil
}
