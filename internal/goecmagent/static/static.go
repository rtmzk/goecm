package static

import _ "embed"

//go:embed scripts/env_prepare.sh
var ENVPREPARE []byte

//go:embed scripts/cronded_scripts.sh
var CRONDEDSCRIPT []byte

//go:embed check_rule/check_rule.json
var CHECK_RULES []byte

//go:embed check_rule/envcheck.sh
var ENVCHECK []byte

//go:embed docker.rpm
var DOCKERPKG []byte
