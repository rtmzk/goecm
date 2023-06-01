package options

import "time"

type JwtOptions struct {
	Realm      string
	Key        string
	Timeout    time.Duration
	MaxRefresh time.Duration
}

func (jwt *JwtOptions) Validate() []error {
	var errors []error

	return errors
}
