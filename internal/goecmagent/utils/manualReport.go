package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/gin-gonic/gin"
	"go-ecm/internal/goecmagent/agent"
	"go-ecm/internal/goecmagent/constand"
	"go-ecm/internal/pkg/docker"
	"net/http"
	"strings"
)

func ManualReport(c *gin.Context, cli docker.DockerOperator) {
	var reportItemsResult agent.ReportItems
	var buf bytes.Buffer
	var url = fmt.Sprintf("http://%s/v1/agent/report", constand.ServerAddr)

	reportItemsResult.Host = strings.Split(c.Request.Host, ":")[0]
	//	reportItemsResult.Items.ServiceItem, _ = cli.GetServices(context.Background(), types.ServiceListOptions{})
	reportItemsResult.Items.ContainerItem, _ = cli.GetContainers(context.Background(), types.ContainerListOptions{All: true})
	reportItemsResult.Items.ImageItem, _ = cli.GetImages(context.Background(), types.ImageListOptions{})

	_ = json.NewEncoder(&buf).Encode(reportItemsResult)

	req, _ := http.NewRequest(http.MethodPost, url, &buf)
	req.Close = true
	resp, _ := http.DefaultClient.Do(req)
	if resp != nil {
		resp.Body.Close()
	}
}
