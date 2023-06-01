package resource

import _ "embed"

//go:embed template/middleware-master-cluster.yml.tpl
var MiddlewareClusterYML []byte

//go:embed template/middleware-master-standalone.yml.tpl
var MiddlewareStandaloneYML []byte

//go:embed template/envfile.env.tpl
var EnvFile []byte

//go:embed template/docker-compose-mainregion.yml.tpl
var MainRegionYML []byte

//go:embed template/docker-compose-inbiz.yml.tpl
var InbizYML []byte

//go:embed template/docker-compose-apm-mainregion.yml.tpl
var APMYML []byte

//go:embed check_rule/check_rule.json
var CHECK_RULES []byte

//go:embed template/daemon.json.tpl
var DOCKERDAMONECONF []byte

//go:embed docker.rpm
var DOCKER_PACKAGE []byte
