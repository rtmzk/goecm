package app

import cliflag "go-ecm/utils/flag"

type CliOptions interface {
	Flags() (fss cliflag.NamedFlagSets)
	Validate() []error
}

type ConfigurableOptions interface {
	ApplyFlag() []error
}

type CompleteableOptions interface {
	Complete() error
}

type PritableOptions interface {
	String() string
}
