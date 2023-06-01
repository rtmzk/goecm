package v1

import metav1 "go-ecm/internal/pkg/meta/v1"

type SSHAuthMode int

type HostInfo struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	HostIP            string `json:"host_ip" form:"host_ip" gorm:"unique;column:host_ip;comment:'初始化节点IP'"`
	// HostName    string `json:"host_name" form:"host_name" gorm:"column:host_name;comment:'初始化节点hostname'"`
	SSHPort     int         `json:"ssh_port" form:"ssh_port" gorm:"column:ssh_port;comment:'节点的ssh port'"`
	SSHAuthMode SSHAuthMode `json:"ssh_auth_mode" form:"ssh_auth_mode" gorm:"column:ssh_auth_mode;comment:'节点登录验证方式,1:key,2:password'"`
	SSHUser     string      `json:"ssh_user" form:"ssh_user" gorm:"column:ssh_user;comment:'节点ssh用户'"`
	SSHPass     string      `json:"ssh_pass" form:"ssh_pass" gorm:"column:ssh_pass;comment:'节点ssh用户密码'"`
	Connected   bool        `json:"connected" form:"connected" gorm:"default:false;column:connected;comment:'节点ssh是否可连接'"`
	EnvInited   bool        `json:"env_inited" form:"env_inited" gorm:"default:false;column:env_inited;comment:'节点操作系统环境是否初始化'"`
	SwarmJoined bool        `json:"swarm_joined" form:"swarm_joined" gorm:"default:false;column:swarm_joined;comment:'节点操作系统环境是否初始化'"`
}

type HostConnectionCheck struct {
	HostIPAddress string `json:"IP"`
	SSHUsername   string `json:"Username"`
	SSHPort       string `json:"SSHPort"`
}

type KeyInfo struct {
	PrivateKey string `json:"private_key" form:"private_key" gorm:"column:private_key;comment:'节点ssh用户私钥'"`
	PublicKey  string `json:"public_key" form:"public_key" gorm:"column:public_key;comment:'节点ssh用户公钥'"`
}

type SwarmJoin struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	SwarmID           string `json:"swarm_id,omitempty" form:"swarm_id" gorm:"column:swarm_id;comment:'swarm cluster id'"`
	MgrJoinToken      string `json:"mgr_join_token" form:"mgr_join_token" gorm:"column:mgr_join_token;comment:'swarm manager join-token'"`
	WrkJoinToken      string `json:"wrk_join_token" form:"wrk_join_token" gorm:"column:wrk_join_token;comment:'swarm worker join-token'"`
}

type SwarmNodeSpec struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	NodeId            string              `json:"node_id" form:"node_id" gorm:"column:node_id;comment:'swarm cluster node id'"`
	Hostname          string              `json:"hostname" form:"hostname" gorm:"column:hostname;comment:'hostname'"`
	Role              string              `json:"role" form:"role" gorm:"column:role;comment:'swarm cluster node role'"`
	Labels            []map[string]string `json:"labels" form:"labels" gorm:"column:labels;comment:'swarm cluster node labels'"`
	Availability      string              `json:"availability" form:"availability" gorm:"column:availability;comment:'swarm cluster node availability'"`
	State             string              `json:"state" form:"state" gorm:"column:state;comment:'swarm cluster node state'"`
	Reachability      string              `json:"reachability" form:"reachability" gorm:"column:reachability;comment:'swarm cluster node reachability'"`
	Addr              string              `json:"address" form:"address" gorm:"column:address;comment:'swarm cluster node address'"`
}

type NodeList struct {
	Items []*HostInfo `json:"items"`
}
