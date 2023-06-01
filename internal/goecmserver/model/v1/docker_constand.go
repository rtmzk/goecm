package v1

import (
	"github.com/docker/docker/api/types"
)

// Docker instance operation api url.
// Do not change them.
// Some url need argument, but DO NOT CHANGE THEM.
// example:
// func xx(id string) {
//		var url = ServiceSpec
//		url.Url = ServiceSpec.Url+id
// }

type OperationAPI struct {
	Method string
	Url    string
}

var (
	ImageList OperationAPI = OperationAPI{Method: "GET", Url: "http://localhost/v1.41/images/json"}
	ImageLoad OperationAPI = OperationAPI{Method: "POST", Url: "http://localhost/v1.41/images/load"}

	// ImageDelete Url need an argument: image full name
	ImageDelete OperationAPI = OperationAPI{Method: "DELETE", Url: "http://localhost/v1.41/images/"}
)

type Images struct {
	Host    string
	Summary *[]types.ImageSummary
}

type ImageDeleteRequest struct {
	Host []string `json:"hosts"`
	Ids  []string `json:"ids"`
}

type ServiceOperationRequest struct {
	Id      string `json:"id,omitempty"`
	Version string `json:"version,omitempty"`
	Name    string `json:"name,omitempty"`
}
