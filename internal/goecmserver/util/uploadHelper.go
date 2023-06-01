package util

import (
	"bytes"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"math"
	"net/http"
	"os"
)

func MultiPartUploadHelper(path, dst, remotePath string, size int) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	uploadId, _ := uuid.NewV4()

	fileInfo, _ := file.Stat()
	var fileSize = fileInfo.Size()
	var chunkSize = size * (1 << 20)

	totalPartNum := uint64(math.Ceil(float64(fileSize) / float64(chunkSize)))

	for i := uint64(0); i < totalPartNum; i++ {
		var isEnd = false
		partSize := int(math.Min(float64(chunkSize), float64(fileSize-int64(i*uint64(chunkSize)))))
		partBuf := make([]byte, partSize)

		_, err := file.Read(partBuf)
		if err != nil {
			fmt.Println(err)
			return err
		}

		if i == totalPartNum-1 {
			isEnd = true
		}

		url := fmt.Sprintf("http://%s/v1/multiPartUpload?uploadId=%s&path=%s&name=%s&end=%t&chunk=%d", dst, uploadId.String(), remotePath, fileInfo.Name(), isEnd, i)
		resp, err := http.Post(url, "multipart/form-data", bytes.NewReader(partBuf))
		if err != nil {
			return err
		}
		defer resp.Body.Close()
	}

	return nil
}
