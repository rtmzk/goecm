package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	v1 "go-ecm/internal/goecmagent/model/v1"
	"go-ecm/pkg/core"
	"go-ecm/utils"
	"os"
	"os/exec"
	"path/filepath"
)

func DeployCore(c *gin.Context) {
	var action = c.Param("action")
	var err error
	switch action {
	case "CheckDataPath":
		err = checkDataPath(c)
	case "LoadImage":
		err = loadImage(c)
	}

	if err != nil {
		core.WriteResponse(c, err, map[string]string{"result": "failed"})
		return
	}
	core.WriteResponse(c, err, map[string]string{"result": "ok"})
}

func checkDataPath(c *gin.Context) error {
	var deploySpec v1.DeploySpec
	var paths []string
	var errs []error
	err := c.ShouldBindJSON(&deploySpec)
	if err != nil {
		return err
	}
	paths = append(paths,
		deploySpec.Elasticsearch.ElasticsearchDataPath,
		deploySpec.Elasticsearch.ElasticsearchBackupPath,
		deploySpec.Database.DBDataPath,
		deploySpec.Database.DBBackupPath,
		deploySpec.Redis.RedisDataPath,
		deploySpec.Rabbitmq.RabbitmqDataPath,
		deploySpec.ECMDeploySpec.ECMStorageSpec.StoragePath,
	)
	for _, p := range paths {
		if utils.IsNotExist(p) {
			err = os.MkdirAll(p, os.ModePerm)
		}
	}
	if len(errs) != 0 {
		return errors.New("create data path failed.")
	}

	return nil
}

//func loadImage(c *gin.Context) error {
//	loaded := c.Query("loaded")
//	if loaded == "true" {
//		return nil
//	}
//	cli, err := docker.NewDockerClient()
//	if err != nil {
//		return err
//	}
//	var path = filepath.Join("/opt", c.Query("name"))
//	fd, err := os.Open(path)
//	defer fd.Close()
//	if err != nil {
//		return err
//	}
//
//	resp, err := cli.LoadImage(c, fd, false)
//	if err != nil {
//		return err
//	}
//	fmt.Printf("%#v", resp)
//	return nil
//}

func loadImage(c *gin.Context) error {
	loaded := c.Query("loaded")
	if loaded == "true" {
		return nil
	}
	var path = filepath.Join("/opt", c.Query("name"))
	var cmd = exec.Command("docker", "image", "load", "-i", path)
	return cmd.Run()
}
