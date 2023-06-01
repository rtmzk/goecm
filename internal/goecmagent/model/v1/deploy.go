package v1

type DeploySpec struct {
	MiddlewareDeploySpec
	ECMDeploySpec `json:"app"`
}

type ECMDeploySpec struct {
	ECMStorageSpec                `json:"storage"`
	Hosts                         string `json:"hosts"`
	Scheme                        string `json:"scheme"`
	AccessPort                    int    `json:"access_port"`
	Docker0Network                string `json:"docker0_network"`
	DockerGwbridgeNetwork         string `json:"docker_gwbridge_network"`
	DockerOverlayMacrowingNetwork string `json:"docker_overlay_macrowing_network"`
}

type ECMStorageSpec struct {
	StorageType        string `json:"storage_type"`
	StoragePath        string `json:"storage_path,omitempty"`
	StorageUrl         string `json:"storage_url,omitempty"`
	StorageHealthCheck string `json:"storage_health_check,omitempty"`
	StorageAK          string `json:"storage_ak,omitempty"`
	StorageSK          string `json:"storage_sk,omitempty"`
	StorageBucket      string `json:"storage_bucket,omitempty"`
}

type MiddlewareDeploySpec struct {
	Mode          string                  `json:"middleware_mode"`
	Redis         RedisDeploySpec         `json:"redis"`
	Rabbitmq      RabbitmqDeploySpec      `json:"rabbitmq"`
	Database      DatabaseDeploySpec      `json:"database"`
	Elasticsearch ElasticsearchDeploySpec `json:"elasticsearch"`
}

type DatabaseDeploySpec struct {
	IsExternal   bool   `json:"is_external"`
	DBType       string `json:"db_type"`
	DBHosts      string `json:"db_hosts"`
	DBPort       string `json:"db_port,omitempty"`
	DBUser       string `json:"db_user,omitempty"`
	DBPass       string `json:"db_pass,omitempty"`
	DBDataPath   string `json:"db_data_path,omitempty"`
	DBBackupPath string `json:"db_backup_path,omitempty"`
}

type RedisDeploySpec struct {
	RedisHosts    string `json:"redis_hosts"`
	RedisDataPath string `json:"redis_data_path"`
}

type ElasticsearchDeploySpec struct {
	ElasticsearchHosts      string `json:"elasticsearch_hosts"`
	ElasticsearchDataPath   string `json:"elasticsearch_data_path"`
	ElasticsearchBackupPath string `json:"elasticsearch_backup_path,omitempty"`
}

type RabbitmqDeploySpec struct {
	RabbitmqHosts    string `json:"rabbitmq_hosts"`
	RabbitmqDataPath string `json:"rabbitmq_data_path"`
}

type SwarmJoin struct {
	SwarmID      string `json:"swarm_id,omitempty" form:"swarm_id" gorm:"column:swarm_id;comment:'swarm cluster id'"`
	MgrJoinToken string `json:"mgr_join_token" form:"mgr_join_token" gorm:"column:mgr_join_token;comment:'swarm manager join-token'"`
	WrkJoinToken string `json:"wrk_join_token" form:"wrk_join_token" gorm:"column:wrk_join_token;comment:'swarm worker join-token'"`
}
