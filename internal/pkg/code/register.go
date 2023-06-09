package code

func init() {
	register(ErrUserNotFound, 404, "User not found.")
	register(ErrDatabase, 500, "Database Error")
	register(ErrSSHConnectionFailed, 400, "Error opening sshkey connection.")
	register(ErrSuccess, 200, "OK")
	register(ErrSignatureInvalid, 401, "Signature is  invalid")
	register(ErrPermissionDenied, 403, "Permission denied")
	register(ErrBind, 400, "Error occurred while binding the request to the struct")
	register(ErrTokenInvalid, 401, "Token invalid")
	register(ErrUnknown, 500, "Internal server error")
	register(ErrUserAlreadyExist, 400, "User already exist")
	register(ErrValidation, 400, "Validation failed")
	register(ErrUnmarshal, 400, "Object Unmarshal failed.")
	register(ErrPrivateKeyNotFound, 500, "Ssh private key not found")
	register(ErrSystemStorageIssue, 500, "Can not save file.")
	register(ErrCommandExec, 500, "Command exec failed.")
	register(ErrGetDockerClient, 500, "Failed get docker client")
	register(ErrGetServiceList, 500, "Failed get service list")
	register(ErrGetServiceSpec, 500, "Failed get service spec")
	register(ErrServiceUpdate, 500, "Failed update service")
	register(ErrNodeList, 500, "Failed list swarm nodes")
	register(ErrNodeDelete, 500, "Failed remove swarm nodes")
	register(ErrNodeUpdate, 500, "Failed update swarm nodes")
	register(ErrAlreadyRegister, 200, "Agent Already registered")
	register(ErrUpload, 500, "Upload file to agent failed")
	register(ErrFlushTemplate, 200, "Can not flush template file by input struct. May be server error.")
	register(ErrImagePullWithConnectionErr, 500, "无法拉取镜像,可能未信任该仓库或需要鉴权")
	register(ErrImagePullWithConnectionErrWrapOK, 200, "无法拉取镜像,可能未信任该仓库或需要鉴权")
	register(ErrNetConnection, 200, "无法获取目标主机的检查信息")
	register(ErrImageUpload, 200, "无法保存块数据")
	register(ErrImageImportMerge, 200, "无法合并镜像分块文件")
}
