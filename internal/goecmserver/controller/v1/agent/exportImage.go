package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/component-base/pkg/time"
	"go-ecm/internal/goecmserver/constand"
	v1 "go-ecm/internal/goecmserver/model/v1"
	"go-ecm/pkg/core"
	"net/http"
	"strconv"
	"strings"
)

func (a *AgentController) ExportImage(c *gin.Context) {
	var images v1.ImageOperation
	var host string
	var buf bytes.Buffer
	if err := c.ShouldBindJSON(&images); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	var day = time.Now().Day()
	var yearday = time.Now().YearDay()
	var hour, mini, second = time.Now().Clock()
	var timestamp = strconv.Itoa(yearday) + strconv.Itoa(day) + strconv.Itoa(hour) + strconv.Itoa(mini) + strconv.Itoa(second)

	for _, h := range constand.HeartbeatItem {
		if strings.Contains(h, host) {
			host = h
		}
	}
	_ = json.NewEncoder(&buf).Encode(images.ImageId)
	url := fmt.Sprintf("http://%s/v1/image/export", host)

	req, _ := http.NewRequest(http.MethodPost, url, &buf)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	var extratHeader = map[string]string{
		"Content-Disposition": `attachment; filename=images` + timestamp + `.tar`,
	}

	c.DataFromReader(http.StatusOK, resp.ContentLength, "application/x-tar", resp.Body, extratHeader)
}
