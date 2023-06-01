package agent

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	"go-ecm/internal/goecmserver/constand"
	"go-ecm/internal/goecmserver/util"
	"go-ecm/internal/pkg/code"
	"go-ecm/internal/pkg/docker"
	"go-ecm/pkg/core"
	"go-ecm/utils"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
)

// ImageImport
// 1. 合并文件    25
// 2. 本地导入    25
// 3. 包传到其他节点  25
// 4. 其他导入文件  25
func (a *AgentController) ImageImport(c *gin.Context) {
	var totalChunk, _ = strconv.Atoi(c.PostForm("chunks"))
	var fileName = c.PostForm("fileName")
	var hosts = strings.Split(c.PostForm("hosts"), ",")
	core.WriteResponse(c, nil, map[string]bool{"task": true})
	var filterd []string
	var mu sync.Mutex
	resetImageImportProgress()
	updateImportProgress(&mu, 50, "starting", nil, false)

	//Merge file
	err := mergeAndLoad(totalChunk, fileName, &mu)
	updateImportProgress(&mu, float64(25)/float64(2), "本地已导入", err, false)
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrImageImportMerge, "合并镜像分块文件失败"), nil)
		return
	}

	var server = strings.Split(c.Request.Host, ":")[0]
	for k, host := range hosts {
		if strings.Split(host, ":")[0] == server {
			fmt.Println(hosts[:k])
			fmt.Println(hosts[k+1:])
			filterd = append(hosts[:k], hosts[k+1:]...)
		} else {
			filterd = append(filterd, host)
		}
	}

	if len(filterd) > 0 {
		fullPath := "/tmp/images/" + fileName
		var uploadWg sync.WaitGroup
		uploadWg.Add(len(filterd))
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			var num = len(filterd)
			for _, h := range filterd {
				err = util.MultiPartUploadHelper(fullPath, h, "/tmp/images/", 20)
				updateImportProgress(&mu, float64(25)/float64(num)/float64(2), h+"正在传输镜像文件", err, false)
				err = util.ImageImportHelper(fullPath, h)
				updateImportProgress(&mu, float64(25)/float64(num)/float64(2), h+"已导入", err, false)
			}
		}(&uploadWg)

		uploadWg.Wait()
	}

	updateImportProgress(&mu, 0, "", nil, true)
}

func updateImportProgress(mu *sync.Mutex, increment float64, message string, err error, end bool) {
	mu.Lock()
	if increment != 0 {
		constand.ImportProgress.Percent = constand.ImportProgress.Percent + increment
	}
	if err != nil {
		constand.ImportProgress.Status = "error"
	} else {
		constand.ImportProgress.Status = "importing"
	}
	constand.ImportProgress.Message = message
	if end {
		constand.ImportProgress.Status = "end"
	}
	mu.Unlock()
}

func mergeAndLoad(totalChunk int, fileName string, mu *sync.Mutex) error {
	if utils.IsNotExist("/tmp/images") {
		_ = os.MkdirAll("/tmp/images", 0755)
	}
	dst, err := os.OpenFile("/tmp/images/"+fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer dst.Close()

	for i := 1; i <= totalChunk; i++ {
		src, _ := os.Open("/tmp/" + fileName + strconv.Itoa(i))
		_, err = io.Copy(dst, src)
		if err == nil {
			_ = os.RemoveAll("/tmp/" + fileName + strconv.Itoa(i))
		}
		updateImportProgress(mu, float64((float64(1)/float64(totalChunk))*float64(25)/float64(2)), "本地已合并文件", nil, false)
	}
	return loadLocal("/tmp/images/" + fileName)
}

func loadLocal(path string) error {
	var cli, _ = docker.NewDockerClient()
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	_, err = cli.LoadImage(context.Background(), file, true)
	return err
}

func (a *AgentController) GetImageImportProgress(c *gin.Context) {
	core.WriteResponse(c, nil, constand.ImportProgress)
	if constand.ImportProgress.Status == "end" {
		resetImageImportProgress()
	}
}

func resetImageImportProgress() {
	constand.ImportProgress.Percent = 0
	constand.ImportProgress.Status = ""
	constand.ImportProgress.Message = ""
}
