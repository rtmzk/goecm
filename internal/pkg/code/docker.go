package code

const (
	ErrGetDockerClient = iota + 100020
	ErrGetServiceList
	ErrGetServiceSpec
	ErrServiceUpdate
	ErrNodeList
	ErrNodeDelete
	ErrNodeInspect
	ErrNodeUpdate
	ErrImagePullWithConnectionErr
	ErrImagePullWithConnectionErrWrapOK = iota + ErrImagePullWithConnectionErr
)
