package util

import (
	"fmt"
	"net/http"
)

func ImageImportHelper(filePath string, addr string) error {
	var url = fmt.Sprintf("http://%s/v1/image/import?filePath=%s", addr, filePath)
	cli := http.DefaultClient
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return err
	}
	resp, err := cli.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	return nil
}
